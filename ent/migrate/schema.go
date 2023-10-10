// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "item", Type: field.TypeString, Size: 255},
		{Name: "price", Type: field.TypeInt},
		{Name: "remaining_amount", Type: field.TypeInt},
		{Name: "sold_amount", Type: field.TypeInt},
		{Name: "exp", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "vendor_id", Type: field.TypeUUID},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "items_vendors_item",
				Columns:    []*schema.Column{ItemsColumns[8]},
				RefColumns: []*schema.Column{VendorsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "token", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "full_name", Type: field.TypeString, Size: 255},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "password", Type: field.TypeString, Size: 255},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "USER"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VendorsColumns holds the columns for the "vendors" table.
	VendorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "vendor", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// VendorsTable holds the schema information for the "vendors" table.
	VendorsTable = &schema.Table{
		Name:       "vendors",
		Columns:    VendorsColumns,
		PrimaryKey: []*schema.Column{VendorsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ItemsTable,
		TokensTable,
		UsersTable,
		VendorsTable,
	}
)

func init() {
	ItemsTable.ForeignKeys[0].RefTable = VendorsTable
	ItemsTable.Annotation = &entsql.Annotation{
		Table: "items",
	}
	TokensTable.Annotation = &entsql.Annotation{
		Table: "tokens",
	}
	UsersTable.Annotation = &entsql.Annotation{
		Table: "users",
	}
	VendorsTable.Annotation = &entsql.Annotation{
		Table: "vendors",
	}
}
