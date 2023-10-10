package schema

import (
	"entgo.io/ent/schema/edge"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Items holds the schema definition for the Items entity.
type Item struct {
	ent.Schema
}

// Fields of the Items.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
		field.String("item").MaxLen(255).Annotations(entgql.OrderField("ITEM")),
		field.Int("price").Annotations(entgql.OrderField("PRICE")),
		field.Int("remaining_amount").Annotations(entgql.OrderField("REMAINING_AMOUNT")),
		field.Int("sold_amount").Annotations(entgql.OrderField("SOLD_AMOUNT")),
		field.Time("exp").Annotations(entgql.OrderField("EXP")),
		field.UUID("vendor_id", uuid.UUID{}).Annotations(entgql.OrderField("VENDOR_ID")),

		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Annotations of the Alert.
func (Item) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "items"},
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("vendors", Vendor.Type).Ref("item").Field("vendor_id").Unique().Required(),
	}
}
