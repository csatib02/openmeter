// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/notificationrule"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// NotificationRuleDelete is the builder for deleting a NotificationRule entity.
type NotificationRuleDelete struct {
	config
	hooks    []Hook
	mutation *NotificationRuleMutation
}

// Where appends a list predicates to the NotificationRuleDelete builder.
func (nrd *NotificationRuleDelete) Where(ps ...predicate.NotificationRule) *NotificationRuleDelete {
	nrd.mutation.Where(ps...)
	return nrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nrd *NotificationRuleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nrd.sqlExec, nrd.mutation, nrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nrd *NotificationRuleDelete) ExecX(ctx context.Context) int {
	n, err := nrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nrd *NotificationRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(notificationrule.Table, sqlgraph.NewFieldSpec(notificationrule.FieldID, field.TypeString))
	if ps := nrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nrd.mutation.done = true
	return affected, err
}

// NotificationRuleDeleteOne is the builder for deleting a single NotificationRule entity.
type NotificationRuleDeleteOne struct {
	nrd *NotificationRuleDelete
}

// Where appends a list predicates to the NotificationRuleDelete builder.
func (nrdo *NotificationRuleDeleteOne) Where(ps ...predicate.NotificationRule) *NotificationRuleDeleteOne {
	nrdo.nrd.mutation.Where(ps...)
	return nrdo
}

// Exec executes the deletion query.
func (nrdo *NotificationRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := nrdo.nrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{notificationrule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nrdo *NotificationRuleDeleteOne) ExecX(ctx context.Context) {
	if err := nrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
