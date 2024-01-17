// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ticketing/ent/concert"
	"ticketing/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConcertDelete is the builder for deleting a Concert entity.
type ConcertDelete struct {
	config
	hooks    []Hook
	mutation *ConcertMutation
}

// Where appends a list predicates to the ConcertDelete builder.
func (cd *ConcertDelete) Where(ps ...predicate.Concert) *ConcertDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *ConcertDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *ConcertDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *ConcertDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(concert.Table, sqlgraph.NewFieldSpec(concert.FieldID, field.TypeUUID))
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// ConcertDeleteOne is the builder for deleting a single Concert entity.
type ConcertDeleteOne struct {
	cd *ConcertDelete
}

// Where appends a list predicates to the ConcertDelete builder.
func (cdo *ConcertDeleteOne) Where(ps ...predicate.Concert) *ConcertDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *ConcertDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{concert.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *ConcertDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}
