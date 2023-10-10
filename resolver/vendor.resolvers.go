package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang-boilerplate/ent"
	graphql1 "golang-boilerplate/graphql"
	"golang-boilerplate/model"
)

// Create is the resolver for the create field.
func (r *mutationResolver) Create(ctx context.Context, input model.VendorInput) (*ent.Vendor, error) {
	panic(fmt.Errorf("not implemented: Create - create"))
}

// List is the resolver for the list field.
func (r *queryResolver) List(ctx context.Context, id string) ([]*ent.Vendor, error) {
	panic(fmt.Errorf("not implemented: List - list"))
}

// ID is the resolver for the id field.
func (r *vendorResolver) ID(ctx context.Context, obj *ent.Vendor) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Vendor returns graphql1.VendorResolver implementation.
func (r *Resolver) Vendor() graphql1.VendorResolver { return &vendorResolver{r} }

type vendorResolver struct{ *Resolver }
