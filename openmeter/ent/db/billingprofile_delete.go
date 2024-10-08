// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billingprofile"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// BillingProfileDelete is the builder for deleting a BillingProfile entity.
type BillingProfileDelete struct {
	config
	hooks    []Hook
	mutation *BillingProfileMutation
}

// Where appends a list predicates to the BillingProfileDelete builder.
func (bpd *BillingProfileDelete) Where(ps ...predicate.BillingProfile) *BillingProfileDelete {
	bpd.mutation.Where(ps...)
	return bpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bpd *BillingProfileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bpd.sqlExec, bpd.mutation, bpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bpd *BillingProfileDelete) ExecX(ctx context.Context) int {
	n, err := bpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bpd *BillingProfileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(billingprofile.Table, sqlgraph.NewFieldSpec(billingprofile.FieldID, field.TypeString))
	if ps := bpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bpd.mutation.done = true
	return affected, err
}

// BillingProfileDeleteOne is the builder for deleting a single BillingProfile entity.
type BillingProfileDeleteOne struct {
	bpd *BillingProfileDelete
}

// Where appends a list predicates to the BillingProfileDelete builder.
func (bpdo *BillingProfileDeleteOne) Where(ps ...predicate.BillingProfile) *BillingProfileDeleteOne {
	bpdo.bpd.mutation.Where(ps...)
	return bpdo
}

// Exec executes the deletion query.
func (bpdo *BillingProfileDeleteOne) Exec(ctx context.Context) error {
	n, err := bpdo.bpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{billingprofile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bpdo *BillingProfileDeleteOne) ExecX(ctx context.Context) {
	if err := bpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
