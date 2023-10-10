package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type Token struct {
	ent.Schema
}

// Fields of the User.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
		field.String("token").Annotations(entgql.OrderField("TOKEN")),
		field.UUID("user_id", uuid.UUID{}).Annotations(entgql.OrderField("USER_ID")).Optional(),
	}
}

// Annotations of the Alert.
func (Token) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tokens"},
	}
}

// Edges of the User.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{}
}
