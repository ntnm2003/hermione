// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"golang-boilerplate/ent/item"
	"golang-boilerplate/ent/predicate"
	"golang-boilerplate/ent/vendor"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// VendorQuery is the builder for querying Vendor entities.
type VendorQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.Vendor
	withItem      *ItemQuery
	modifiers     []func(*sql.Selector)
	loadTotal     []func(context.Context, []*Vendor) error
	withNamedItem map[string]*ItemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VendorQuery builder.
func (vq *VendorQuery) Where(ps ...predicate.Vendor) *VendorQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit adds a limit step to the query.
func (vq *VendorQuery) Limit(limit int) *VendorQuery {
	vq.limit = &limit
	return vq
}

// Offset adds an offset step to the query.
func (vq *VendorQuery) Offset(offset int) *VendorQuery {
	vq.offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VendorQuery) Unique(unique bool) *VendorQuery {
	vq.unique = &unique
	return vq
}

// Order adds an order step to the query.
func (vq *VendorQuery) Order(o ...OrderFunc) *VendorQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryItem chains the current query on the "item" edge.
func (vq *VendorQuery) QueryItem() *ItemQuery {
	query := &ItemQuery{config: vq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vendor.Table, vendor.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, vendor.ItemTable, vendor.ItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Vendor entity from the query.
// Returns a *NotFoundError when no Vendor was found.
func (vq *VendorQuery) First(ctx context.Context) (*Vendor, error) {
	nodes, err := vq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vendor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VendorQuery) FirstX(ctx context.Context) *Vendor {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Vendor ID from the query.
// Returns a *NotFoundError when no Vendor ID was found.
func (vq *VendorQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vendor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VendorQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Vendor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Vendor entity is found.
// Returns a *NotFoundError when no Vendor entities are found.
func (vq *VendorQuery) Only(ctx context.Context) (*Vendor, error) {
	nodes, err := vq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vendor.Label}
	default:
		return nil, &NotSingularError{vendor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VendorQuery) OnlyX(ctx context.Context) *Vendor {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Vendor ID in the query.
// Returns a *NotSingularError when more than one Vendor ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VendorQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vendor.Label}
	default:
		err = &NotSingularError{vendor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VendorQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Vendors.
func (vq *VendorQuery) All(ctx context.Context) ([]*Vendor, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vq *VendorQuery) AllX(ctx context.Context) []*Vendor {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Vendor IDs.
func (vq *VendorQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := vq.Select(vendor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VendorQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VendorQuery) Count(ctx context.Context) (int, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VendorQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VendorQuery) Exist(ctx context.Context) (bool, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VendorQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VendorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VendorQuery) Clone() *VendorQuery {
	if vq == nil {
		return nil
	}
	return &VendorQuery{
		config:     vq.config,
		limit:      vq.limit,
		offset:     vq.offset,
		order:      append([]OrderFunc{}, vq.order...),
		predicates: append([]predicate.Vendor{}, vq.predicates...),
		withItem:   vq.withItem.Clone(),
		// clone intermediate query.
		sql:    vq.sql.Clone(),
		path:   vq.path,
		unique: vq.unique,
	}
}

// WithItem tells the query-builder to eager-load the nodes that are connected to
// the "item" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VendorQuery) WithItem(opts ...func(*ItemQuery)) *VendorQuery {
	query := &ItemQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	vq.withItem = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Vendor string `json:"vendor,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Vendor.Query().
//		GroupBy(vendor.FieldVendor).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vq *VendorQuery) GroupBy(field string, fields ...string) *VendorGroupBy {
	grbuild := &VendorGroupBy{config: vq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vq.sqlQuery(ctx), nil
	}
	grbuild.label = vendor.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Vendor string `json:"vendor,omitempty"`
//	}
//
//	client.Vendor.Query().
//		Select(vendor.FieldVendor).
//		Scan(ctx, &v)
//
func (vq *VendorQuery) Select(fields ...string) *VendorSelect {
	vq.fields = append(vq.fields, fields...)
	selbuild := &VendorSelect{VendorQuery: vq}
	selbuild.label = vendor.Label
	selbuild.flds, selbuild.scan = &vq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a VendorSelect configured with the given aggregations.
func (vq *VendorQuery) Aggregate(fns ...AggregateFunc) *VendorSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VendorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vq.fields {
		if !vendor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VendorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Vendor, error) {
	var (
		nodes       = []*Vendor{}
		_spec       = vq.querySpec()
		loadedTypes = [1]bool{
			vq.withItem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Vendor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Vendor{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withItem; query != nil {
		if err := vq.loadItem(ctx, query, nodes,
			func(n *Vendor) { n.Edges.Item = []*Item{} },
			func(n *Vendor, e *Item) { n.Edges.Item = append(n.Edges.Item, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range vq.withNamedItem {
		if err := vq.loadItem(ctx, query, nodes,
			func(n *Vendor) { n.appendNamedItem(name) },
			func(n *Vendor, e *Item) { n.appendNamedItem(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range vq.loadTotal {
		if err := vq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VendorQuery) loadItem(ctx context.Context, query *ItemQuery, nodes []*Vendor, init func(*Vendor), assign func(*Vendor, *Item)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Vendor)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Item(func(s *sql.Selector) {
		s.Where(sql.InValues(vendor.ItemColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.VendorID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "vendor_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (vq *VendorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	_spec.Node.Columns = vq.fields
	if len(vq.fields) > 0 {
		_spec.Unique = vq.unique != nil && *vq.unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VendorQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (vq *VendorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vendor.Table,
			Columns: vendor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vendor.FieldID,
			},
		},
		From:   vq.sql,
		Unique: true,
	}
	if unique := vq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vendor.FieldID)
		for i := range fields {
			if fields[i] != vendor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VendorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(vendor.Table)
	columns := vq.fields
	if len(columns) == 0 {
		columns = vendor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.unique != nil && *vq.unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedItem tells the query-builder to eager-load the nodes that are connected to the "item"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (vq *VendorQuery) WithNamedItem(name string, opts ...func(*ItemQuery)) *VendorQuery {
	query := &ItemQuery{config: vq.config}
	for _, opt := range opts {
		opt(query)
	}
	if vq.withNamedItem == nil {
		vq.withNamedItem = make(map[string]*ItemQuery)
	}
	vq.withNamedItem[name] = query
	return vq
}

// VendorGroupBy is the group-by builder for Vendor entities.
type VendorGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VendorGroupBy) Aggregate(fns ...AggregateFunc) *VendorGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vgb *VendorGroupBy) Scan(ctx context.Context, v any) error {
	query, err := vgb.path(ctx)
	if err != nil {
		return err
	}
	vgb.sql = query
	return vgb.sqlScan(ctx, v)
}

func (vgb *VendorGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range vgb.fields {
		if !vendor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vgb *VendorGroupBy) sqlQuery() *sql.Selector {
	selector := vgb.sql.Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vgb.fields)+len(vgb.fns))
		for _, f := range vgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vgb.fields...)...)
}

// VendorSelect is the builder for selecting fields of Vendor entities.
type VendorSelect struct {
	*VendorQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VendorSelect) Aggregate(fns ...AggregateFunc) *VendorSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VendorSelect) Scan(ctx context.Context, v any) error {
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	vs.sql = vs.VendorQuery.sqlQuery(ctx)
	return vs.sqlScan(ctx, v)
}

func (vs *VendorSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(vs.sql))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		vs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		vs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := vs.sql.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
