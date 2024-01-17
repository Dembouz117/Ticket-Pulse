package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Annotations(
			entproto.Field(1),
		),
		field.String("name").Annotations(
			entproto.Field(2),
		),
		field.String("email").Unique().Annotations(
			entproto.Field(3),
		),
		field.String("password").Sensitive().Annotations( // Store the hashed password in a sensitive field.
			entproto.Field(4),
		),
		field.String("phone").Annotations(
			entproto.Field(5),
		),
		field.Enum("role").Values("admin", "user").Annotations(
			entproto.Field(6),
			entproto.Enum(
				map[string]int32{
					"admin": 1,
					"user":  2,
				},
			),
		),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
