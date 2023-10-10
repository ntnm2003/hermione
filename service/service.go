package service

import (
	"golang-boilerplate/ent"
	"golang-boilerplate/repository"

	"golang-boilerplate/service/item"

	"golang-boilerplate/service/user"
	"golang-boilerplate/service/vendor"

	"go.uber.org/zap"
)

// ServiceRegistry is the interface for the service registry.
type ServiceRegistry interface {
	User() user.Service
	Vendor() vendor.Service
	Item() Item.Service
}

// impl is the implementation of the service registry.
type impl struct {
	user   user.Service
	vendor vendor.Service
	item   Item.Service
}

func (i impl) Vendor() vendor.Service {
	return i.vendor
}

func (i impl) Item() Item.Service {
	return i.item
}

// New creates a new service registry.
func New(entClient *ent.Client, logger *zap.Logger) ServiceRegistry {
	repoRegistry := repository.New(entClient)

	return &impl{
		user:   user.New(repoRegistry, logger),
		item:   Item.New(repoRegistry, logger),
		vendor: vendor.New(repoRegistry, logger),
	}
}

// User returns the user service.
func (i impl) User() user.Service {
	return i.user
}
