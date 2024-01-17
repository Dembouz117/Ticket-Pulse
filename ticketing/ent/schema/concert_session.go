package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ConcertSession holds the schema definition for the ConcertSession entity.
type ConcertSession struct {
	ent.Schema
}

func (ConcertSession) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Annotations(
			entproto.Field(1),
		),
		field.Int("sessionDateTime").Unique().Annotations(
			entproto.Field(2),
		),
	}
}

func (ConcertSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ofConcert", Concert.Type).Ref("hasConcertSessions").Required().Annotations(
			entproto.Field(3),
		),
		edge.To("hasSections", Section.Type).Annotations(
			entproto.Field(4),
		),
	}
}

func (ConcertSession) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
