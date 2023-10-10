package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type Vendor struct {
	ent.Schema
}

// Fields of the User.
func (Vendor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
		field.String("vendor").MaxLen(255).Annotations(entgql.OrderField("VENDOR")),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Annotations of the Alert.
func (Vendor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "vendors"},
	}
}

// Edges of the Vendor.
func (Vendor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item", Item.Type),
	}
}
