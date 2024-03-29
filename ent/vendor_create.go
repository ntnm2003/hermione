// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"golang-boilerplate/ent/item"
	"golang-boilerplate/ent/vendor"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// VendorCreate is the builder for creating a Vendor entity.
type VendorCreate struct {
	config
	mutation *VendorMutation
	hooks    []Hook
}

// SetVendor sets the "vendor" field.
func (vc *VendorCreate) SetVendor(s string) *VendorCreate {
	vc.mutation.SetVendor(s)
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VendorCreate) SetCreatedAt(t time.Time) *VendorCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VendorCreate) SetNillableCreatedAt(t *time.Time) *VendorCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VendorCreate) SetUpdatedAt(t time.Time) *VendorCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VendorCreate) SetNillableUpdatedAt(t *time.Time) *VendorCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetID sets the "id" field.
func (vc *VendorCreate) SetID(u uuid.UUID) *VendorCreate {
	vc.mutation.SetID(u)
	return vc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vc *VendorCreate) SetNillableID(u *uuid.UUID) *VendorCreate {
	if u != nil {
		vc.SetID(*u)
	}
	return vc
}

// AddItemIDs adds the "item" edge to the Item entity by IDs.
func (vc *VendorCreate) AddItemIDs(ids ...uuid.UUID) *VendorCreate {
	vc.mutation.AddItemIDs(ids...)
	return vc
}

// AddItem adds the "item" edges to the Item entity.
func (vc *VendorCreate) AddItem(i ...*Item) *VendorCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return vc.AddItemIDs(ids...)
}

// Mutation returns the VendorMutation object of the builder.
func (vc *VendorCreate) Mutation() *VendorMutation {
	return vc.mutation
}

// Save creates the Vendor in the database.
func (vc *VendorCreate) Save(ctx context.Context) (*Vendor, error) {
	var (
		err  error
		node *Vendor
	)
	vc.defaults()
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VendorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, vc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Vendor)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from VendorMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VendorCreate) SaveX(ctx context.Context) *Vendor {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VendorCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VendorCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VendorCreate) defaults() {
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := vendor.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := vendor.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vc.mutation.ID(); !ok {
		v := vendor.DefaultID()
		vc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VendorCreate) check() error {
	if _, ok := vc.mutation.Vendor(); !ok {
		return &ValidationError{Name: "vendor", err: errors.New(`ent: missing required field "Vendor.vendor"`)}
	}
	if v, ok := vc.mutation.Vendor(); ok {
		if err := vendor.VendorValidator(v); err != nil {
			return &ValidationError{Name: "vendor", err: fmt.Errorf(`ent: validator failed for field "Vendor.vendor": %w`, err)}
		}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Vendor.created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Vendor.updated_at"`)}
	}
	return nil
}

func (vc *VendorCreate) sqlSave(ctx context.Context) (*Vendor, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (vc *VendorCreate) createSpec() (*Vendor, *sqlgraph.CreateSpec) {
	var (
		_node = &Vendor{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vendor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vendor.FieldID,
			},
		}
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := vc.mutation.Vendor(); ok {
		_spec.SetField(vendor.FieldVendor, field.TypeString, value)
		_node.Vendor = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(vendor.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.SetField(vendor.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := vc.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vendor.ItemTable,
			Columns: []string{vendor.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VendorCreateBulk is the builder for creating many Vendor entities in bulk.
type VendorCreateBulk struct {
	config
	builders []*VendorCreate
}

// Save creates the Vendor entities in the database.
func (vcb *VendorCreateBulk) Save(ctx context.Context) ([]*Vendor, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vendor, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VendorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VendorCreateBulk) SaveX(ctx context.Context) []*Vendor {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VendorCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VendorCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
