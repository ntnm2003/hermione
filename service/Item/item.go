package Item

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang-boilerplate/ent"
	"golang-boilerplate/internal/util"
	"golang-boilerplate/model"
	"golang-boilerplate/repository"
)

// Service is the interface for the user service.
type Service interface {
	// Create creates a new user.
	Create(ctx context.Context, input model.CreateItemInput) (*ent.Item, error)
	List(ctx context.Context) ([]*ent.Item, error)
	TopList(context context.Context) ([]*ent.Item, error)
	ExpList(context context.Context) ([]*ent.Item, error)
	Revenue(context context.Context, id *string) (*int, error)
	Buy(context context.Context, input model.BuyItemInput) (*ent.Item, error)
}

// impl is the implementation of the user service.
type impl struct {
	repoRegistry repository.RepositoryRegistry
	logger       *zap.Logger
}

func (i impl) Create(ctx context.Context, input model.CreateItemInput) (*ent.Item, error) {
	item, err := i.repoRegistry.Item().Create(ctx, input)
	if err != nil {
		i.logger.Error("create_item_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(ctx, errors.Unwrap(err).Error(), "create_item_has_failed")
	}

	return item, nil
}

func (i impl) List(ctx context.Context) ([]*ent.Item, error) {
	items, err := i.repoRegistry.Item().List(ctx)
	if err != nil {
		i.logger.Error("List_items_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(ctx, errors.Unwrap(err).Error(), "List_items_has_failed")
	}

	return items, nil
}

func (i impl) TopList(context context.Context) ([]*ent.Item, error) {
	items, err := i.repoRegistry.Item().TopList(context)
	if err != nil {
		i.logger.Error("List_top_items_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(context, errors.Unwrap(err).Error(), "List_top_items_has_failed")
	}

	return items, nil
}

func (i impl) ExpList(context context.Context) ([]*ent.Item, error) {
	items, err := i.repoRegistry.Item().ExpList(context)
	if err != nil {
		i.logger.Error("List_items_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(context, errors.Unwrap(err).Error(), "Lists_items_has_failed")
	}

	return items, nil
}

func (i impl) Revenue(context context.Context, id *string) (*int, error) {
	revenue, err := i.repoRegistry.Item().Revenue(context, id)
	if err != nil {
		i.logger.Error("Get_Revenue_items_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(context, errors.Unwrap(err).Error(), "Get_Revenue_items_has_failed")
	}
	return revenue, nil
}

func (i impl) Buy(context context.Context, input model.BuyItemInput) (*ent.Item, error) {
	items, err := i.repoRegistry.Item().Buy(context, input)
	if err != nil {
		i.logger.Error("Buy_items_has_failed", zap.Error(err))
		return nil, util.WrapGQLError(context, errors.Unwrap(err).Error(), "Buy_items_has_failed")
	}

	return items, nil
}

// New Lists a new user service.
func New(repoRegistry repository.RepositoryRegistry, logger *zap.Logger) Service {
	return &impl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}
