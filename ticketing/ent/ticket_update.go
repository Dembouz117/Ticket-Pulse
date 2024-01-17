// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"ticketing/ent/predicate"
	"ticketing/ent/section"
	"ticketing/ent/ticket"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TicketUpdate is the builder for updating Ticket entities.
type TicketUpdate struct {
	config
	hooks    []Hook
	mutation *TicketMutation
}

// Where appends a list predicates to the TicketUpdate builder.
func (tu *TicketUpdate) Where(ps ...predicate.Ticket) *TicketUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetSeatNumber sets the "seatNumber" field.
func (tu *TicketUpdate) SetSeatNumber(i int) *TicketUpdate {
	tu.mutation.ResetSeatNumber()
	tu.mutation.SetSeatNumber(i)
	return tu
}

// AddSeatNumber adds i to the "seatNumber" field.
func (tu *TicketUpdate) AddSeatNumber(i int) *TicketUpdate {
	tu.mutation.AddSeatNumber(i)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TicketUpdate) SetStatus(t ticket.Status) *TicketUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetUserId sets the "userId" field.
func (tu *TicketUpdate) SetUserId(u uuid.UUID) *TicketUpdate {
	tu.mutation.SetUserId(u)
	return tu
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (tu *TicketUpdate) SetNillableUserId(u *uuid.UUID) *TicketUpdate {
	if u != nil {
		tu.SetUserId(*u)
	}
	return tu
}

// ClearUserId clears the value of the "userId" field.
func (tu *TicketUpdate) ClearUserId() *TicketUpdate {
	tu.mutation.ClearUserId()
	return tu
}

// SetReservedAt sets the "reservedAt" field.
func (tu *TicketUpdate) SetReservedAt(i int) *TicketUpdate {
	tu.mutation.ResetReservedAt()
	tu.mutation.SetReservedAt(i)
	return tu
}

// SetNillableReservedAt sets the "reservedAt" field if the given value is not nil.
func (tu *TicketUpdate) SetNillableReservedAt(i *int) *TicketUpdate {
	if i != nil {
		tu.SetReservedAt(*i)
	}
	return tu
}

// AddReservedAt adds i to the "reservedAt" field.
func (tu *TicketUpdate) AddReservedAt(i int) *TicketUpdate {
	tu.mutation.AddReservedAt(i)
	return tu
}

// ClearReservedAt clears the value of the "reservedAt" field.
func (tu *TicketUpdate) ClearReservedAt() *TicketUpdate {
	tu.mutation.ClearReservedAt()
	return tu
}

// SetWithinSectionID sets the "withinSection" edge to the Section entity by ID.
func (tu *TicketUpdate) SetWithinSectionID(id uuid.UUID) *TicketUpdate {
	tu.mutation.SetWithinSectionID(id)
	return tu
}

// SetWithinSection sets the "withinSection" edge to the Section entity.
func (tu *TicketUpdate) SetWithinSection(s *Section) *TicketUpdate {
	return tu.SetWithinSectionID(s.ID)
}

// Mutation returns the TicketMutation object of the builder.
func (tu *TicketUpdate) Mutation() *TicketMutation {
	return tu.mutation
}

// ClearWithinSection clears the "withinSection" edge to the Section entity.
func (tu *TicketUpdate) ClearWithinSection() *TicketUpdate {
	tu.mutation.ClearWithinSection()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TicketUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TicketUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TicketUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TicketUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TicketUpdate) check() error {
	if v, ok := tu.mutation.Status(); ok {
		if err := ticket.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Ticket.status": %w`, err)}
		}
	}
	if _, ok := tu.mutation.WithinSectionID(); tu.mutation.WithinSectionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Ticket.withinSection"`)
	}
	return nil
}

func (tu *TicketUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(ticket.Table, ticket.Columns, sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.SeatNumber(); ok {
		_spec.SetField(ticket.FieldSeatNumber, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedSeatNumber(); ok {
		_spec.AddField(ticket.FieldSeatNumber, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(ticket.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.UserId(); ok {
		_spec.SetField(ticket.FieldUserId, field.TypeUUID, value)
	}
	if tu.mutation.UserIdCleared() {
		_spec.ClearField(ticket.FieldUserId, field.TypeUUID)
	}
	if value, ok := tu.mutation.ReservedAt(); ok {
		_spec.SetField(ticket.FieldReservedAt, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedReservedAt(); ok {
		_spec.AddField(ticket.FieldReservedAt, field.TypeInt, value)
	}
	if tu.mutation.ReservedAtCleared() {
		_spec.ClearField(ticket.FieldReservedAt, field.TypeInt)
	}
	if tu.mutation.WithinSectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.WithinSectionTable,
			Columns: []string{ticket.WithinSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(section.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.WithinSectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.WithinSectionTable,
			Columns: []string{ticket.WithinSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(section.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ticket.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TicketUpdateOne is the builder for updating a single Ticket entity.
type TicketUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TicketMutation
}

// SetSeatNumber sets the "seatNumber" field.
func (tuo *TicketUpdateOne) SetSeatNumber(i int) *TicketUpdateOne {
	tuo.mutation.ResetSeatNumber()
	tuo.mutation.SetSeatNumber(i)
	return tuo
}

// AddSeatNumber adds i to the "seatNumber" field.
func (tuo *TicketUpdateOne) AddSeatNumber(i int) *TicketUpdateOne {
	tuo.mutation.AddSeatNumber(i)
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TicketUpdateOne) SetStatus(t ticket.Status) *TicketUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetUserId sets the "userId" field.
func (tuo *TicketUpdateOne) SetUserId(u uuid.UUID) *TicketUpdateOne {
	tuo.mutation.SetUserId(u)
	return tuo
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (tuo *TicketUpdateOne) SetNillableUserId(u *uuid.UUID) *TicketUpdateOne {
	if u != nil {
		tuo.SetUserId(*u)
	}
	return tuo
}

// ClearUserId clears the value of the "userId" field.
func (tuo *TicketUpdateOne) ClearUserId() *TicketUpdateOne {
	tuo.mutation.ClearUserId()
	return tuo
}

// SetReservedAt sets the "reservedAt" field.
func (tuo *TicketUpdateOne) SetReservedAt(i int) *TicketUpdateOne {
	tuo.mutation.ResetReservedAt()
	tuo.mutation.SetReservedAt(i)
	return tuo
}

// SetNillableReservedAt sets the "reservedAt" field if the given value is not nil.
func (tuo *TicketUpdateOne) SetNillableReservedAt(i *int) *TicketUpdateOne {
	if i != nil {
		tuo.SetReservedAt(*i)
	}
	return tuo
}

// AddReservedAt adds i to the "reservedAt" field.
func (tuo *TicketUpdateOne) AddReservedAt(i int) *TicketUpdateOne {
	tuo.mutation.AddReservedAt(i)
	return tuo
}

// ClearReservedAt clears the value of the "reservedAt" field.
func (tuo *TicketUpdateOne) ClearReservedAt() *TicketUpdateOne {
	tuo.mutation.ClearReservedAt()
	return tuo
}

// SetWithinSectionID sets the "withinSection" edge to the Section entity by ID.
func (tuo *TicketUpdateOne) SetWithinSectionID(id uuid.UUID) *TicketUpdateOne {
	tuo.mutation.SetWithinSectionID(id)
	return tuo
}

// SetWithinSection sets the "withinSection" edge to the Section entity.
func (tuo *TicketUpdateOne) SetWithinSection(s *Section) *TicketUpdateOne {
	return tuo.SetWithinSectionID(s.ID)
}

// Mutation returns the TicketMutation object of the builder.
func (tuo *TicketUpdateOne) Mutation() *TicketMutation {
	return tuo.mutation
}

// ClearWithinSection clears the "withinSection" edge to the Section entity.
func (tuo *TicketUpdateOne) ClearWithinSection() *TicketUpdateOne {
	tuo.mutation.ClearWithinSection()
	return tuo
}

// Where appends a list predicates to the TicketUpdate builder.
func (tuo *TicketUpdateOne) Where(ps ...predicate.Ticket) *TicketUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TicketUpdateOne) Select(field string, fields ...string) *TicketUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Ticket entity.
func (tuo *TicketUpdateOne) Save(ctx context.Context) (*Ticket, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TicketUpdateOne) SaveX(ctx context.Context) *Ticket {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TicketUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TicketUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TicketUpdateOne) check() error {
	if v, ok := tuo.mutation.Status(); ok {
		if err := ticket.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Ticket.status": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.WithinSectionID(); tuo.mutation.WithinSectionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Ticket.withinSection"`)
	}
	return nil
}

func (tuo *TicketUpdateOne) sqlSave(ctx context.Context) (_node *Ticket, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(ticket.Table, ticket.Columns, sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Ticket.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ticket.FieldID)
		for _, f := range fields {
			if !ticket.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ticket.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.SeatNumber(); ok {
		_spec.SetField(ticket.FieldSeatNumber, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedSeatNumber(); ok {
		_spec.AddField(ticket.FieldSeatNumber, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(ticket.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.UserId(); ok {
		_spec.SetField(ticket.FieldUserId, field.TypeUUID, value)
	}
	if tuo.mutation.UserIdCleared() {
		_spec.ClearField(ticket.FieldUserId, field.TypeUUID)
	}
	if value, ok := tuo.mutation.ReservedAt(); ok {
		_spec.SetField(ticket.FieldReservedAt, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedReservedAt(); ok {
		_spec.AddField(ticket.FieldReservedAt, field.TypeInt, value)
	}
	if tuo.mutation.ReservedAtCleared() {
		_spec.ClearField(ticket.FieldReservedAt, field.TypeInt)
	}
	if tuo.mutation.WithinSectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.WithinSectionTable,
			Columns: []string{ticket.WithinSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(section.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.WithinSectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.WithinSectionTable,
			Columns: []string{ticket.WithinSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(section.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Ticket{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ticket.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
