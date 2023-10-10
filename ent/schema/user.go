package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
		field.String("fullName").MaxLen(255).Annotations(entgql.OrderField("FULLNAME")),
		field.String("username").MaxLen(255).Annotations(entgql.OrderField("USERNAME")).Unique(),
		field.String("password").MaxLen(255).Annotations(entgql.OrderField("PASSWORD")),
		field.String("email").Annotations(entgql.OrderField("EMAIL")).Unique(),
		field.Enum("role").Values("ADMIN", "USER").Annotations(entgql.OrderField("ROLE")),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Annotations of the Alert.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
