// Copyright © 2024 Tailfin Cloud Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package entitlement

import (
	"context"
	"fmt"
	"time"

	eventmodels "github.com/openmeterio/openmeter/openmeter/event/models"
	"github.com/openmeterio/openmeter/openmeter/meter"
	"github.com/openmeterio/openmeter/openmeter/productcatalog"
	"github.com/openmeterio/openmeter/openmeter/watermill/eventbus"
	"github.com/openmeterio/openmeter/pkg/framework/entutils"
	"github.com/openmeterio/openmeter/pkg/models"
	"github.com/openmeterio/openmeter/pkg/pagination"
	"github.com/openmeterio/openmeter/pkg/slicesx"
	"github.com/openmeterio/openmeter/pkg/sortx"
)

type ListEntitlementsOrderBy string

const (
	ListEntitlementsOrderByCreatedAt ListEntitlementsOrderBy = "created_at"
	ListEntitlementsOrderByUpdatedAt ListEntitlementsOrderBy = "updated_at"
)

func (o ListEntitlementsOrderBy) Values() []ListEntitlementsOrderBy {
	return []ListEntitlementsOrderBy{
		ListEntitlementsOrderByCreatedAt,
		ListEntitlementsOrderByUpdatedAt,
	}
}

func (o ListEntitlementsOrderBy) StrValues() []string {
	return slicesx.Map(o.Values(), func(v ListEntitlementsOrderBy) string {
		return string(v)
	})
}

type ListEntitlementsParams struct {
	IDs                 []string
	Namespaces          []string
	SubjectKeys         []string
	FeatureIDs          []string
	FeatureKeys         []string
	FeatureIDsOrKeys    []string
	EntitlementTypes    []EntitlementType
	OrderBy             ListEntitlementsOrderBy
	Order               sortx.Order
	IncludeDeleted      bool
	IncludeDeletedAfter time.Time
	Page                pagination.Page
	// will be deprecated
	Limit int
	// will be deprecated
	Offset int
}

type Connector interface {
	CreateEntitlement(ctx context.Context, input CreateEntitlementInputs) (*Entitlement, error)
	OverrideEntitlement(ctx context.Context, subject string, entitlementIdOrFeatureKey string, input CreateEntitlementInputs) (*Entitlement, error)
	GetEntitlement(ctx context.Context, namespace string, id string) (*Entitlement, error)
	DeleteEntitlement(ctx context.Context, namespace string, id string) error

	GetEntitlementValue(ctx context.Context, namespace string, subjectKey string, idOrFeatureKey string, at time.Time) (EntitlementValue, error)

	GetEntitlementsOfSubject(ctx context.Context, namespace string, subjectKey models.SubjectKey) ([]Entitlement, error)
	ListEntitlements(ctx context.Context, params ListEntitlementsParams) (pagination.PagedResponse[Entitlement], error)
}

type entitlementConnector struct {
	meteredEntitlementConnector SubTypeConnector
	staticEntitlementConnector  SubTypeConnector
	booleanEntitlementConnector SubTypeConnector

	entitlementRepo  EntitlementRepo
	featureConnector productcatalog.FeatureConnector
	meterRepo        meter.Repository

	publisher eventbus.Publisher
}

func NewEntitlementConnector(
	entitlementRepo EntitlementRepo,
	featureConnector productcatalog.FeatureConnector,
	meterRepo meter.Repository,
	meteredEntitlementConnector SubTypeConnector,
	staticEntitlementConnector SubTypeConnector,
	booleanEntitlementConnector SubTypeConnector,
	publisher eventbus.Publisher,
) Connector {
	return &entitlementConnector{
		meteredEntitlementConnector: meteredEntitlementConnector,
		staticEntitlementConnector:  staticEntitlementConnector,
		booleanEntitlementConnector: booleanEntitlementConnector,
		entitlementRepo:             entitlementRepo,
		featureConnector:            featureConnector,
		meterRepo:                   meterRepo,
		publisher:                   publisher,
	}
}

func (c *entitlementConnector) CreateEntitlement(ctx context.Context, input CreateEntitlementInputs) (*Entitlement, error) {
	doInTx := func(ctx context.Context, tx *entutils.TxDriver) (*Entitlement, error) {
		// ID has priority over key
		featureIdOrKey := input.FeatureID
		if featureIdOrKey == nil {
			featureIdOrKey = input.FeatureKey
		}
		if featureIdOrKey == nil {
			return nil, &models.GenericUserError{Message: "Feature ID or Key is required"}
		}

		feature, err := c.featureConnector.GetFeature(ctx, input.Namespace, *featureIdOrKey, productcatalog.IncludeArchivedFeatureFalse)
		if err != nil || feature == nil {
			return nil, &productcatalog.FeatureNotFoundError{ID: *featureIdOrKey}
		}

		// fill featureId and featureKey
		input.FeatureID = &feature.ID
		input.FeatureKey = &feature.Key

		currentEntitlements, err := c.entitlementRepo.WithTx(ctx, tx).GetEntitlementsOfSubject(ctx, input.Namespace, models.SubjectKey(input.SubjectKey))
		if err != nil {
			return nil, err
		}
		for _, ent := range currentEntitlements {
			// you can only have a single entitlemnet per feature key
			if ent.FeatureKey == feature.Key || ent.FeatureID == feature.ID {
				return nil, &AlreadyExistsError{EntitlementID: ent.ID, FeatureID: feature.ID, SubjectKey: input.SubjectKey}
			}
		}

		connector, err := c.getTypeConnector(input)
		if err != nil {
			return nil, err
		}
		repoInputs, err := connector.BeforeCreate(input, *feature)
		if err != nil {
			return nil, err
		}

		txCtx := entutils.NewTxContext(ctx, tx)

		ent, err := c.entitlementRepo.WithTx(txCtx, tx).CreateEntitlement(txCtx, *repoInputs)
		if err != nil {
			return nil, err
		}

		err = connector.AfterCreate(txCtx, ent)
		if err != nil {
			return nil, err
		}

		err = c.publisher.Publish(ctx, EntitlementCreatedEvent{
			Entitlement: *ent,
			Namespace: eventmodels.NamespaceID{
				ID: input.Namespace,
			},
		})
		if err != nil {
			return nil, err
		}

		return ent, err
	}

	if ctxTx, err := entutils.GetTxDriver(ctx); err == nil {
		// we're already in a tx
		return doInTx(ctx, ctxTx)
	} else {
		return entutils.StartAndRunTx(ctx, c.entitlementRepo, doInTx)
	}
}

func (c *entitlementConnector) OverrideEntitlement(ctx context.Context, subject string, entitlementIdOrFeatureKey string, input CreateEntitlementInputs) (*Entitlement, error) {
	// Validate input
	if subject != input.SubjectKey {
		return nil, &models.GenericUserError{Message: "Subject key in path and body do not match"}
	}

	oldEnt, err := c.entitlementRepo.GetEntitlementOfSubject(ctx, input.Namespace, input.SubjectKey, entitlementIdOrFeatureKey)
	if err != nil {
		return nil, err
	}

	if oldEnt == nil {
		return nil, fmt.Errorf("inconsistency error, entitlement not found: %s", entitlementIdOrFeatureKey)
	}

	if oldEnt.DeletedAt != nil {
		return nil, fmt.Errorf("inconsistency error, entitlement already deleted: %s", oldEnt.ID)
	}

	// ID has priority over key
	featureIdOrKey := input.FeatureID
	if featureIdOrKey == nil {
		featureIdOrKey = input.FeatureKey
	}
	if featureIdOrKey == nil {
		return nil, &models.GenericUserError{Message: "Feature ID or Key is required"}
	}

	feature, err := c.featureConnector.GetFeature(ctx, input.Namespace, *featureIdOrKey, productcatalog.IncludeArchivedFeatureFalse)
	if err != nil || feature == nil {
		return nil, &productcatalog.FeatureNotFoundError{ID: *featureIdOrKey}
	}

	if feature.ID != oldEnt.FeatureID {
		return nil, &models.GenericUserError{Message: "Feature in path and body do not match"}
	}

	// Do the override in TX
	return entutils.StartAndRunTx(ctx, c.entitlementRepo, func(ctx context.Context, tx *entutils.TxDriver) (*Entitlement, error) {
		ctx = entutils.NewTxContext(ctx, tx)

		// Delete previous entitlement
		// FIXME: we publish an event during this even if we fail later
		err := c.DeleteEntitlement(ctx, input.Namespace, oldEnt.ID)
		if err != nil {
			return nil, err
		}

		// Create new entitlement
		return c.CreateEntitlement(ctx, input)
	})
}

func (c *entitlementConnector) GetEntitlement(ctx context.Context, namespace string, id string) (*Entitlement, error) {
	return c.entitlementRepo.GetEntitlement(ctx, models.NamespacedID{Namespace: namespace, ID: id})
}

func (c *entitlementConnector) DeleteEntitlement(ctx context.Context, namespace string, id string) error {
	doInTx := func(ctx context.Context, tx *entutils.TxDriver) (*Entitlement, error) {
		ent, err := c.entitlementRepo.WithTx(ctx, tx).GetEntitlement(ctx, models.NamespacedID{Namespace: namespace, ID: id})
		if err != nil {
			return nil, err
		}

		err = c.entitlementRepo.WithTx(ctx, tx).DeleteEntitlement(ctx, models.NamespacedID{Namespace: namespace, ID: id})
		if err != nil {
			return nil, err
		}

		err = c.publisher.Publish(ctx, EntitlementDeletedEvent{
			Entitlement: *ent,
			Namespace: eventmodels.NamespaceID{
				ID: namespace,
			},
		})
		if err != nil {
			return nil, err
		}

		return ent, nil
	}

	if ctxTx, err := entutils.GetTxDriver(ctx); err == nil {
		// we're already in a tx
		_, err := doInTx(ctx, ctxTx)
		return err
	} else {
		_, err := entutils.StartAndRunTx(ctx, c.entitlementRepo, doInTx)
		return err
	}
}

func (c *entitlementConnector) GetEntitlementsOfSubject(ctx context.Context, namespace string, subjectKey models.SubjectKey) ([]Entitlement, error) {
	return c.entitlementRepo.GetEntitlementsOfSubject(ctx, namespace, subjectKey)
}

func (c *entitlementConnector) GetEntitlementValue(ctx context.Context, namespace string, subjectKey string, idOrFeatureKey string, at time.Time) (EntitlementValue, error) {
	ent, err := c.entitlementRepo.GetEntitlementOfSubject(ctx, namespace, subjectKey, idOrFeatureKey)
	if err != nil {
		return nil, err
	}
	connector, err := c.getTypeConnector(ent)
	if err != nil {
		return nil, err
	}
	return connector.GetValue(ent, at)
}

func (c *entitlementConnector) ListEntitlements(ctx context.Context, params ListEntitlementsParams) (pagination.PagedResponse[Entitlement], error) {
	if !params.Page.IsZero() {
		if err := params.Page.Validate(); err != nil {
			return pagination.PagedResponse[Entitlement]{}, err
		}
	}
	return c.entitlementRepo.ListEntitlements(ctx, params)
}

func (c *entitlementConnector) getTypeConnector(inp TypedEntitlement) (SubTypeConnector, error) {
	entitlementType := inp.GetType()
	switch entitlementType {
	case EntitlementTypeMetered:
		return c.meteredEntitlementConnector, nil
	case EntitlementTypeStatic:
		return c.staticEntitlementConnector, nil
	case EntitlementTypeBoolean:
		return c.booleanEntitlementConnector, nil
	default:
		return nil, fmt.Errorf("unsupported entitlement type: %s", entitlementType)
	}
}
