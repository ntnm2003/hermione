// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"golang-boilerplate/ent/item"
	"golang-boilerplate/ent/predicate"
	"golang-boilerplate/ent/vendor"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetItem sets the "item" field.
func (iu *ItemUpdate) SetItem(s string) *ItemUpdate {
	iu.mutation.SetItem(s)
	return iu
}

// SetPrice sets the "price" field.
func (iu *ItemUpdate) SetPrice(i int) *ItemUpdate {
	iu.mutation.ResetPrice()
	iu.mutation.SetPrice(i)
	return iu
}

// AddPrice adds i to the "price" field.
func (iu *ItemUpdate) AddPrice(i int) *ItemUpdate {
	iu.mutation.AddPrice(i)
	return iu
}

// SetRemainingAmount sets the "remaining_amount" field.
func (iu *ItemUpdate) SetRemainingAmount(i int) *ItemUpdate {
	iu.mutation.ResetRemainingAmount()
	iu.mutation.SetRemainingAmount(i)
	return iu
}

// AddRemainingAmount adds i to the "remaining_amount" field.
func (iu *ItemUpdate) AddRemainingAmount(i int) *ItemUpdate {
	iu.mutation.AddRemainingAmount(i)
	return iu
}

// SetSoldAmount sets the "sold_amount" field.
func (iu *ItemUpdate) SetSoldAmount(i int) *ItemUpdate {
	iu.mutation.ResetSoldAmount()
	iu.mutation.SetSoldAmount(i)
	return iu
}

// AddSoldAmount adds i to the "sold_amount" field.
func (iu *ItemUpdate) AddSoldAmount(i int) *ItemUpdate {
	iu.mutation.AddSoldAmount(i)
	return iu
}

// SetExp sets the "exp" field.
func (iu *ItemUpdate) SetExp(t time.Time) *ItemUpdate {
	iu.mutation.SetExp(t)
	return iu
}

// SetVendorID sets the "vendor_id" field.
func (iu *ItemUpdate) SetVendorID(u uuid.UUID) *ItemUpdate {
	iu.mutation.SetVendorID(u)
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *ItemUpdate) SetUpdatedAt(t time.Time) *ItemUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetVendorsID sets the "vendors" edge to the Vendor entity by ID.
func (iu *ItemUpdate) SetVendorsID(id uuid.UUID) *ItemUpdate {
	iu.mutation.SetVendorsID(id)
	return iu
}

// SetVendors sets the "vendors" edge to the Vendor entity.
func (iu *ItemUpdate) SetVendors(v *Vendor) *ItemUpdate {
	return iu.SetVendorsID(v.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// ClearVendors clears the "vendors" edge to the Vendor entity.
func (iu *ItemUpdate) ClearVendors() *ItemUpdate {
	iu.mutation.ClearVendors()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *ItemUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := item.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ItemUpdate) check() error {
	if v, ok := iu.mutation.Item(); ok {
		if err := item.ItemValidator(v); err != nil {
			return &ValidationError{Name: "item", err: fmt.Errorf(`ent: validator failed for field "Item.item": %w`, err)}
		}
	}
	if _, ok := iu.mutation.VendorsID(); iu.mutation.VendorsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Item.vendors"`)
	}
	return nil
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: item.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Item(); ok {
		_spec.SetField(item.FieldItem, field.TypeString, value)
	}
	if value, ok := iu.mutation.Price(); ok {
		_spec.SetField(item.FieldPrice, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedPrice(); ok {
		_spec.AddField(item.FieldPrice, field.TypeInt, value)
	}
	if value, ok := iu.mutation.RemainingAmount(); ok {
		_spec.SetField(item.FieldRemainingAmount, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedRemainingAmount(); ok {
		_spec.AddField(item.FieldRemainingAmount, field.TypeInt, value)
	}
	if value, ok := iu.mutation.SoldAmount(); ok {
		_spec.SetField(item.FieldSoldAmount, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedSoldAmount(); ok {
		_spec.AddField(item.FieldSoldAmount, field.TypeInt, value)
	}
	if value, ok := iu.mutation.Exp(); ok {
		_spec.SetField(item.FieldExp, field.TypeTime, value)
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(item.FieldUpdatedAt, field.TypeTime, value)
	}
	if iu.mutation.VendorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.VendorsTable,
			Columns: []string{item.VendorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vendor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.VendorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.VendorsTable,
			Columns: []string{item.VendorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vendor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemMutation
}

// SetItem sets the "item" field.
func (iuo *ItemUpdateOne) SetItem(s string) *ItemUpdateOne {
	iuo.mutation.SetItem(s)
	return iuo
}

// SetPrice sets the "price" field.
func (iuo *ItemUpdateOne) SetPrice(i int) *ItemUpdateOne {
	iuo.mutation.ResetPrice()
	iuo.mutation.SetPrice(i)
	return iuo
}

// AddPrice adds i to the "price" field.
func (iuo *ItemUpdateOne) AddPrice(i int) *ItemUpdateOne {
	iuo.mutation.AddPrice(i)
	return iuo
}

// SetRemainingAmount sets the "remaining_amount" field.
func (iuo *ItemUpdateOne) SetRemainingAmount(i int) *ItemUpdateOne {
	iuo.mutation.ResetRemainingAmount()
	iuo.mutation.SetRemainingAmount(i)
	return iuo
}

// AddRemainingAmount adds i to the "remaining_amount" field.
func (iuo *ItemUpdateOne) AddRemainingAmount(i int) *ItemUpdateOne {
	iuo.mutation.AddRemainingAmount(i)
	return iuo
}

// SetSoldAmount sets the "sold_amount" field.
func (iuo *ItemUpdateOne) SetSoldAmount(i int) *ItemUpdateOne {
	iuo.mutation.ResetSoldAmount()
	iuo.mutation.SetSoldAmount(i)
	return iuo
}

// AddSoldAmount adds i to the "sold_amount" field.
func (iuo *ItemUpdateOne) AddSoldAmount(i int) *ItemUpdateOne {
	iuo.mutation.AddSoldAmount(i)
	return iuo
}

// SetExp sets the "exp" field.
func (iuo *ItemUpdateOne) SetExp(t time.Time) *ItemUpdateOne {
	iuo.mutation.SetExp(t)
	return iuo
}

// SetVendorID sets the "vendor_id" field.
func (iuo *ItemUpdateOne) SetVendorID(u uuid.UUID) *ItemUpdateOne {
	iuo.mutation.SetVendorID(u)
	return iuo
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *ItemUpdateOne) SetUpdatedAt(t time.Time) *ItemUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetVendorsID sets the "vendors" edge to the Vendor entity by ID.
func (iuo *ItemUpdateOne) SetVendorsID(id uuid.UUID) *ItemUpdateOne {
	iuo.mutation.SetVendorsID(id)
	return iuo
}

// SetVendors sets the "vendors" edge to the Vendor entity.
func (iuo *ItemUpdateOne) SetVendors(v *Vendor) *ItemUpdateOne {
	return iuo.SetVendorsID(v.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// ClearVendors clears the "vendors" edge to the Vendor entity.
func (iuo *ItemUpdateOne) ClearVendors() *ItemUpdateOne {
	iuo.mutation.ClearVendors()
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	var (
		err  error
		node *Item
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, iuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Item)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ItemMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *ItemUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := item.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ItemUpdateOne) check() error {
	if v, ok := iuo.mutation.Item(); ok {
		if err := item.ItemValidator(v); err != nil {
			return &ValidationError{Name: "item", err: fmt.Errorf(`ent: validator failed for field "Item.item": %w`, err)}
		}
	}
	if _, ok := iuo.mutation.VendorsID(); iuo.mutation.VendorsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Item.vendors"`)
	}
	return nil
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: item.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Item(); ok {
		_spec.SetField(item.FieldItem, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Price(); ok {
		_spec.SetField(item.FieldPrice, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedPrice(); ok {
		_spec.AddField(item.FieldPrice, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.RemainingAmount(); ok {
		_spec.SetField(item.FieldRemainingAmount, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedRemainingAmount(); ok {
		_spec.AddField(item.FieldRemainingAmount, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.SoldAmount(); ok {
		_spec.SetField(item.FieldSoldAmount, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedSoldAmount(); ok {
		_spec.AddField(item.FieldSoldAmount, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.Exp(); ok {
		_spec.SetField(item.FieldExp, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(item.FieldUpdatedAt, field.TypeTime, value)
	}
	if iuo.mutation.VendorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.VendorsTable,
			Columns: []string{item.VendorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vendor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.VendorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.VendorsTable,
			Columns: []string{item.VendorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vendor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
