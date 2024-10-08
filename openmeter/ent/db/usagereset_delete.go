// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
	"github.com/openmeterio/openmeter/openmeter/ent/db/usagereset"
)

// UsageResetDelete is the builder for deleting a UsageReset entity.
type UsageResetDelete struct {
	config
	hooks    []Hook
	mutation *UsageResetMutation
}

// Where appends a list predicates to the UsageResetDelete builder.
func (urd *UsageResetDelete) Where(ps ...predicate.UsageReset) *UsageResetDelete {
	urd.mutation.Where(ps...)
	return urd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (urd *UsageResetDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, urd.sqlExec, urd.mutation, urd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (urd *UsageResetDelete) ExecX(ctx context.Context) int {
	n, err := urd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (urd *UsageResetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usagereset.Table, sqlgraph.NewFieldSpec(usagereset.FieldID, field.TypeString))
	if ps := urd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, urd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	urd.mutation.done = true
	return affected, err
}

// UsageResetDeleteOne is the builder for deleting a single UsageReset entity.
type UsageResetDeleteOne struct {
	urd *UsageResetDelete
}

// Where appends a list predicates to the UsageResetDelete builder.
func (urdo *UsageResetDeleteOne) Where(ps ...predicate.UsageReset) *UsageResetDeleteOne {
	urdo.urd.mutation.Where(ps...)
	return urdo
}

// Exec executes the deletion query.
func (urdo *UsageResetDeleteOne) Exec(ctx context.Context) error {
	n, err := urdo.urd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usagereset.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (urdo *UsageResetDeleteOne) ExecX(ctx context.Context) {
	if err := urdo.Exec(ctx); err != nil {
		panic(err)
	}
}
