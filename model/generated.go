// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"golang-boilerplate/ent"
	"io"
	"strconv"
	"time"
)

type BuyItemInput struct {
	ID         string `json:"id"`
	Item       string `json:"item"`
	SoldAmount int    `json:"sold_amount"`
}

type CreateItemInput struct {
	Item            string    `json:"item"`
	Price           int       `json:"price"`
	RemainingAmount int       `json:"remainingAmount"`
	Exp             time.Time `json:"exp"`
	VendorID        string    `json:"vendorID"`
}

type CreateUserInput struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	FullName string   `json:"full_name"`
	Email    string   `json:"email"`
	Role     RoleType `json:"role"`
}

type GetTokenInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ItemOps struct {
	Create *ent.Item `json:"create"`
	Buy    *ent.Item `json:"buy"`
}

type ItemQueries struct {
	List    []*ent.Item `json:"list"`
	TopList []*ent.Item `json:"topList,omitempty"`
	ExpList []*ent.Item `json:"expList,omitempty"`
	Revenue *int        `json:"revenue,omitempty"`
}

type TokenOps struct {
	GetToken      *ent.Token `json:"getToken"`
	ExchangeToken *ent.Token `json:"exchangeToken"`
	Logout        *ent.Token `json:"logout,omitempty"`
}

type UserQueries struct {
	List *ent.UserConnection `json:"list"`
}

type VendorInput struct {
	Vendor string `json:"vendor"`
}

type RoleType string

const (
	RoleTypeAdmin RoleType = "ADMIN"
	RoleTypeUser  RoleType = "USER"
)

var AllRoleType = []RoleType{
	RoleTypeAdmin,
	RoleTypeUser,
}

func (e RoleType) IsValid() bool {
	switch e {
	case RoleTypeAdmin, RoleTypeUser:
		return true
	}
	return false
}

func (e RoleType) String() string {
	return string(e)
}

func (e *RoleType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ROLE_TYPE", str)
	}
	return nil
}

func (e RoleType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
