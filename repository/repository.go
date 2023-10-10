package repository

import (
	"golang-boilerplate/ent"
	"golang-boilerplate/repository/item"
	"golang-boilerplate/repository/user"
	"golang-boilerplate/repository/vendor"
)

// RepositoryRegistry is the interface for the repository registry.
type RepositoryRegistry interface {
	User() user.Repository
	Item() item.Repository
	Vendor() vendor.Repository
}

// impl is the implementation of the repository registry.
type impl struct {
	user   user.Repository
	item   item.Repository
	vendor vendor.Repository
}

func (i impl) Item() item.Repository {
	return i.item
}

func (i impl) Vendor() vendor.Repository {
	return i.vendor
}

// New creates a new repository registry.
func New(client *ent.Client) RepositoryRegistry {
	return &impl{
		user:   user.New(client),
		item:   item.New(client),
		vendor: vendor.New(client),
	}
}

// User returns the user repository.
func (i impl) User() user.Repository {
	return i.user
}
