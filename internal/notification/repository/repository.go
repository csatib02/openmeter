package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	entdb "github.com/openmeterio/openmeter/internal/ent/db"
	channeldb "github.com/openmeterio/openmeter/internal/ent/db/notificationchannel"
	"github.com/openmeterio/openmeter/internal/notification"
	"github.com/openmeterio/openmeter/pkg/clock"
	"github.com/openmeterio/openmeter/pkg/framework/entutils"
	"github.com/openmeterio/openmeter/pkg/models"
	"github.com/openmeterio/openmeter/pkg/pagination"
	"github.com/openmeterio/openmeter/pkg/sortx"
)

type Config struct {
	Client *entdb.Client
	Logger *slog.Logger
}

func (c Config) Validate() error {
	if c.Client == nil {
		return errors.New("postgres client is required")
	}

	if c.Logger == nil {
		return errors.New("logger must not be nil")
	}

	return nil
}

func New(config Config) (notification.Repository, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &repository{
		db:     config.Client,
		logger: config.Logger,
	}, nil
}

var _ notification.Repository = (*repository)(nil)

type repository struct {
	db *entdb.Client

	logger *slog.Logger
}

func (r repository) ListChannels(ctx context.Context, params notification.ListChannelsInput) (pagination.PagedResponse[notification.Channel], error) {
	query := r.db.NotificationChannel.Query().
		Where(channeldb.DeletedAtIsNil()) // Do not return deleted channels

	if len(params.Namespaces) > 0 {
		query = query.Where(channeldb.NamespaceIn(params.Namespaces...))
	}

	if len(params.Channels) > 0 {
		query = query.Where(channeldb.IDIn(params.Channels...))
	}

	if !params.IncludeDisabled {
		query = query.Where(channeldb.Disabled(false))
	}

	order := entutils.GetOrdering(sortx.OrderDefault)
	if !params.Order.IsDefaultValue() {
		order = entutils.GetOrdering(params.Order)
	}

	switch params.OrderBy {
	case notification.ChannelOrderByCreatedAt:
		query = query.Order(channeldb.ByCreatedAt(order...))
	case notification.ChannelOrderByUpdatedAt:
		query = query.Order(channeldb.ByUpdatedAt(order...))
	case notification.ChannelOrderByType:
		query = query.Order(channeldb.ByType(order...))
	case notification.ChannelOrderByID:
		fallthrough
	default:
		query = query.Order(channeldb.ByID(order...))
	}

	response := pagination.PagedResponse[notification.Channel]{
		Page: params.Page,
	}

	paged, err := query.Paginate(ctx, params.Page)
	if err != nil {
		return response, err
	}

	result := make([]notification.Channel, 0, len(paged.Items))
	for _, item := range paged.Items {
		if item == nil {
			r.logger.Warn("invalid query result: nil notification channel received")
			continue
		}

		result = append(result, *ChannelFromDBEntity(*item))
	}

	response.TotalCount = paged.TotalCount
	response.Items = result

	return response, nil
}

func (r repository) CreateChannel(ctx context.Context, params notification.CreateChannelInput) (*notification.Channel, error) {
	query := r.db.NotificationChannel.Create().
		SetType(params.Type).
		SetName(params.Name).
		SetNamespace(params.Namespace).
		SetDisabled(params.Disabled).
		SetConfig(params.Config)

	channel, err := query.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create notification channel: %w", err)
	}

	if channel == nil {
		return nil, fmt.Errorf("invalid query result: nil notification channel received")
	}

	return ChannelFromDBEntity(*channel), nil
}

func (r repository) DeleteChannel(ctx context.Context, params notification.DeleteChannelInput) error {
	query := r.db.NotificationChannel.UpdateOneID(params.ID).
		SetDeletedAt(clock.Now().UTC()).
		SetDisabled(true)

	_, err := query.Save(ctx)
	if err != nil {
		if entdb.IsNotFound(err) {
			return notification.NotFoundError{
				NamespacedID: models.NamespacedID{
					Namespace: params.Namespace,
					ID:        params.ID,
				},
			}
		}

		return fmt.Errorf("failed to delete notification channel: %w", err)
	}

	return nil
}

func (r repository) GetChannel(ctx context.Context, params notification.GetChannelInput) (*notification.Channel, error) {
	query := r.db.NotificationChannel.Query().
		Where(channeldb.ID(params.ID)).
		Where(channeldb.Namespace(params.Namespace))

	queryRow, err := query.First(ctx)
	if err != nil {
		if entdb.IsNotFound(err) {
			return nil, notification.NotFoundError{
				NamespacedID: models.NamespacedID{
					Namespace: params.Namespace,
					ID:        params.ID,
				},
			}
		}

		return nil, fmt.Errorf("failed to fetch notification channel: %w", err)
	}

	if queryRow == nil {
		return nil, fmt.Errorf("invalid query result: nil notification channel received")
	}

	return ChannelFromDBEntity(*queryRow), nil
}

func (r repository) UpdateChannel(ctx context.Context, params notification.UpdateChannelInput) (*notification.Channel, error) {
	query := r.db.NotificationChannel.UpdateOneID(params.ID).
		SetUpdatedAt(clock.Now().UTC()).
		SetDisabled(params.Disabled).
		SetConfig(params.Config).
		SetName(params.Name)

	queryRow, err := query.Save(ctx)
	if err != nil {
		if entdb.IsNotFound(err) {
			return nil, notification.NotFoundError{
				NamespacedID: models.NamespacedID{
					Namespace: params.Namespace,
					ID:        params.ID,
				},
			}
		}

		return nil, fmt.Errorf("failed to update notification channel: %w", err)
	}

	if queryRow == nil {
		return nil, fmt.Errorf("invalid query result: nil notification channel received")
	}

	return ChannelFromDBEntity(*queryRow), nil
}