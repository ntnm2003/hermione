package vendor

import (
"context"
"errors"
"golang-boilerplate/ent"
"golang-boilerplate/internal/util"
"golang-boilerplate/model"
"golang-boilerplate/repository"

"go.uber.org/zap"
)

// Service is the interface for the user service.
type Service interface {
	Create(ctx context.Context, input model.VendorInput) (*ent.Vendor, error)
	List(ctx context.Context, id *string) ([]*ent.Vendor, error)
}
// impl is the implementation of the user service.
type impl struct {
	repoRegistry repository.RepositoryRegistry
	logger       *zap.Logger
}

func (i impl) Create(ctx context.Context, input model.VendorInput) (*ent.Vendor, error) {
	vendor, err := i.repoRegistry.Vendor().Create(ctx, input)
	if err != nil {
		i.logger.Error("Create_vendor_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(ctx, errors.Unwrap(err).Error(), "Create_vendor_items_has_failed")
	}
	return vendor, nil
}

func (i impl) List(ctx context.Context, id *string) ([]*ent.Vendor, error) {
	vendors, err := i.repoRegistry.Vendor().List(ctx, id)
	if err != nil {
		i.logger.Error("List_vendor_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(ctx, errors.Unwrap(err).Error(), "List_vendor_items_has_failed")
	}
	return vendors, nil
}

// New creates a new user service.
func New(repoRegistry repository.RepositoryRegistry, logger *zap.Logger) Service {
	return &impl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

