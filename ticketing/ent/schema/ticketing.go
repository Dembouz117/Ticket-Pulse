package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Section holds the schema definition for the Section entity.
type Section struct {
	ent.Schema
}

func (Section) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Annotations(
			entproto.Field(1),
		),
		field.String("name").NotEmpty().Annotations(
			entproto.Field(2),
		),
		field.Int("capacity").Annotations(
			entproto.Field(3),
		),
		field.Int("reserved").Annotations(
			entproto.Field(4),
		),
		field.Int("bought").Annotations(
			entproto.Field(5),
		),
		field.Enum("category").Values("CAT1", "CAT2", "CAT3", "CAT4", "CAT5").Annotations(
			entproto.Field(6),
			entproto.Enum(map[string]int32{
				"CAT1": 1,
				"CAT2": 2,
				"CAT3": 3,
				"CAT4": 4,
				"CAT5": 5,
			},
			),
		),
		field.Int("price").Annotations(
			entproto.Field(7),
		),
	}
}

func (Section) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hasTickets", Ticket.Type).Annotations(
			entproto.Field(8),
		),
		edge.From("atConcertSession", ConcertSession.Type).Ref("hasSections").Unique().Required().Annotations(
			entproto.Field(9),
		),
	}
}

func (Section) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Tickets holds the schema definition for the Tickets entity.
type Ticket struct {
	ent.Schema
}

func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Annotations(
			entproto.Field(1),
		),
		field.Int("seatNumber").Annotations(
			entproto.Field(2),
		),
		field.Enum("status").Values("AVAILABLE", "BOUGHT", "RESERVED").Annotations(
			entproto.Field(3),
			entproto.Enum(map[string]int32{
				"AVAILABLE": 1,
				"BOUGHT":    2,
				"RESERVED":  3,
			},
			)),
		field.UUID("userId", uuid.UUID{}).Optional().Annotations(
			entproto.Field(4),
		),
		field.Int("reservedAt").Optional().Nillable().Annotations(
			entproto.Field(5),
		),
	}
}

func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("withinSection", Section.Type).Ref("hasTickets").Unique().Required().Annotations(
			entproto.Field(5),
		),
	}
}

func (Ticket) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

func (Ticket) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("userId"),
	}
}
