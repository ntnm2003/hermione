package user

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/ent/user"
	"golang-boilerplate/model"
	"strings"

	"github.com/pkg/errors"
)

// Repository is the interface for the user repository.
type Repository interface {
	// Create creates a new user.
	Create(ctx context.Context, input model.CreateUserInput) (*ent.User, error)
	// List lists all users by filter.
	List(ctx context.Context, filter ent.UserFilterInput) (*ent.UserConnection, error)
}

// impl is the implementation of the user repository.
type impl struct {
	client *ent.Client
}

// New creates a new user repository.
func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}

// Create creates a new user.
func (i impl) Create(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	user, err := i.client.User.Create().SetFullName(input.FullName).SetUsername(input.Username).SetPassword(input.Password).
		SetEmail(input.Email).SetRole(user.Role(input.Role)).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}

// List lists all users by filter.
func (i impl) List(ctx context.Context, filter ent.UserFilterInput) (*ent.UserConnection, error) {
	query := i.client.User.Query()
	if filter.Name != nil {
		query = query.Where(user.FullNameContainsFold(strings.TrimSpace(*filter.Name)))
	}

	users, err := query.Paginate(ctx, filter.Pagination.After, filter.Pagination.First, filter.Pagination.Before, filter.Pagination.Last)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return users, nil
}
