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
	"github.com/openmeterio/openmeter/internal/ent/db/notificationchannel"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationrule"
	"github.com/openmeterio/openmeter/internal/ent/db/predicate"
	"github.com/openmeterio/openmeter/internal/notification"
)

// NotificationChannelUpdate is the builder for updating NotificationChannel entities.
type NotificationChannelUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationChannelMutation
}

// Where appends a list predicates to the NotificationChannelUpdate builder.
func (ncu *NotificationChannelUpdate) Where(ps ...predicate.NotificationChannel) *NotificationChannelUpdate {
	ncu.mutation.Where(ps...)
	return ncu
}

// SetUpdatedAt sets the "updated_at" field.
func (ncu *NotificationChannelUpdate) SetUpdatedAt(t time.Time) *NotificationChannelUpdate {
	ncu.mutation.SetUpdatedAt(t)
	return ncu
}

// SetDeletedAt sets the "deleted_at" field.
func (ncu *NotificationChannelUpdate) SetDeletedAt(t time.Time) *NotificationChannelUpdate {
	ncu.mutation.SetDeletedAt(t)
	return ncu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ncu *NotificationChannelUpdate) SetNillableDeletedAt(t *time.Time) *NotificationChannelUpdate {
	if t != nil {
		ncu.SetDeletedAt(*t)
	}
	return ncu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ncu *NotificationChannelUpdate) ClearDeletedAt() *NotificationChannelUpdate {
	ncu.mutation.ClearDeletedAt()
	return ncu
}

// SetName sets the "name" field.
func (ncu *NotificationChannelUpdate) SetName(s string) *NotificationChannelUpdate {
	ncu.mutation.SetName(s)
	return ncu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ncu *NotificationChannelUpdate) SetNillableName(s *string) *NotificationChannelUpdate {
	if s != nil {
		ncu.SetName(*s)
	}
	return ncu
}

// SetDisabled sets the "disabled" field.
func (ncu *NotificationChannelUpdate) SetDisabled(b bool) *NotificationChannelUpdate {
	ncu.mutation.SetDisabled(b)
	return ncu
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (ncu *NotificationChannelUpdate) SetNillableDisabled(b *bool) *NotificationChannelUpdate {
	if b != nil {
		ncu.SetDisabled(*b)
	}
	return ncu
}

// ClearDisabled clears the value of the "disabled" field.
func (ncu *NotificationChannelUpdate) ClearDisabled() *NotificationChannelUpdate {
	ncu.mutation.ClearDisabled()
	return ncu
}

// SetConfig sets the "config" field.
func (ncu *NotificationChannelUpdate) SetConfig(nc notification.ChannelConfig) *NotificationChannelUpdate {
	ncu.mutation.SetConfig(nc)
	return ncu
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (ncu *NotificationChannelUpdate) SetNillableConfig(nc *notification.ChannelConfig) *NotificationChannelUpdate {
	if nc != nil {
		ncu.SetConfig(*nc)
	}
	return ncu
}

// AddRuleIDs adds the "rules" edge to the NotificationRule entity by IDs.
func (ncu *NotificationChannelUpdate) AddRuleIDs(ids ...string) *NotificationChannelUpdate {
	ncu.mutation.AddRuleIDs(ids...)
	return ncu
}

// AddRules adds the "rules" edges to the NotificationRule entity.
func (ncu *NotificationChannelUpdate) AddRules(n ...*NotificationRule) *NotificationChannelUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ncu.AddRuleIDs(ids...)
}

// Mutation returns the NotificationChannelMutation object of the builder.
func (ncu *NotificationChannelUpdate) Mutation() *NotificationChannelMutation {
	return ncu.mutation
}

// ClearRules clears all "rules" edges to the NotificationRule entity.
func (ncu *NotificationChannelUpdate) ClearRules() *NotificationChannelUpdate {
	ncu.mutation.ClearRules()
	return ncu
}

// RemoveRuleIDs removes the "rules" edge to NotificationRule entities by IDs.
func (ncu *NotificationChannelUpdate) RemoveRuleIDs(ids ...string) *NotificationChannelUpdate {
	ncu.mutation.RemoveRuleIDs(ids...)
	return ncu
}

// RemoveRules removes "rules" edges to NotificationRule entities.
func (ncu *NotificationChannelUpdate) RemoveRules(n ...*NotificationRule) *NotificationChannelUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ncu.RemoveRuleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ncu *NotificationChannelUpdate) Save(ctx context.Context) (int, error) {
	ncu.defaults()
	return withHooks(ctx, ncu.sqlSave, ncu.mutation, ncu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ncu *NotificationChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := ncu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ncu *NotificationChannelUpdate) Exec(ctx context.Context) error {
	_, err := ncu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncu *NotificationChannelUpdate) ExecX(ctx context.Context) {
	if err := ncu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ncu *NotificationChannelUpdate) defaults() {
	if _, ok := ncu.mutation.UpdatedAt(); !ok {
		v := notificationchannel.UpdateDefaultUpdatedAt()
		ncu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ncu *NotificationChannelUpdate) check() error {
	if v, ok := ncu.mutation.Name(); ok {
		if err := notificationchannel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "NotificationChannel.name": %w`, err)}
		}
	}
	if v, ok := ncu.mutation.Config(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "config", err: fmt.Errorf(`db: validator failed for field "NotificationChannel.config": %w`, err)}
		}
	}
	return nil
}

func (ncu *NotificationChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ncu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notificationchannel.Table, notificationchannel.Columns, sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeString))
	if ps := ncu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncu.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationchannel.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ncu.mutation.DeletedAt(); ok {
		_spec.SetField(notificationchannel.FieldDeletedAt, field.TypeTime, value)
	}
	if ncu.mutation.DeletedAtCleared() {
		_spec.ClearField(notificationchannel.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ncu.mutation.Name(); ok {
		_spec.SetField(notificationchannel.FieldName, field.TypeString, value)
	}
	if value, ok := ncu.mutation.Disabled(); ok {
		_spec.SetField(notificationchannel.FieldDisabled, field.TypeBool, value)
	}
	if ncu.mutation.DisabledCleared() {
		_spec.ClearField(notificationchannel.FieldDisabled, field.TypeBool)
	}
	if value, ok := ncu.mutation.Config(); ok {
		vv, err := notificationchannel.ValueScanner.Config.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(notificationchannel.FieldConfig, field.TypeString, vv)
	}
	if ncu.mutation.RulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   notificationchannel.RulesTable,
			Columns: notificationchannel.RulesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationrule.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ncu.mutation.RemovedRulesIDs(); len(nodes) > 0 && !ncu.mutation.RulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   notificationchannel.RulesTable,
			Columns: notificationchannel.RulesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationrule.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ncu.mutation.RulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   notificationchannel.RulesTable,
			Columns: notificationchannel.RulesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationrule.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ncu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationchannel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ncu.mutation.done = true
	return n, nil
}

// NotificationChannelUpdateOne is the builder for updating a single NotificationChannel entity.
type NotificationChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationChannelMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ncuo *NotificationChannelUpdateOne) SetUpdatedAt(t time.Time) *NotificationChannelUpdateOne {
	ncuo.mutation.SetUpdatedAt(t)
	return ncuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ncuo *NotificationChannelUpdateOne) SetDeletedAt(t time.Time) *NotificationChannelUpdateOne {
	ncuo.mutation.SetDeletedAt(t)
	return ncuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ncuo *NotificationChannelUpdateOne) SetNillableDeletedAt(t *time.Time) *NotificationChannelUpdateOne {
	if t != nil {
		ncuo.SetDeletedAt(*t)
	}
	return ncuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ncuo *NotificationChannelUpdateOne) ClearDeletedAt() *NotificationChannelUpdateOne {
	ncuo.mutation.ClearDeletedAt()
	return ncuo
}

// SetName sets the "name" field.
func (ncuo *NotificationChannelUpdateOne) SetName(s string) *NotificationChannelUpdateOne {
	ncuo.mutation.SetName(s)
	return ncuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ncuo *NotificationChannelUpdateOne) SetNillableName(s *string) *NotificationChannelUpdateOne {
	if s != nil {
		ncuo.SetName(*s)
	}
	return ncuo
}

// SetDisabled sets the "disabled" field.
func (ncuo *NotificationChannelUpdateOne) SetDisabled(b bool) *NotificationChannelUpdateOne {
	ncuo.mutation.SetDisabled(b)
	return ncuo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (ncuo *NotificationChannelUpdateOne) SetNillableDisabled(b *bool) *NotificationChannelUpdateOne {
	if b != nil {
		ncuo.SetDisabled(*b)
	}
	return ncuo
}

// ClearDisabled clears the value of the "disabled" field.
func (ncuo *NotificationChannelUpdateOne) ClearDisabled() *NotificationChannelUpdateOne {
	ncuo.mutation.ClearDisabled()
	return ncuo
}

// SetConfig sets the "config" field.
func (ncuo *NotificationChannelUpdateOne) SetConfig(nc notification.ChannelConfig) *NotificationChannelUpdateOne {
	ncuo.mutation.SetConfig(nc)
	return ncuo
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (ncuo *NotificationChannelUpdateOne) SetNillableConfig(nc *notification.ChannelConfig) *NotificationChannelUpdateOne {
	if nc != nil {
		ncuo.SetConfig(*nc)
	}
	return ncuo
}

// AddRuleIDs adds the "rules" edge to the NotificationRule entity by IDs.
func (ncuo *NotificationChannelUpdateOne) AddRuleIDs(ids ...string) *NotificationChannelUpdateOne {
	ncuo.mutation.AddRuleIDs(ids...)
	return ncuo
}

// AddRules adds the "rules" edges to the NotificationRule entity.
func (ncuo *NotificationChannelUpdateOne) AddRules(n ...*NotificationRule) *NotificationChannelUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ncuo.AddRuleIDs(ids...)
}

// Mutation returns the NotificationChannelMutation object of the builder.
func (ncuo *NotificationChannelUpdateOne) Mutation() *NotificationChannelMutation {
	return ncuo.mutation
}

// ClearRules clears all "rules" edges to the NotificationRule entity.
func (ncuo *NotificationChannelUpdateOne) ClearRules() *NotificationChannelUpdateOne {
	ncuo.mutation.ClearRules()
	return ncuo
}

// RemoveRuleIDs removes the "rules" edge to NotificationRule entities by IDs.
func (ncuo *NotificationChannelUpdateOne) RemoveRuleIDs(ids ...string) *NotificationChannelUpdateOne {
	ncuo.mutation.RemoveRuleIDs(ids...)
	return ncuo
}

// RemoveRules removes "rules" edges to NotificationRule entities.
func (ncuo *NotificationChannelUpdateOne) RemoveRules(n ...*NotificationRule) *NotificationChannelUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ncuo.RemoveRuleIDs(ids...)
}

// Where appends a list predicates to the NotificationChannelUpdate builder.
func (ncuo *NotificationChannelUpdateOne) Where(ps ...predicate.NotificationChannel) *NotificationChannelUpdateOne {
	ncuo.mutation.Where(ps...)
	return ncuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ncuo *NotificationChannelUpdateOne) Select(field string, fields ...string) *NotificationChannelUpdateOne {
	ncuo.fields = append([]string{field}, fields...)
	return ncuo
}

// Save executes the query and returns the updated NotificationChannel entity.
func (ncuo *NotificationChannelUpdateOne) Save(ctx context.Context) (*NotificationChannel, error) {
	ncuo.defaults()
	return withHooks(ctx, ncuo.sqlSave, ncuo.mutation, ncuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ncuo *NotificationChannelUpdateOne) SaveX(ctx context.Context) *NotificationChannel {
	node, err := ncuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ncuo *NotificationChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := ncuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncuo *NotificationChannelUpdateOne) ExecX(ctx context.Context) {
	if err := ncuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ncuo *NotificationChannelUpdateOne) defaults() {
	if _, ok := ncuo.mutation.UpdatedAt(); !ok {
		v := notificationchannel.UpdateDefaultUpdatedAt()
		ncuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ncuo *NotificationChannelUpdateOne) check() error {
	if v, ok := ncuo.mutation.Name(); ok {
		if err := notificationchannel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "NotificationChannel.name": %w`, err)}
		}
	}
	if v, ok := ncuo.mutation.Config(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "config", err: fmt.Errorf(`db: validator failed for field "NotificationChannel.config": %w`, err)}
		}
	}
	return nil
}

func (ncuo *NotificationChannelUpdateOne) sqlSave(ctx context.Context) (_node *NotificationChannel, err error) {
	if err := ncuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notificationchannel.Table, notificationchannel.Columns, sqlgraph.NewFieldSpec(notificationchannel.FieldID, field.TypeString))
	id, ok := ncuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "NotificationChannel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ncuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationchannel.FieldID)
		for _, f := range fields {
			if !notificationchannel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != notificationchannel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ncuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notificationchannel.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ncuo.mutation.DeletedAt(); ok {
		_spec.SetField(notificationchannel.FieldDeletedAt, field.TypeTime, value)
	}
	if ncuo.mutation.DeletedAtCleared() {
		_spec.ClearField(notificationchannel.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ncuo.mutation.Name(); ok {
		_spec.SetField(notificationchannel.FieldName, field.TypeString, value)
	}
	if value, ok := ncuo.mutation.Disabled(); ok {
		_spec.SetField(notificationchannel.FieldDisabled, field.TypeBool, value)
	}
	if ncuo.mutation.DisabledCleared() {
		_spec.ClearField(notificationchannel.FieldDisabled, field.TypeBool)
	}
	if value, ok := ncuo.mutation.Config(); ok {
		vv, err := notificationchannel.ValueScanner.Config.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(notificationchannel.FieldConfig, field.TypeString, vv)
	}
	if ncuo.mutation.RulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   notificationchannel.RulesTable,
			Columns: notificationchannel.RulesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationrule.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ncuo.mutation.RemovedRulesIDs(); len(nodes) > 0 && !ncuo.mutation.RulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   notificationchannel.RulesTable,
			Columns: notificationchannel.RulesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationrule.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ncuo.mutation.RulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   notificationchannel.RulesTable,
			Columns: notificationchannel.RulesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationrule.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NotificationChannel{config: ncuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ncuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notificationchannel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ncuo.mutation.done = true
	return _node, nil
}
