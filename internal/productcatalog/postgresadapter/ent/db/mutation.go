// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/openmeterio/openmeter/internal/productcatalog/postgresadapter/ent/db/feature"
	"github.com/openmeterio/openmeter/internal/productcatalog/postgresadapter/ent/db/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFeature = "Feature"
)

// FeatureMutation represents an operation that mutates the Feature nodes in the graph.
type FeatureMutation struct {
	config
	op                     Op
	typ                    string
	id                     *string
	created_at             *time.Time
	updated_at             *time.Time
	deleted_at             *time.Time
	metadata               *map[string]string
	namespace              *string
	name                   *string
	key                    *string
	meter_slug             *string
	meter_group_by_filters *map[string]string
	archived_at            *time.Time
	clearedFields          map[string]struct{}
	done                   bool
	oldValue               func(context.Context) (*Feature, error)
	predicates             []predicate.Feature
}

var _ ent.Mutation = (*FeatureMutation)(nil)

// featureOption allows management of the mutation configuration using functional options.
type featureOption func(*FeatureMutation)

// newFeatureMutation creates new mutation for the Feature entity.
func newFeatureMutation(c config, op Op, opts ...featureOption) *FeatureMutation {
	m := &FeatureMutation{
		config:        c,
		op:            op,
		typ:           TypeFeature,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFeatureID sets the ID field of the mutation.
func withFeatureID(id string) featureOption {
	return func(m *FeatureMutation) {
		var (
			err   error
			once  sync.Once
			value *Feature
		)
		m.oldValue = func(ctx context.Context) (*Feature, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Feature.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFeature sets the old Feature of the mutation.
func withFeature(node *Feature) featureOption {
	return func(m *FeatureMutation) {
		m.oldValue = func(context.Context) (*Feature, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FeatureMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FeatureMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("db: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Feature entities.
func (m *FeatureMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *FeatureMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *FeatureMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Feature.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *FeatureMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *FeatureMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Feature entity.
// If the Feature object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeatureMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *FeatureMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *FeatureMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *FeatureMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Feature entity.
// If the Feature object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeatureMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *FeatureMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetDeletedAt sets the "deleted_at" field.
func (m *FeatureMutation) SetDeletedAt(t time.Time) {
	m.deleted_at = &t
}

// DeletedAt returns the value of the "deleted_at" field in the mutation.
func (m *FeatureMutation) DeletedAt() (r time.Time, exists bool) {
	v := m.deleted_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeletedAt returns the old "deleted_at" field's value of the Feature entity.
// If the Feature object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeatureMutation) OldDeletedAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeletedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeletedAt: %w", err)
	}
	return oldValue.DeletedAt, nil
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (m *FeatureMutation) ClearDeletedAt() {
	m.deleted_at = nil
	m.clearedFields[feature.FieldDeletedAt] = struct{}{}
}

// DeletedAtCleared returns if the "deleted_at" field was cleared in this mutation.
func (m *FeatureMutation) DeletedAtCleared() bool {
	_, ok := m.clearedFields[feature.FieldDeletedAt]
	return ok
}

// ResetDeletedAt resets all changes to the "deleted_at" field.
func (m *FeatureMutation) ResetDeletedAt() {
	m.deleted_at = nil
	delete(m.clearedFields, feature.FieldDeletedAt)
}

// SetMetadata sets the "metadata" field.
func (m *FeatureMutation) SetMetadata(value map[string]string) {
	m.metadata = &value
}

// Metadata returns the value of the "metadata" field in the mutation.
func (m *FeatureMutation) Metadata() (r map[string]string, exists bool) {
	v := m.metadata
	if v == nil {
		return
	}
	return *v, true
}

// OldMetadata returns the old "metadata" field's value of the Feature entity.
// If the Feature object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeatureMutation) OldMetadata(ctx context.Context) (v map[string]string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMetadata is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMetadata requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMetadata: %w", err)
	}
	return oldValue.Metadata, nil
}

// ClearMetadata clears the value of the "metadata" field.
func (m *FeatureMutation) ClearMetadata() {
	m.metadata = nil
	m.clearedFields[feature.FieldMetadata] = struct{}{}
}

// MetadataCleared returns if the "metadata" field was cleared in this mutation.
func (m *FeatureMutation) MetadataCleared() bool {
	_, ok := m.clearedFields[feature.FieldMetadata]
	return ok
}

// ResetMetadata resets all changes to the "metadata" field.
func (m *FeatureMutation) ResetMetadata() {
	m.metadata = nil
	delete(m.clearedFields, feature.FieldMetadata)
}

// SetNamespace sets the "namespace" field.
func (m *FeatureMutation) SetNamespace(s string) {
	m.namespace = &s
}

// Namespace returns the value of the "namespace" field in the mutation.
func (m *FeatureMutation) Namespace() (r string, exists bool) {
	v := m.namespace
	if v == nil {
		return
	}
	return *v, true
}

// OldNamespace returns the old "namespace" field's value of the Feature entity.
// If the Feature object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeatureMutation) OldNamespace(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNamespace is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNamespace requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNamespace: %w", err)
	}
	return oldValue.Namespace, nil
}

// ResetNamespace resets all changes to the "namespace" field.
func (m *FeatureMutation) ResetNamespace() {
	m.namespace = nil
}

// SetName sets the "name" field.
func (m *FeatureMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *FeatureMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Feature entity.
// If the Feature object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeatureMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *FeatureMutation) ResetName() {
	m.name = nil
}

// SetKey sets the "key" field.
func (m *FeatureMutation) SetKey(s string) {
	m.key = &s
}

// Key returns the value of the "key" field in the mutation.
func (m *FeatureMutation) Key() (r string, exists bool) {
	v := m.key
	if v == nil {
		return
	}
	return *v, true
}

// OldKey returns the old "key" field's value of the Feature entity.
// If the Feature object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeatureMutation) OldKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldKey: %w", err)
	}
	return oldValue.Key, nil
}

// ResetKey resets all changes to the "key" field.
func (m *FeatureMutation) ResetKey() {
	m.key = nil
}

// SetMeterSlug sets the "meter_slug" field.
func (m *FeatureMutation) SetMeterSlug(s string) {
	m.meter_slug = &s
}

// MeterSlug returns the value of the "meter_slug" field in the mutation.
func (m *FeatureMutation) MeterSlug() (r string, exists bool) {
	v := m.meter_slug
	if v == nil {
		return
	}
	return *v, true
}

// OldMeterSlug returns the old "meter_slug" field's value of the Feature entity.
// If the Feature object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeatureMutation) OldMeterSlug(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMeterSlug is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMeterSlug requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMeterSlug: %w", err)
	}
	return oldValue.MeterSlug, nil
}

// ClearMeterSlug clears the value of the "meter_slug" field.
func (m *FeatureMutation) ClearMeterSlug() {
	m.meter_slug = nil
	m.clearedFields[feature.FieldMeterSlug] = struct{}{}
}

// MeterSlugCleared returns if the "meter_slug" field was cleared in this mutation.
func (m *FeatureMutation) MeterSlugCleared() bool {
	_, ok := m.clearedFields[feature.FieldMeterSlug]
	return ok
}

// ResetMeterSlug resets all changes to the "meter_slug" field.
func (m *FeatureMutation) ResetMeterSlug() {
	m.meter_slug = nil
	delete(m.clearedFields, feature.FieldMeterSlug)
}

// SetMeterGroupByFilters sets the "meter_group_by_filters" field.
func (m *FeatureMutation) SetMeterGroupByFilters(value map[string]string) {
	m.meter_group_by_filters = &value
}

// MeterGroupByFilters returns the value of the "meter_group_by_filters" field in the mutation.
func (m *FeatureMutation) MeterGroupByFilters() (r map[string]string, exists bool) {
	v := m.meter_group_by_filters
	if v == nil {
		return
	}
	return *v, true
}

// OldMeterGroupByFilters returns the old "meter_group_by_filters" field's value of the Feature entity.
// If the Feature object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeatureMutation) OldMeterGroupByFilters(ctx context.Context) (v map[string]string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMeterGroupByFilters is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMeterGroupByFilters requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMeterGroupByFilters: %w", err)
	}
	return oldValue.MeterGroupByFilters, nil
}

// ClearMeterGroupByFilters clears the value of the "meter_group_by_filters" field.
func (m *FeatureMutation) ClearMeterGroupByFilters() {
	m.meter_group_by_filters = nil
	m.clearedFields[feature.FieldMeterGroupByFilters] = struct{}{}
}

// MeterGroupByFiltersCleared returns if the "meter_group_by_filters" field was cleared in this mutation.
func (m *FeatureMutation) MeterGroupByFiltersCleared() bool {
	_, ok := m.clearedFields[feature.FieldMeterGroupByFilters]
	return ok
}

// ResetMeterGroupByFilters resets all changes to the "meter_group_by_filters" field.
func (m *FeatureMutation) ResetMeterGroupByFilters() {
	m.meter_group_by_filters = nil
	delete(m.clearedFields, feature.FieldMeterGroupByFilters)
}

// SetArchivedAt sets the "archived_at" field.
func (m *FeatureMutation) SetArchivedAt(t time.Time) {
	m.archived_at = &t
}

// ArchivedAt returns the value of the "archived_at" field in the mutation.
func (m *FeatureMutation) ArchivedAt() (r time.Time, exists bool) {
	v := m.archived_at
	if v == nil {
		return
	}
	return *v, true
}

// OldArchivedAt returns the old "archived_at" field's value of the Feature entity.
// If the Feature object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FeatureMutation) OldArchivedAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldArchivedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldArchivedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldArchivedAt: %w", err)
	}
	return oldValue.ArchivedAt, nil
}

// ClearArchivedAt clears the value of the "archived_at" field.
func (m *FeatureMutation) ClearArchivedAt() {
	m.archived_at = nil
	m.clearedFields[feature.FieldArchivedAt] = struct{}{}
}

// ArchivedAtCleared returns if the "archived_at" field was cleared in this mutation.
func (m *FeatureMutation) ArchivedAtCleared() bool {
	_, ok := m.clearedFields[feature.FieldArchivedAt]
	return ok
}

// ResetArchivedAt resets all changes to the "archived_at" field.
func (m *FeatureMutation) ResetArchivedAt() {
	m.archived_at = nil
	delete(m.clearedFields, feature.FieldArchivedAt)
}

// Where appends a list predicates to the FeatureMutation builder.
func (m *FeatureMutation) Where(ps ...predicate.Feature) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the FeatureMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *FeatureMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Feature, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *FeatureMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *FeatureMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Feature).
func (m *FeatureMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *FeatureMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.created_at != nil {
		fields = append(fields, feature.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, feature.FieldUpdatedAt)
	}
	if m.deleted_at != nil {
		fields = append(fields, feature.FieldDeletedAt)
	}
	if m.metadata != nil {
		fields = append(fields, feature.FieldMetadata)
	}
	if m.namespace != nil {
		fields = append(fields, feature.FieldNamespace)
	}
	if m.name != nil {
		fields = append(fields, feature.FieldName)
	}
	if m.key != nil {
		fields = append(fields, feature.FieldKey)
	}
	if m.meter_slug != nil {
		fields = append(fields, feature.FieldMeterSlug)
	}
	if m.meter_group_by_filters != nil {
		fields = append(fields, feature.FieldMeterGroupByFilters)
	}
	if m.archived_at != nil {
		fields = append(fields, feature.FieldArchivedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *FeatureMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case feature.FieldCreatedAt:
		return m.CreatedAt()
	case feature.FieldUpdatedAt:
		return m.UpdatedAt()
	case feature.FieldDeletedAt:
		return m.DeletedAt()
	case feature.FieldMetadata:
		return m.Metadata()
	case feature.FieldNamespace:
		return m.Namespace()
	case feature.FieldName:
		return m.Name()
	case feature.FieldKey:
		return m.Key()
	case feature.FieldMeterSlug:
		return m.MeterSlug()
	case feature.FieldMeterGroupByFilters:
		return m.MeterGroupByFilters()
	case feature.FieldArchivedAt:
		return m.ArchivedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *FeatureMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case feature.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case feature.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case feature.FieldDeletedAt:
		return m.OldDeletedAt(ctx)
	case feature.FieldMetadata:
		return m.OldMetadata(ctx)
	case feature.FieldNamespace:
		return m.OldNamespace(ctx)
	case feature.FieldName:
		return m.OldName(ctx)
	case feature.FieldKey:
		return m.OldKey(ctx)
	case feature.FieldMeterSlug:
		return m.OldMeterSlug(ctx)
	case feature.FieldMeterGroupByFilters:
		return m.OldMeterGroupByFilters(ctx)
	case feature.FieldArchivedAt:
		return m.OldArchivedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Feature field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FeatureMutation) SetField(name string, value ent.Value) error {
	switch name {
	case feature.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case feature.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case feature.FieldDeletedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeletedAt(v)
		return nil
	case feature.FieldMetadata:
		v, ok := value.(map[string]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMetadata(v)
		return nil
	case feature.FieldNamespace:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNamespace(v)
		return nil
	case feature.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case feature.FieldKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKey(v)
		return nil
	case feature.FieldMeterSlug:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMeterSlug(v)
		return nil
	case feature.FieldMeterGroupByFilters:
		v, ok := value.(map[string]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMeterGroupByFilters(v)
		return nil
	case feature.FieldArchivedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetArchivedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Feature field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *FeatureMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *FeatureMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FeatureMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Feature numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *FeatureMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(feature.FieldDeletedAt) {
		fields = append(fields, feature.FieldDeletedAt)
	}
	if m.FieldCleared(feature.FieldMetadata) {
		fields = append(fields, feature.FieldMetadata)
	}
	if m.FieldCleared(feature.FieldMeterSlug) {
		fields = append(fields, feature.FieldMeterSlug)
	}
	if m.FieldCleared(feature.FieldMeterGroupByFilters) {
		fields = append(fields, feature.FieldMeterGroupByFilters)
	}
	if m.FieldCleared(feature.FieldArchivedAt) {
		fields = append(fields, feature.FieldArchivedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *FeatureMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *FeatureMutation) ClearField(name string) error {
	switch name {
	case feature.FieldDeletedAt:
		m.ClearDeletedAt()
		return nil
	case feature.FieldMetadata:
		m.ClearMetadata()
		return nil
	case feature.FieldMeterSlug:
		m.ClearMeterSlug()
		return nil
	case feature.FieldMeterGroupByFilters:
		m.ClearMeterGroupByFilters()
		return nil
	case feature.FieldArchivedAt:
		m.ClearArchivedAt()
		return nil
	}
	return fmt.Errorf("unknown Feature nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *FeatureMutation) ResetField(name string) error {
	switch name {
	case feature.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case feature.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case feature.FieldDeletedAt:
		m.ResetDeletedAt()
		return nil
	case feature.FieldMetadata:
		m.ResetMetadata()
		return nil
	case feature.FieldNamespace:
		m.ResetNamespace()
		return nil
	case feature.FieldName:
		m.ResetName()
		return nil
	case feature.FieldKey:
		m.ResetKey()
		return nil
	case feature.FieldMeterSlug:
		m.ResetMeterSlug()
		return nil
	case feature.FieldMeterGroupByFilters:
		m.ResetMeterGroupByFilters()
		return nil
	case feature.FieldArchivedAt:
		m.ResetArchivedAt()
		return nil
	}
	return fmt.Errorf("unknown Feature field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *FeatureMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *FeatureMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *FeatureMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *FeatureMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *FeatureMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *FeatureMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *FeatureMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Feature unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *FeatureMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Feature edge %s", name)
}
