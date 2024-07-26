package credit

import (
	"context"
	"fmt"
	"time"

	eventmodels "github.com/openmeterio/openmeter/internal/event/models"
	"github.com/openmeterio/openmeter/internal/event/publisher"
	"github.com/openmeterio/openmeter/internal/event/spec"
	"github.com/openmeterio/openmeter/pkg/clock"
	"github.com/openmeterio/openmeter/pkg/framework/entutils"
	"github.com/openmeterio/openmeter/pkg/models"
	"github.com/openmeterio/openmeter/pkg/pagination"
	"github.com/openmeterio/openmeter/pkg/recurrence"
	"github.com/openmeterio/openmeter/pkg/sortx"
)

type CreateGrantInput struct {
	Amount           float64
	Priority         uint8
	EffectiveAt      time.Time
	Expiration       ExpirationPeriod
	Metadata         map[string]string
	ResetMaxRollover float64
	ResetMinRollover float64
	Recurrence       *recurrence.Recurrence
}

type GrantConnector interface {
	CreateGrant(ctx context.Context, owner NamespacedGrantOwner, grant CreateGrantInput) (*Grant, error)
	VoidGrant(ctx context.Context, grantID models.NamespacedID) error
	ListGrants(ctx context.Context, params ListGrantsParams) (pagination.PagedResponse[Grant], error)
	ListActiveGrantsBetween(ctx context.Context, owner NamespacedGrantOwner, from, to time.Time) ([]Grant, error)
	GetGrant(ctx context.Context, grantID models.NamespacedID) (Grant, error)
}

type GrantOrderBy string

const (
	GrantOrderByCreatedAt   GrantOrderBy = "created_at"
	GrantOrderByUpdatedAt   GrantOrderBy = "updated_at"
	GrantOrderByExpiresAt   GrantOrderBy = "expires_at"
	GrantOrderByEffectiveAt GrantOrderBy = "effective_at"
	GrantOrderByOwner       GrantOrderBy = "owner_id" // check
)

type ListGrantsParams struct {
	Namespace      string
	OwnerID        *GrantOwner
	IncludeDeleted bool
	Page           pagination.Page
	OrderBy        GrantOrderBy
	Order          sortx.Order
	// will be deprecated
	Limit int
	// will be deprecated
	Offset int
}

type GrantRepoCreateGrantInput struct {
	OwnerID          GrantOwner
	Namespace        string
	Amount           float64
	Priority         uint8
	EffectiveAt      time.Time
	Expiration       ExpirationPeriod
	ExpiresAt        time.Time
	Metadata         map[string]string
	ResetMaxRollover float64
	ResetMinRollover float64
	Recurrence       *recurrence.Recurrence
}

type GrantRepo interface {
	CreateGrant(ctx context.Context, grant GrantRepoCreateGrantInput) (*Grant, error)
	VoidGrant(ctx context.Context, grantID models.NamespacedID, at time.Time) error
	// For bw compatibility, if pagination is not provided we
	ListGrants(ctx context.Context, params ListGrantsParams) (pagination.PagedResponse[Grant], error)
	// ListActiveGrantsBetween returns all grants that are active at any point between the given time range.
	ListActiveGrantsBetween(ctx context.Context, owner NamespacedGrantOwner, from, to time.Time) ([]Grant, error)
	GetGrant(ctx context.Context, grantID models.NamespacedID) (Grant, error)

	entutils.TxCreator
	entutils.TxUser[GrantRepo]
}

type grantConnector struct {
	ownerConnector           OwnerConnector
	grantRepo                GrantRepo
	balanceSnapshotConnector BalanceSnapshotRepo
	granularity              time.Duration

	publisher publisher.TopicPublisher
}

func NewGrantConnector(
	ownerConnector OwnerConnector,
	grantRepo GrantRepo,
	balanceSnapshotConnector BalanceSnapshotRepo,
	granularity time.Duration,
	publisher publisher.TopicPublisher,
) GrantConnector {
	return &grantConnector{
		ownerConnector:           ownerConnector,
		grantRepo:                grantRepo,
		balanceSnapshotConnector: balanceSnapshotConnector,
		granularity:              granularity,
		publisher:                publisher,
	}
}

func (m *grantConnector) CreateGrant(ctx context.Context, owner NamespacedGrantOwner, input CreateGrantInput) (*Grant, error) {
	doInTx := func(ctx context.Context, tx *entutils.TxDriver) (*Grant, error) {
		// All metering information is stored in windowSize chunks,
		// so we cannot do accurate calculations unless we follow that same windowing.
		meter, err := m.ownerConnector.GetMeter(ctx, owner)
		if err != nil {
			return nil, err
		}
		granularity := meter.WindowSize.Duration()
		input.EffectiveAt = input.EffectiveAt.Truncate(granularity)
		if input.Recurrence != nil {
			input.Recurrence.Anchor = input.Recurrence.Anchor.Truncate(granularity)
		}
		periodStart, err := m.ownerConnector.GetUsagePeriodStartAt(ctx, owner, clock.Now())
		if err != nil {
			return nil, err
		}

		if input.EffectiveAt.Before(periodStart) {
			return nil, &models.GenericUserError{Message: "grant effective date is before the current usage period"}
		}

		err = m.ownerConnector.LockOwnerForTx(ctx, tx, owner)
		if err != nil {
			return nil, err
		}
		grant, err := m.grantRepo.WithTx(ctx, tx).CreateGrant(ctx, GrantRepoCreateGrantInput{
			OwnerID:          owner.ID,
			Namespace:        owner.Namespace,
			Amount:           input.Amount,
			Priority:         input.Priority,
			EffectiveAt:      input.EffectiveAt,
			Expiration:       input.Expiration,
			ExpiresAt:        input.Expiration.GetExpiration(input.EffectiveAt),
			Metadata:         input.Metadata,
			ResetMaxRollover: input.ResetMaxRollover,
			ResetMinRollover: input.ResetMinRollover,
			Recurrence:       input.Recurrence,
		})
		if err != nil {
			return nil, err
		}

		// invalidate snapshots
		err = m.balanceSnapshotConnector.WithTx(ctx, tx).InvalidateAfter(ctx, owner, grant.EffectiveAt)
		if err != nil {
			return nil, fmt.Errorf("failed to invalidate snapshots after %s: %w", grant.EffectiveAt, err)
		}

		// publish event
		subjectKey, err := m.ownerConnector.GetOwnerSubjectKey(ctx, owner)
		if err != nil {
			return nil, err
		}

		event, err := spec.NewCloudEvent(
			spec.EventSpec{
				Source:  spec.ComposeResourcePath(owner.Namespace, spec.EntityEntitlement, string(owner.ID), spec.EntityGrant, grant.ID),
				Subject: spec.ComposeResourcePath(owner.Namespace, spec.EntitySubjectKey, subjectKey),
			},
			GrantCreatedEvent{
				Grant:     *grant,
				Namespace: eventmodels.NamespaceID{ID: owner.Namespace},
				Subject:   eventmodels.SubjectKeyAndID{Key: subjectKey},
			},
		)
		if err != nil {
			return nil, err
		}

		if err := m.publisher.Publish(event); err != nil {
			return nil, err
		}

		return grant, err
	}

	if ctxTx, err := entutils.GetTxDriver(ctx); err == nil {
		// we're already in a tx
		return doInTx(ctx, ctxTx)
	} else {
		return entutils.StartAndRunTx(ctx, m.grantRepo, doInTx)
	}
}

func (m *grantConnector) VoidGrant(ctx context.Context, grantID models.NamespacedID) error {
	// can we void grants that have been used?
	grant, err := m.grantRepo.GetGrant(ctx, grantID)
	if err != nil {
		return err
	}

	if grant.VoidedAt != nil {
		return &models.GenericUserError{Message: "grant already voided"}
	}

	owner := NamespacedGrantOwner{Namespace: grantID.Namespace, ID: grant.OwnerID}

	_, err = entutils.StartAndRunTx(ctx, m.grantRepo, func(ctx context.Context, tx *entutils.TxDriver) (*interface{}, error) {
		err := m.ownerConnector.LockOwnerForTx(ctx, tx, owner)
		if err != nil {
			return nil, err
		}

		now := clock.Now().Truncate(m.granularity)
		err = m.grantRepo.WithTx(ctx, tx).VoidGrant(ctx, grantID, now)
		if err != nil {
			return nil, err
		}

		err = m.balanceSnapshotConnector.WithTx(ctx, tx).InvalidateAfter(ctx, owner, now)
		if err != nil {
			return nil, fmt.Errorf("failed to invalidate snapshots after %s: %w", now, err)
		}

		// publish an event
		subjectKey, err := m.ownerConnector.GetOwnerSubjectKey(ctx, owner)
		if err != nil {
			return nil, err
		}

		event, err := spec.NewCloudEvent(
			spec.EventSpec{
				Source:  spec.ComposeResourcePath(grantID.Namespace, spec.EntityEntitlement, string(owner.ID), spec.EntityGrant, grantID.ID),
				Subject: spec.ComposeResourcePath(grantID.Namespace, spec.EntitySubjectKey, subjectKey),
			},
			GrantVoidedEvent{
				Grant:     grant,
				Namespace: eventmodels.NamespaceID{ID: owner.Namespace},
				Subject:   eventmodels.SubjectKeyAndID{Key: subjectKey},
			},
		)
		if err != nil {
			return nil, err
		}

		return nil, m.publisher.Publish(event)
	})
	return err
}

func (m *grantConnector) ListGrants(ctx context.Context, params ListGrantsParams) (pagination.PagedResponse[Grant], error) {
	return m.grantRepo.ListGrants(ctx, params)
}

func (m *grantConnector) ListActiveGrantsBetween(ctx context.Context, owner NamespacedGrantOwner, from, to time.Time) ([]Grant, error) {
	return m.grantRepo.ListActiveGrantsBetween(ctx, owner, from, to)
}

func (m *grantConnector) GetGrant(ctx context.Context, grantID models.NamespacedID) (Grant, error) {
	return m.grantRepo.GetGrant(ctx, grantID)
}

type GrantNotFoundError struct {
	GrantID string
}

func (e *GrantNotFoundError) Error() string {
	return fmt.Sprintf("grant not found: %s", e.GrantID)
}
