package vendor

import (
	"context"
	"golang-boilerplate/model"
	"golang-boilerplate/ent"
	"github.com/pkg/errors"
	"golang-boilerplate/ent/vendor"
	"github.com/google/uuid"
)

type Repository interface {

	Create(ctx context.Context, input model.VendorInput ) (*ent.Vendor, error)
	List(ctx context.Context, id *string) ([]*ent.Vendor, error)
}
type impl struct {
	client *ent.Client
}

func (i impl) Create(ctx context.Context, input model.VendorInput) (*ent.Vendor, error) {
	v, err := i.client.Vendor.Create().SetVendor(input.Vendor).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return v, nil
}

func (i impl) List(ctx context.Context, id *string) ([]*ent.Vendor, error) {

	query := i.client.Vendor.Query()

		if *id == "0" {
			vendors, err := query.All(ctx)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			return vendors, nil
		}
		vendorId, err := uuid.Parse(*id)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		vendors, err := query.Where(vendor.ID(vendorId)).All(ctx)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		return vendors, nil
}

func New (client *ent.Client) Repository{
	 return &impl{
		 client: client,
	 }
 }