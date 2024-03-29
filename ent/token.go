// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"golang-boilerplate/ent/token"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Token is the model entity for the Token schema.
type Token struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Token) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case token.FieldToken:
			values[i] = new(sql.NullString)
		case token.FieldID, token.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Token", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Token fields.
func (t *Token) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case token.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case token.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				t.Token = value.String
			}
		case token.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				t.UserID = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Token.
// Note that you need to call Token.Unwrap() before calling this method if this Token
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Token) Update() *TokenUpdateOne {
	return (&TokenClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Token entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Token) Unwrap() *Token {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Token is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Token) String() string {
	var builder strings.Builder
	builder.WriteString("Token(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("token=")
	builder.WriteString(t.Token)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", t.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Tokens is a parsable slice of Token.
type Tokens []*Token

func (t Tokens) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
