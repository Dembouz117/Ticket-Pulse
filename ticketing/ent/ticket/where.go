// Code generated by ent, DO NOT EDIT.

package ticket

import (
	"ticketing/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldID, id))
}

// SeatNumber applies equality check predicate on the "seatNumber" field. It's identical to SeatNumberEQ.
func SeatNumber(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldSeatNumber, v))
}

// UserId applies equality check predicate on the "userId" field. It's identical to UserIdEQ.
func UserId(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldUserId, v))
}

// ReservedAt applies equality check predicate on the "reservedAt" field. It's identical to ReservedAtEQ.
func ReservedAt(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldReservedAt, v))
}

// SeatNumberEQ applies the EQ predicate on the "seatNumber" field.
func SeatNumberEQ(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldSeatNumber, v))
}

// SeatNumberNEQ applies the NEQ predicate on the "seatNumber" field.
func SeatNumberNEQ(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldSeatNumber, v))
}

// SeatNumberIn applies the In predicate on the "seatNumber" field.
func SeatNumberIn(vs ...int) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldSeatNumber, vs...))
}

// SeatNumberNotIn applies the NotIn predicate on the "seatNumber" field.
func SeatNumberNotIn(vs ...int) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldSeatNumber, vs...))
}

// SeatNumberGT applies the GT predicate on the "seatNumber" field.
func SeatNumberGT(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldSeatNumber, v))
}

// SeatNumberGTE applies the GTE predicate on the "seatNumber" field.
func SeatNumberGTE(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldSeatNumber, v))
}

// SeatNumberLT applies the LT predicate on the "seatNumber" field.
func SeatNumberLT(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldSeatNumber, v))
}

// SeatNumberLTE applies the LTE predicate on the "seatNumber" field.
func SeatNumberLTE(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldSeatNumber, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldStatus, vs...))
}

// UserIdEQ applies the EQ predicate on the "userId" field.
func UserIdEQ(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldUserId, v))
}

// UserIdNEQ applies the NEQ predicate on the "userId" field.
func UserIdNEQ(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldUserId, v))
}

// UserIdIn applies the In predicate on the "userId" field.
func UserIdIn(vs ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldUserId, vs...))
}

// UserIdNotIn applies the NotIn predicate on the "userId" field.
func UserIdNotIn(vs ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldUserId, vs...))
}

// UserIdGT applies the GT predicate on the "userId" field.
func UserIdGT(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldUserId, v))
}

// UserIdGTE applies the GTE predicate on the "userId" field.
func UserIdGTE(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldUserId, v))
}

// UserIdLT applies the LT predicate on the "userId" field.
func UserIdLT(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldUserId, v))
}

// UserIdLTE applies the LTE predicate on the "userId" field.
func UserIdLTE(v uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldUserId, v))
}

// UserIdIsNil applies the IsNil predicate on the "userId" field.
func UserIdIsNil() predicate.Ticket {
	return predicate.Ticket(sql.FieldIsNull(FieldUserId))
}

// UserIdNotNil applies the NotNil predicate on the "userId" field.
func UserIdNotNil() predicate.Ticket {
	return predicate.Ticket(sql.FieldNotNull(FieldUserId))
}

// ReservedAtEQ applies the EQ predicate on the "reservedAt" field.
func ReservedAtEQ(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldReservedAt, v))
}

// ReservedAtNEQ applies the NEQ predicate on the "reservedAt" field.
func ReservedAtNEQ(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldReservedAt, v))
}

// ReservedAtIn applies the In predicate on the "reservedAt" field.
func ReservedAtIn(vs ...int) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldReservedAt, vs...))
}

// ReservedAtNotIn applies the NotIn predicate on the "reservedAt" field.
func ReservedAtNotIn(vs ...int) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldReservedAt, vs...))
}

// ReservedAtGT applies the GT predicate on the "reservedAt" field.
func ReservedAtGT(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldReservedAt, v))
}

// ReservedAtGTE applies the GTE predicate on the "reservedAt" field.
func ReservedAtGTE(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldReservedAt, v))
}

// ReservedAtLT applies the LT predicate on the "reservedAt" field.
func ReservedAtLT(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldReservedAt, v))
}

// ReservedAtLTE applies the LTE predicate on the "reservedAt" field.
func ReservedAtLTE(v int) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldReservedAt, v))
}

// ReservedAtIsNil applies the IsNil predicate on the "reservedAt" field.
func ReservedAtIsNil() predicate.Ticket {
	return predicate.Ticket(sql.FieldIsNull(FieldReservedAt))
}

// ReservedAtNotNil applies the NotNil predicate on the "reservedAt" field.
func ReservedAtNotNil() predicate.Ticket {
	return predicate.Ticket(sql.FieldNotNull(FieldReservedAt))
}

// HasWithinSection applies the HasEdge predicate on the "withinSection" edge.
func HasWithinSection() predicate.Ticket {
	return predicate.Ticket(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WithinSectionTable, WithinSectionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWithinSectionWith applies the HasEdge predicate on the "withinSection" edge with a given conditions (other predicates).
func HasWithinSectionWith(preds ...predicate.Section) predicate.Ticket {
	return predicate.Ticket(func(s *sql.Selector) {
		step := newWithinSectionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ticket) predicate.Ticket {
	return predicate.Ticket(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ticket) predicate.Ticket {
	return predicate.Ticket(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Ticket) predicate.Ticket {
	return predicate.Ticket(sql.NotPredicates(p))
}