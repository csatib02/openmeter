// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/creditentry"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/predicate"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/pgulid"
)

// CreditEntryUpdate is the builder for updating CreditEntry entities.
type CreditEntryUpdate struct {
	config
	hooks    []Hook
	mutation *CreditEntryMutation
}

// Where appends a list predicates to the CreditEntryUpdate builder.
func (ceu *CreditEntryUpdate) Where(ps ...predicate.CreditEntry) *CreditEntryUpdate {
	ceu.mutation.Where(ps...)
	return ceu
}

// SetUpdatedAt sets the "updated_at" field.
func (ceu *CreditEntryUpdate) SetUpdatedAt(t time.Time) *CreditEntryUpdate {
	ceu.mutation.SetUpdatedAt(t)
	return ceu
}

// SetMetadata sets the "metadata" field.
func (ceu *CreditEntryUpdate) SetMetadata(m map[string]string) *CreditEntryUpdate {
	ceu.mutation.SetMetadata(m)
	return ceu
}

// ClearMetadata clears the value of the "metadata" field.
func (ceu *CreditEntryUpdate) ClearMetadata() *CreditEntryUpdate {
	ceu.mutation.ClearMetadata()
	return ceu
}

// SetChildrenID sets the "children" edge to the CreditEntry entity by ID.
func (ceu *CreditEntryUpdate) SetChildrenID(id pgulid.ULID) *CreditEntryUpdate {
	ceu.mutation.SetChildrenID(id)
	return ceu
}

// SetNillableChildrenID sets the "children" edge to the CreditEntry entity by ID if the given value is not nil.
func (ceu *CreditEntryUpdate) SetNillableChildrenID(id *pgulid.ULID) *CreditEntryUpdate {
	if id != nil {
		ceu = ceu.SetChildrenID(*id)
	}
	return ceu
}

// SetChildren sets the "children" edge to the CreditEntry entity.
func (ceu *CreditEntryUpdate) SetChildren(c *CreditEntry) *CreditEntryUpdate {
	return ceu.SetChildrenID(c.ID)
}

// Mutation returns the CreditEntryMutation object of the builder.
func (ceu *CreditEntryUpdate) Mutation() *CreditEntryMutation {
	return ceu.mutation
}

// ClearChildren clears the "children" edge to the CreditEntry entity.
func (ceu *CreditEntryUpdate) ClearChildren() *CreditEntryUpdate {
	ceu.mutation.ClearChildren()
	return ceu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ceu *CreditEntryUpdate) Save(ctx context.Context) (int, error) {
	ceu.defaults()
	return withHooks(ctx, ceu.sqlSave, ceu.mutation, ceu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ceu *CreditEntryUpdate) SaveX(ctx context.Context) int {
	affected, err := ceu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ceu *CreditEntryUpdate) Exec(ctx context.Context) error {
	_, err := ceu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceu *CreditEntryUpdate) ExecX(ctx context.Context) {
	if err := ceu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ceu *CreditEntryUpdate) defaults() {
	if _, ok := ceu.mutation.UpdatedAt(); !ok {
		v := creditentry.UpdateDefaultUpdatedAt()
		ceu.mutation.SetUpdatedAt(v)
	}
}

func (ceu *CreditEntryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(creditentry.Table, creditentry.Columns, sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther))
	if ps := ceu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceu.mutation.UpdatedAt(); ok {
		_spec.SetField(creditentry.FieldUpdatedAt, field.TypeTime, value)
	}
	if ceu.mutation.TypeCleared() {
		_spec.ClearField(creditentry.FieldType, field.TypeEnum)
	}
	if ceu.mutation.AmountCleared() {
		_spec.ClearField(creditentry.FieldAmount, field.TypeFloat64)
	}
	if ceu.mutation.ExpirationPeriodDurationCleared() {
		_spec.ClearField(creditentry.FieldExpirationPeriodDuration, field.TypeEnum)
	}
	if ceu.mutation.ExpirationPeriodCountCleared() {
		_spec.ClearField(creditentry.FieldExpirationPeriodCount, field.TypeUint8)
	}
	if ceu.mutation.ExpirationAtCleared() {
		_spec.ClearField(creditentry.FieldExpirationAt, field.TypeTime)
	}
	if ceu.mutation.RolloverTypeCleared() {
		_spec.ClearField(creditentry.FieldRolloverType, field.TypeEnum)
	}
	if ceu.mutation.RolloverMaxAmountCleared() {
		_spec.ClearField(creditentry.FieldRolloverMaxAmount, field.TypeFloat64)
	}
	if value, ok := ceu.mutation.Metadata(); ok {
		_spec.SetField(creditentry.FieldMetadata, field.TypeJSON, value)
	}
	if ceu.mutation.MetadataCleared() {
		_spec.ClearField(creditentry.FieldMetadata, field.TypeJSON)
	}
	if ceu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   creditentry.ChildrenTable,
			Columns: []string{creditentry.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   creditentry.ChildrenTable,
			Columns: []string{creditentry.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ceu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{creditentry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ceu.mutation.done = true
	return n, nil
}

// CreditEntryUpdateOne is the builder for updating a single CreditEntry entity.
type CreditEntryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CreditEntryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ceuo *CreditEntryUpdateOne) SetUpdatedAt(t time.Time) *CreditEntryUpdateOne {
	ceuo.mutation.SetUpdatedAt(t)
	return ceuo
}

// SetMetadata sets the "metadata" field.
func (ceuo *CreditEntryUpdateOne) SetMetadata(m map[string]string) *CreditEntryUpdateOne {
	ceuo.mutation.SetMetadata(m)
	return ceuo
}

// ClearMetadata clears the value of the "metadata" field.
func (ceuo *CreditEntryUpdateOne) ClearMetadata() *CreditEntryUpdateOne {
	ceuo.mutation.ClearMetadata()
	return ceuo
}

// SetChildrenID sets the "children" edge to the CreditEntry entity by ID.
func (ceuo *CreditEntryUpdateOne) SetChildrenID(id pgulid.ULID) *CreditEntryUpdateOne {
	ceuo.mutation.SetChildrenID(id)
	return ceuo
}

// SetNillableChildrenID sets the "children" edge to the CreditEntry entity by ID if the given value is not nil.
func (ceuo *CreditEntryUpdateOne) SetNillableChildrenID(id *pgulid.ULID) *CreditEntryUpdateOne {
	if id != nil {
		ceuo = ceuo.SetChildrenID(*id)
	}
	return ceuo
}

// SetChildren sets the "children" edge to the CreditEntry entity.
func (ceuo *CreditEntryUpdateOne) SetChildren(c *CreditEntry) *CreditEntryUpdateOne {
	return ceuo.SetChildrenID(c.ID)
}

// Mutation returns the CreditEntryMutation object of the builder.
func (ceuo *CreditEntryUpdateOne) Mutation() *CreditEntryMutation {
	return ceuo.mutation
}

// ClearChildren clears the "children" edge to the CreditEntry entity.
func (ceuo *CreditEntryUpdateOne) ClearChildren() *CreditEntryUpdateOne {
	ceuo.mutation.ClearChildren()
	return ceuo
}

// Where appends a list predicates to the CreditEntryUpdate builder.
func (ceuo *CreditEntryUpdateOne) Where(ps ...predicate.CreditEntry) *CreditEntryUpdateOne {
	ceuo.mutation.Where(ps...)
	return ceuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ceuo *CreditEntryUpdateOne) Select(field string, fields ...string) *CreditEntryUpdateOne {
	ceuo.fields = append([]string{field}, fields...)
	return ceuo
}

// Save executes the query and returns the updated CreditEntry entity.
func (ceuo *CreditEntryUpdateOne) Save(ctx context.Context) (*CreditEntry, error) {
	ceuo.defaults()
	return withHooks(ctx, ceuo.sqlSave, ceuo.mutation, ceuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ceuo *CreditEntryUpdateOne) SaveX(ctx context.Context) *CreditEntry {
	node, err := ceuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ceuo *CreditEntryUpdateOne) Exec(ctx context.Context) error {
	_, err := ceuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceuo *CreditEntryUpdateOne) ExecX(ctx context.Context) {
	if err := ceuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ceuo *CreditEntryUpdateOne) defaults() {
	if _, ok := ceuo.mutation.UpdatedAt(); !ok {
		v := creditentry.UpdateDefaultUpdatedAt()
		ceuo.mutation.SetUpdatedAt(v)
	}
}

func (ceuo *CreditEntryUpdateOne) sqlSave(ctx context.Context) (_node *CreditEntry, err error) {
	_spec := sqlgraph.NewUpdateSpec(creditentry.Table, creditentry.Columns, sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther))
	id, ok := ceuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "CreditEntry.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ceuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, creditentry.FieldID)
		for _, f := range fields {
			if !creditentry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != creditentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ceuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceuo.mutation.UpdatedAt(); ok {
		_spec.SetField(creditentry.FieldUpdatedAt, field.TypeTime, value)
	}
	if ceuo.mutation.TypeCleared() {
		_spec.ClearField(creditentry.FieldType, field.TypeEnum)
	}
	if ceuo.mutation.AmountCleared() {
		_spec.ClearField(creditentry.FieldAmount, field.TypeFloat64)
	}
	if ceuo.mutation.ExpirationPeriodDurationCleared() {
		_spec.ClearField(creditentry.FieldExpirationPeriodDuration, field.TypeEnum)
	}
	if ceuo.mutation.ExpirationPeriodCountCleared() {
		_spec.ClearField(creditentry.FieldExpirationPeriodCount, field.TypeUint8)
	}
	if ceuo.mutation.ExpirationAtCleared() {
		_spec.ClearField(creditentry.FieldExpirationAt, field.TypeTime)
	}
	if ceuo.mutation.RolloverTypeCleared() {
		_spec.ClearField(creditentry.FieldRolloverType, field.TypeEnum)
	}
	if ceuo.mutation.RolloverMaxAmountCleared() {
		_spec.ClearField(creditentry.FieldRolloverMaxAmount, field.TypeFloat64)
	}
	if value, ok := ceuo.mutation.Metadata(); ok {
		_spec.SetField(creditentry.FieldMetadata, field.TypeJSON, value)
	}
	if ceuo.mutation.MetadataCleared() {
		_spec.ClearField(creditentry.FieldMetadata, field.TypeJSON)
	}
	if ceuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   creditentry.ChildrenTable,
			Columns: []string{creditentry.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   creditentry.ChildrenTable,
			Columns: []string{creditentry.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CreditEntry{config: ceuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ceuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{creditentry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ceuo.mutation.done = true
	return _node, nil
}
