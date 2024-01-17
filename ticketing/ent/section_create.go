// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"ticketing/ent/concertsession"
	"ticketing/ent/section"
	"ticketing/ent/ticket"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SectionCreate is the builder for creating a Section entity.
type SectionCreate struct {
	config
	mutation *SectionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SectionCreate) SetName(s string) *SectionCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetCapacity sets the "capacity" field.
func (sc *SectionCreate) SetCapacity(i int) *SectionCreate {
	sc.mutation.SetCapacity(i)
	return sc
}

// SetReserved sets the "reserved" field.
func (sc *SectionCreate) SetReserved(i int) *SectionCreate {
	sc.mutation.SetReserved(i)
	return sc
}

// SetBought sets the "bought" field.
func (sc *SectionCreate) SetBought(i int) *SectionCreate {
	sc.mutation.SetBought(i)
	return sc
}

// SetCategory sets the "category" field.
func (sc *SectionCreate) SetCategory(s section.Category) *SectionCreate {
	sc.mutation.SetCategory(s)
	return sc
}

// SetPrice sets the "price" field.
func (sc *SectionCreate) SetPrice(i int) *SectionCreate {
	sc.mutation.SetPrice(i)
	return sc
}

// SetID sets the "id" field.
func (sc *SectionCreate) SetID(u uuid.UUID) *SectionCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SectionCreate) SetNillableID(u *uuid.UUID) *SectionCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// AddHasTicketIDs adds the "hasTickets" edge to the Ticket entity by IDs.
func (sc *SectionCreate) AddHasTicketIDs(ids ...uuid.UUID) *SectionCreate {
	sc.mutation.AddHasTicketIDs(ids...)
	return sc
}

// AddHasTickets adds the "hasTickets" edges to the Ticket entity.
func (sc *SectionCreate) AddHasTickets(t ...*Ticket) *SectionCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddHasTicketIDs(ids...)
}

// SetAtConcertSessionID sets the "atConcertSession" edge to the ConcertSession entity by ID.
func (sc *SectionCreate) SetAtConcertSessionID(id uuid.UUID) *SectionCreate {
	sc.mutation.SetAtConcertSessionID(id)
	return sc
}

// SetAtConcertSession sets the "atConcertSession" edge to the ConcertSession entity.
func (sc *SectionCreate) SetAtConcertSession(c *ConcertSession) *SectionCreate {
	return sc.SetAtConcertSessionID(c.ID)
}

// Mutation returns the SectionMutation object of the builder.
func (sc *SectionCreate) Mutation() *SectionMutation {
	return sc.mutation
}

// Save creates the Section in the database.
func (sc *SectionCreate) Save(ctx context.Context) (*Section, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SectionCreate) SaveX(ctx context.Context) *Section {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SectionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SectionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SectionCreate) defaults() {
	if _, ok := sc.mutation.ID(); !ok {
		v := section.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SectionCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Section.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := section.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Section.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Capacity(); !ok {
		return &ValidationError{Name: "capacity", err: errors.New(`ent: missing required field "Section.capacity"`)}
	}
	if _, ok := sc.mutation.Reserved(); !ok {
		return &ValidationError{Name: "reserved", err: errors.New(`ent: missing required field "Section.reserved"`)}
	}
	if _, ok := sc.mutation.Bought(); !ok {
		return &ValidationError{Name: "bought", err: errors.New(`ent: missing required field "Section.bought"`)}
	}
	if _, ok := sc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Section.category"`)}
	}
	if v, ok := sc.mutation.Category(); ok {
		if err := section.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Section.category": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Section.price"`)}
	}
	if _, ok := sc.mutation.AtConcertSessionID(); !ok {
		return &ValidationError{Name: "atConcertSession", err: errors.New(`ent: missing required edge "Section.atConcertSession"`)}
	}
	return nil
}

func (sc *SectionCreate) sqlSave(ctx context.Context) (*Section, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SectionCreate) createSpec() (*Section, *sqlgraph.CreateSpec) {
	var (
		_node = &Section{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(section.Table, sqlgraph.NewFieldSpec(section.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(section.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Capacity(); ok {
		_spec.SetField(section.FieldCapacity, field.TypeInt, value)
		_node.Capacity = value
	}
	if value, ok := sc.mutation.Reserved(); ok {
		_spec.SetField(section.FieldReserved, field.TypeInt, value)
		_node.Reserved = value
	}
	if value, ok := sc.mutation.Bought(); ok {
		_spec.SetField(section.FieldBought, field.TypeInt, value)
		_node.Bought = value
	}
	if value, ok := sc.mutation.Category(); ok {
		_spec.SetField(section.FieldCategory, field.TypeEnum, value)
		_node.Category = value
	}
	if value, ok := sc.mutation.Price(); ok {
		_spec.SetField(section.FieldPrice, field.TypeInt, value)
		_node.Price = value
	}
	if nodes := sc.mutation.HasTicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   section.HasTicketsTable,
			Columns: []string{section.HasTicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.AtConcertSessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   section.AtConcertSessionTable,
			Columns: []string{section.AtConcertSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(concertsession.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.concert_session_has_sections = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SectionCreateBulk is the builder for creating many Section entities in bulk.
type SectionCreateBulk struct {
	config
	err      error
	builders []*SectionCreate
}

// Save creates the Section entities in the database.
func (scb *SectionCreateBulk) Save(ctx context.Context) ([]*Section, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Section, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SectionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SectionCreateBulk) SaveX(ctx context.Context) []*Section {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SectionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SectionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}