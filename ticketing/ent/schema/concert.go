package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Concert struct {
	ent.Schema
}

func (Concert) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Annotations(
			entproto.Field(1),
		),
		field.String("title").NotEmpty().Annotations(
			entproto.Field(2),
		),
		field.String("artist").NotEmpty().Annotations(
			entproto.Field(3),
		),
		field.String("imageUrl").NotEmpty().Annotations(
			entproto.Field(4),
		),
		field.String("description").Default("").Annotations(
			entproto.Field(5),
		),
		field.String("headline").Default("").Annotations(
			entproto.Field(6),
		),
		field.Bool("featured").Default(false).Annotations(
			entproto.Field(7),
		),
	}
}

func (Concert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hasConcertSessions", ConcertSession.Type).Annotations(
			entproto.Field(8),
		),
		edge.To("genres", Genre.Type).Annotations(
			entproto.Field(9),
		),
	}
}

func (Concert) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

type Genre struct {
	ent.Schema
}

func (Genre) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Annotations(
			entproto.Field(1),
		),
		field.String("name").NotEmpty().Annotations(
			entproto.Field(2),
		),
	}
}

func (Genre) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("concerts", Concert.Type).Ref("genres").Annotations(
			entproto.Field(3),
		),
	}
}

func (Genre) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
