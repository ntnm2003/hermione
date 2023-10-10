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

// Token is the resolver for the token field.
func (r *mutationResolver) Token(ctx context.Context) (*model.TokenOps, error) {
	panic(fmt.Errorf("not implemented: Token - token"))
}

// ID is the resolver for the id field.
func (r *tokenResolver) ID(ctx context.Context, obj *ent.Token) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// AccessToken is the resolver for the accessToken field.
func (r *tokenResolver) AccessToken(ctx context.Context, obj *ent.Token) (string, error) {
	panic(fmt.Errorf("not implemented: AccessToken - accessToken"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *tokenResolver) RefreshToken(ctx context.Context, obj *ent.Token) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// User is the resolver for the user field.
func (r *tokenResolver) User(ctx context.Context, obj *ent.Token) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// ExchangeToken is the resolver for the exchangeToken field.
func (r *tokenOpsResolver) ExchangeToken(ctx context.Context, obj *model.TokenOps, refreshToken string) (*ent.Token, error) {
	panic(fmt.Errorf("not implemented: ExchangeToken - exchangeToken"))
}

// Logout is the resolver for the logout field.
func (r *tokenOpsResolver) Logout(ctx context.Context, obj *model.TokenOps, refreshToken string) (*ent.Token, error) {
	panic(fmt.Errorf("not implemented: Logout - logout"))
}

// Token returns graphql1.TokenResolver implementation.
func (r *Resolver) Token() graphql1.TokenResolver { return &tokenResolver{r} }

// TokenOps returns graphql1.TokenOpsResolver implementation.
func (r *Resolver) TokenOps() graphql1.TokenOpsResolver { return &tokenOpsResolver{r} }

type tokenResolver struct{ *Resolver }
type tokenOpsResolver struct{ *Resolver }
