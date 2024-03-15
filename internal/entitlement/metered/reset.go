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

package meteredentitlement

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/openmeterio/openmeter/internal/credit"
	"github.com/openmeterio/openmeter/internal/entitlement"
	eventmodels "github.com/openmeterio/openmeter/internal/event/models"
	"github.com/openmeterio/openmeter/internal/event/spec"
	"github.com/openmeterio/openmeter/pkg/framework/entutils"
	"github.com/openmeterio/openmeter/pkg/models"
)

func (e *connector) ResetEntitlementUsage(ctx context.Context, entitlementID models.NamespacedID, params ResetEntitlementUsageParams) (*EntitlementBalance, error) {
	return entutils.StartAndRunTx(ctx, e.entitlementRepo, func(ctx context.Context, tx *entutils.TxDriver) (*EntitlementBalance, error) {
		txCtx := entutils.NewTxContext(ctx, tx)

		owner := credit.NamespacedGrantOwner{
			Namespace: entitlementID.Namespace,
			ID:        credit.GrantOwner(entitlementID.ID),
		}

		ent, err := e.entitlementRepo.WithTx(txCtx, tx).GetEntitlement(txCtx, entitlementID)
		if err != nil {
			return nil, fmt.Errorf("failed to get entitlement: %w", err)
		}

		_, err = ParseFromGenericEntitlement(ent)
		if err != nil {
			return nil, fmt.Errorf("failed to parse entitlement: %w", err)
		}

		balanceAfterReset, err := e.balanceConnector.ResetUsageForOwner(txCtx, owner, credit.ResetUsageForOwnerParams{
			At:           params.At,
			RetainAnchor: params.RetainAnchor,
		})
		if err != nil {
			if _, ok := err.(*credit.OwnerNotFoundError); ok {
				return nil, &entitlement.NotFoundError{EntitlementID: entitlementID}
			}
			return nil, err
		}

		event, err := spec.NewCloudEvent(
			spec.EventSpec{
				Source:  spec.ComposeResourcePath(entitlementID.Namespace, spec.EntityEntitlement, entitlementID.ID),
				Subject: spec.ComposeResourcePath(entitlementID.Namespace, spec.EntitySubjectKey, ent.SubjectKey),
			},
			EntitlementResetEvent{
				EntitlementID: entitlementID.ID,
				Namespace: eventmodels.NamespaceID{
					ID: entitlementID.Namespace,
				},
				Subject: eventmodels.SubjectKeyAndID{
					Key: ent.SubjectKey,
				},
				ResetAt:      params.At,
				RetainAnchor: params.RetainAnchor,
			},
		)
		if err != nil {
			return nil, err
		}

		if err := e.publisher.Publish(event); err != nil {
			return nil, err
		}

		return &EntitlementBalance{
			EntitlementID: entitlementID.ID,
			Balance:       balanceAfterReset.Balance(),
			UsageInPeriod: 0.0, // you cannot have usage right after a reset
			Overage:       balanceAfterReset.Overage,
			StartOfPeriod: params.At,
		}, nil
	})
}

func (c *connector) ResetEntitlementsWithExpiredUsagePeriod(ctx context.Context, namespace string, highwatermark time.Time) ([]models.NamespacedID, error) {
	entitlements, err := c.entitlementRepo.ListEntitlementsWithExpiredUsagePeriod(ctx, namespace, highwatermark)
	if err != nil {
		return nil, fmt.Errorf("failed to list entitlements with due reset: %w", err)
	}

	result := make([]models.NamespacedID, 0, len(entitlements))

	var finalError error
	for _, ent := range entitlements {
		namespacedID := models.NamespacedID{Namespace: namespace, ID: ent.ID}

		_, err := c.ResetEntitlementUsage(ctx,
			namespacedID,
			ResetEntitlementUsageParams{
				At:           ent.CurrentUsagePeriod.To,
				RetainAnchor: true,
			})
		if err != nil {
			finalError = errors.Join(finalError, fmt.Errorf("failed to reset entitlement usage ns=%s id=%s: %w", namespace, ent.ID, err))
		}

		result = append(result, namespacedID)
	}
	return result, finalError
}
