package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	graphql1 "golang-boilerplate/graphql"
	"golang-boilerplate/model"
)

// Create is the resolver for the create field.
func (r *itemOpsResolver) Create(ctx context.Context, obj *model.ItemOps, input model.CreateUserInput) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: Create - create"))
}

// Buy is the resolver for the buy field.
func (r *itemOpsResolver) Buy(ctx context.Context, obj *model.ItemOps, input model.BuyItemInput) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: Buy - buy"))
}

// List is the resolver for the list field.
func (r *itemQueriesResolver) List(ctx context.Context, obj *model.ItemQueries) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: List - list"))
}

// TopList is the resolver for the topList field.
func (r *itemQueriesResolver) TopList(ctx context.Context, obj *model.ItemQueries) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: TopList - topList"))
}

// ExpList is the resolver for the expList field.
func (r *itemQueriesResolver) ExpList(ctx context.Context, obj *model.ItemQueries) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: ExpList - expList"))
}

// Revenue is the resolver for the revenue field.
func (r *itemQueriesResolver) Revenue(ctx context.Context, obj *model.ItemQueries) (*int, error) {
	panic(fmt.Errorf("not implemented: Revenue - revenue"))
}

// Items is the resolver for the items field.
func (r *mutationResolver) Items(ctx context.Context) (*model.ItemOps, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Item is the resolver for the item field.
func (r *queryResolver) Item(ctx context.Context) (*model.ItemQueries, error) {
	panic(fmt.Errorf("not implemented: Item - item"))
}

// ItemOps returns graphql1.ItemOpsResolver implementation.
func (r *Resolver) ItemOps() graphql1.ItemOpsResolver { return &itemOpsResolver{r} }

// ItemQueries returns graphql1.ItemQueriesResolver implementation.
func (r *Resolver) ItemQueries() graphql1.ItemQueriesResolver { return &itemQueriesResolver{r} }

type itemOpsResolver struct{ *Resolver }
type itemQueriesResolver struct{ *Resolver }
