// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"golang-boilerplate/ent/item"
	"golang-boilerplate/ent/vendor"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Item holds the value of the "item" field.
	Item string `json:"item,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// RemainingAmount holds the value of the "remaining_amount" field.
	RemainingAmount int `json:"remaining_amount,omitempty"`
	// SoldAmount holds the value of the "sold_amount" field.
	SoldAmount int `json:"sold_amount,omitempty"`
	// Exp holds the value of the "exp" field.
	Exp time.Time `json:"exp,omitempty"`
	// VendorID holds the value of the "vendor_id" field.
	VendorID uuid.UUID `json:"vendor_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemQuery when eager-loading is set.
	Edges ItemEdges `json:"edges"`
}

// ItemEdges holds the relations/edges for other nodes in the graph.
type ItemEdges struct {
	// Vendors holds the value of the vendors edge.
	Vendors *Vendor `json:"vendors,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// VendorsOrErr returns the Vendors value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemEdges) VendorsOrErr() (*Vendor, error) {
	if e.loadedTypes[0] {
		if e.Vendors == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: vendor.Label}
		}
		return e.Vendors, nil
	}
	return nil, &NotLoadedError{edge: "vendors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case item.FieldPrice, item.FieldRemainingAmount, item.FieldSoldAmount:
			values[i] = new(sql.NullInt64)
		case item.FieldItem:
			values[i] = new(sql.NullString)
		case item.FieldExp, item.FieldCreatedAt, item.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case item.FieldID, item.FieldVendorID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Item", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case item.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case item.FieldItem:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field item", values[j])
			} else if value.Valid {
				i.Item = value.String
			}
		case item.FieldPrice:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[j])
			} else if value.Valid {
				i.Price = int(value.Int64)
			}
		case item.FieldRemainingAmount:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field remaining_amount", values[j])
			} else if value.Valid {
				i.RemainingAmount = int(value.Int64)
			}
		case item.FieldSoldAmount:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sold_amount", values[j])
			} else if value.Valid {
				i.SoldAmount = int(value.Int64)
			}
		case item.FieldExp:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field exp", values[j])
			} else if value.Valid {
				i.Exp = value.Time
			}
		case item.FieldVendorID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field vendor_id", values[j])
			} else if value != nil {
				i.VendorID = *value
			}
		case item.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case item.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryVendors queries the "vendors" edge of the Item entity.
func (i *Item) QueryVendors() *VendorQuery {
	return (&ItemClient{config: i.config}).QueryVendors(i)
}

// Update returns a builder for updating this Item.
// Note that you need to call Item.Unwrap() before calling this method if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return (&ItemClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("item=")
	builder.WriteString(i.Item)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", i.Price))
	builder.WriteString(", ")
	builder.WriteString("remaining_amount=")
	builder.WriteString(fmt.Sprintf("%v", i.RemainingAmount))
	builder.WriteString(", ")
	builder.WriteString("sold_amount=")
	builder.WriteString(fmt.Sprintf("%v", i.SoldAmount))
	builder.WriteString(", ")
	builder.WriteString("exp=")
	builder.WriteString(i.Exp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("vendor_id=")
	builder.WriteString(fmt.Sprintf("%v", i.VendorID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item

func (i Items) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
