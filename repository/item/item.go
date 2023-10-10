package item

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang-boilerplate/ent"
	"golang-boilerplate/ent/item"
	"golang-boilerplate/ent/vendor"
	"golang-boilerplate/model"
	"sort"
	"time"
)

type Repository interface {
	Create(ctx context.Context, input model.CreateItemInput) (*ent.Item, error)
	List(ctx context.Context) ([]*ent.Item, error)
	TopList(context context.Context) ([]*ent.Item, error)
	ExpList(context context.Context) ([]*ent.Item, error)
	Revenue(context context.Context, id *string) (*int, error)
	Buy(context context.Context, input model.BuyItemInput) (*ent.Item, error)
}
type impl struct {
	client *ent.Client
}

func (i impl) Buy(context context.Context, input model.BuyItemInput) (*ent.Item, error) {
	query := i.client.Item.Query()

	itemId, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	item, err := query.Where(item.ID(itemId)).Only(context)
	if item.Item != input.Item {
		return nil, fmt.Errorf("itemId not found")
	}
	sold := item.SoldAmount
	it, err := item.Update().SetSoldAmount(sold - input.SoldAmount).Save(context)
	return it, nil
}

func (i impl) Create(ctx context.Context, input model.CreateItemInput) (*ent.Item, error) {
	vendorId, err := uuid.Parse(input.VendorID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	vendor, err := i.client.Vendor.Query().Where(vendor.ID(vendorId)).Only(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if vendor == nil {
		return nil, errors.Errorf("vendor_not_exist")
	}
	item, err := i.client.Item.Create().SetItem(input.Item).
		SetPrice(input.Price).
		SetExp(input.Exp).
		SetRemainingAmount(input.RemainingAmount).
		SetSoldAmount(0).
		SetVendor(vendor).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return item, nil
}

func (i impl) List(ctx context.Context) ([]*ent.Item, error) {
	query := i.client.Item.Query()
	items, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return items, nil
}

func (i impl) TopList(context context.Context) ([]*ent.Item, error) {
	query := i.client.Item.Query()
	items, err := query.All(context)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].SoldAmount > items[j].SoldAmount
	})
	var items1 []*ent.Item
	for k := 0; k < 2; k++ {
		items1 = append(items1, items[k])
	}
	return items1, nil
}

func (i impl) ExpList(context context.Context) ([]*ent.Item, error) {
	query := i.client.Item.Query()
	items, err := query.All(context)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	during10Days := 10 * 24 * time.Hour
	beforeTime := time.Now().Add(-during10Days)
	var items1 []*ent.Item
	for _, item := range items {
		if item.Exp.After(beforeTime) {
			items1 = append(items1, item)
		}
	}
	return items1, nil
}

func (i impl) Revenue(context context.Context, id *string) (*int, error) {
	query := i.client.Item.Query()

	var items1 []*ent.Item
	if *id == "0" {
		items, err := query.All(context)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		items1 = items
	}
	itemId, err := uuid.Parse(*id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	items1, err = query.Where(item.ID(itemId)).All(context)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	sum := 0
	for _, item := range items1 {
		sum += (item.Price) * (item.SoldAmount)
	}
	return &sum, nil
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
