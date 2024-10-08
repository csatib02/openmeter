// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoice"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoiceitem"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// BillingInvoiceItemQuery is the builder for querying BillingInvoiceItem entities.
type BillingInvoiceItemQuery struct {
	config
	ctx                *QueryContext
	order              []billinginvoiceitem.OrderOption
	inters             []Interceptor
	predicates         []predicate.BillingInvoiceItem
	withBillingInvoice *BillingInvoiceQuery
	modifiers          []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillingInvoiceItemQuery builder.
func (biiq *BillingInvoiceItemQuery) Where(ps ...predicate.BillingInvoiceItem) *BillingInvoiceItemQuery {
	biiq.predicates = append(biiq.predicates, ps...)
	return biiq
}

// Limit the number of records to be returned by this query.
func (biiq *BillingInvoiceItemQuery) Limit(limit int) *BillingInvoiceItemQuery {
	biiq.ctx.Limit = &limit
	return biiq
}

// Offset to start from.
func (biiq *BillingInvoiceItemQuery) Offset(offset int) *BillingInvoiceItemQuery {
	biiq.ctx.Offset = &offset
	return biiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (biiq *BillingInvoiceItemQuery) Unique(unique bool) *BillingInvoiceItemQuery {
	biiq.ctx.Unique = &unique
	return biiq
}

// Order specifies how the records should be ordered.
func (biiq *BillingInvoiceItemQuery) Order(o ...billinginvoiceitem.OrderOption) *BillingInvoiceItemQuery {
	biiq.order = append(biiq.order, o...)
	return biiq
}

// QueryBillingInvoice chains the current query on the "billing_invoice" edge.
func (biiq *BillingInvoiceItemQuery) QueryBillingInvoice() *BillingInvoiceQuery {
	query := (&BillingInvoiceClient{config: biiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := biiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := biiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(billinginvoiceitem.Table, billinginvoiceitem.FieldID, selector),
			sqlgraph.To(billinginvoice.Table, billinginvoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, billinginvoiceitem.BillingInvoiceTable, billinginvoiceitem.BillingInvoiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(biiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BillingInvoiceItem entity from the query.
// Returns a *NotFoundError when no BillingInvoiceItem was found.
func (biiq *BillingInvoiceItemQuery) First(ctx context.Context) (*BillingInvoiceItem, error) {
	nodes, err := biiq.Limit(1).All(setContextOp(ctx, biiq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{billinginvoiceitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (biiq *BillingInvoiceItemQuery) FirstX(ctx context.Context) *BillingInvoiceItem {
	node, err := biiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BillingInvoiceItem ID from the query.
// Returns a *NotFoundError when no BillingInvoiceItem ID was found.
func (biiq *BillingInvoiceItemQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = biiq.Limit(1).IDs(setContextOp(ctx, biiq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{billinginvoiceitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (biiq *BillingInvoiceItemQuery) FirstIDX(ctx context.Context) string {
	id, err := biiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BillingInvoiceItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BillingInvoiceItem entity is found.
// Returns a *NotFoundError when no BillingInvoiceItem entities are found.
func (biiq *BillingInvoiceItemQuery) Only(ctx context.Context) (*BillingInvoiceItem, error) {
	nodes, err := biiq.Limit(2).All(setContextOp(ctx, biiq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{billinginvoiceitem.Label}
	default:
		return nil, &NotSingularError{billinginvoiceitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (biiq *BillingInvoiceItemQuery) OnlyX(ctx context.Context) *BillingInvoiceItem {
	node, err := biiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BillingInvoiceItem ID in the query.
// Returns a *NotSingularError when more than one BillingInvoiceItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (biiq *BillingInvoiceItemQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = biiq.Limit(2).IDs(setContextOp(ctx, biiq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{billinginvoiceitem.Label}
	default:
		err = &NotSingularError{billinginvoiceitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (biiq *BillingInvoiceItemQuery) OnlyIDX(ctx context.Context) string {
	id, err := biiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BillingInvoiceItems.
func (biiq *BillingInvoiceItemQuery) All(ctx context.Context) ([]*BillingInvoiceItem, error) {
	ctx = setContextOp(ctx, biiq.ctx, ent.OpQueryAll)
	if err := biiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BillingInvoiceItem, *BillingInvoiceItemQuery]()
	return withInterceptors[[]*BillingInvoiceItem](ctx, biiq, qr, biiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (biiq *BillingInvoiceItemQuery) AllX(ctx context.Context) []*BillingInvoiceItem {
	nodes, err := biiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BillingInvoiceItem IDs.
func (biiq *BillingInvoiceItemQuery) IDs(ctx context.Context) (ids []string, err error) {
	if biiq.ctx.Unique == nil && biiq.path != nil {
		biiq.Unique(true)
	}
	ctx = setContextOp(ctx, biiq.ctx, ent.OpQueryIDs)
	if err = biiq.Select(billinginvoiceitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (biiq *BillingInvoiceItemQuery) IDsX(ctx context.Context) []string {
	ids, err := biiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (biiq *BillingInvoiceItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, biiq.ctx, ent.OpQueryCount)
	if err := biiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, biiq, querierCount[*BillingInvoiceItemQuery](), biiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (biiq *BillingInvoiceItemQuery) CountX(ctx context.Context) int {
	count, err := biiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (biiq *BillingInvoiceItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, biiq.ctx, ent.OpQueryExist)
	switch _, err := biiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (biiq *BillingInvoiceItemQuery) ExistX(ctx context.Context) bool {
	exist, err := biiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillingInvoiceItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (biiq *BillingInvoiceItemQuery) Clone() *BillingInvoiceItemQuery {
	if biiq == nil {
		return nil
	}
	return &BillingInvoiceItemQuery{
		config:             biiq.config,
		ctx:                biiq.ctx.Clone(),
		order:              append([]billinginvoiceitem.OrderOption{}, biiq.order...),
		inters:             append([]Interceptor{}, biiq.inters...),
		predicates:         append([]predicate.BillingInvoiceItem{}, biiq.predicates...),
		withBillingInvoice: biiq.withBillingInvoice.Clone(),
		// clone intermediate query.
		sql:  biiq.sql.Clone(),
		path: biiq.path,
	}
}

// WithBillingInvoice tells the query-builder to eager-load the nodes that are connected to
// the "billing_invoice" edge. The optional arguments are used to configure the query builder of the edge.
func (biiq *BillingInvoiceItemQuery) WithBillingInvoice(opts ...func(*BillingInvoiceQuery)) *BillingInvoiceItemQuery {
	query := (&BillingInvoiceClient{config: biiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	biiq.withBillingInvoice = query
	return biiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BillingInvoiceItem.Query().
//		GroupBy(billinginvoiceitem.FieldNamespace).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (biiq *BillingInvoiceItemQuery) GroupBy(field string, fields ...string) *BillingInvoiceItemGroupBy {
	biiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillingInvoiceItemGroupBy{build: biiq}
	grbuild.flds = &biiq.ctx.Fields
	grbuild.label = billinginvoiceitem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//	}
//
//	client.BillingInvoiceItem.Query().
//		Select(billinginvoiceitem.FieldNamespace).
//		Scan(ctx, &v)
func (biiq *BillingInvoiceItemQuery) Select(fields ...string) *BillingInvoiceItemSelect {
	biiq.ctx.Fields = append(biiq.ctx.Fields, fields...)
	sbuild := &BillingInvoiceItemSelect{BillingInvoiceItemQuery: biiq}
	sbuild.label = billinginvoiceitem.Label
	sbuild.flds, sbuild.scan = &biiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillingInvoiceItemSelect configured with the given aggregations.
func (biiq *BillingInvoiceItemQuery) Aggregate(fns ...AggregateFunc) *BillingInvoiceItemSelect {
	return biiq.Select().Aggregate(fns...)
}

func (biiq *BillingInvoiceItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range biiq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, biiq); err != nil {
				return err
			}
		}
	}
	for _, f := range biiq.ctx.Fields {
		if !billinginvoiceitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if biiq.path != nil {
		prev, err := biiq.path(ctx)
		if err != nil {
			return err
		}
		biiq.sql = prev
	}
	return nil
}

func (biiq *BillingInvoiceItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BillingInvoiceItem, error) {
	var (
		nodes       = []*BillingInvoiceItem{}
		_spec       = biiq.querySpec()
		loadedTypes = [1]bool{
			biiq.withBillingInvoice != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BillingInvoiceItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BillingInvoiceItem{config: biiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(biiq.modifiers) > 0 {
		_spec.Modifiers = biiq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, biiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := biiq.withBillingInvoice; query != nil {
		if err := biiq.loadBillingInvoice(ctx, query, nodes, nil,
			func(n *BillingInvoiceItem, e *BillingInvoice) { n.Edges.BillingInvoice = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (biiq *BillingInvoiceItemQuery) loadBillingInvoice(ctx context.Context, query *BillingInvoiceQuery, nodes []*BillingInvoiceItem, init func(*BillingInvoiceItem), assign func(*BillingInvoiceItem, *BillingInvoice)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*BillingInvoiceItem)
	for i := range nodes {
		fk := nodes[i].InvoiceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(billinginvoice.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "invoice_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (biiq *BillingInvoiceItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := biiq.querySpec()
	if len(biiq.modifiers) > 0 {
		_spec.Modifiers = biiq.modifiers
	}
	_spec.Node.Columns = biiq.ctx.Fields
	if len(biiq.ctx.Fields) > 0 {
		_spec.Unique = biiq.ctx.Unique != nil && *biiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, biiq.driver, _spec)
}

func (biiq *BillingInvoiceItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(billinginvoiceitem.Table, billinginvoiceitem.Columns, sqlgraph.NewFieldSpec(billinginvoiceitem.FieldID, field.TypeString))
	_spec.From = biiq.sql
	if unique := biiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if biiq.path != nil {
		_spec.Unique = true
	}
	if fields := biiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billinginvoiceitem.FieldID)
		for i := range fields {
			if fields[i] != billinginvoiceitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if biiq.withBillingInvoice != nil {
			_spec.Node.AddColumnOnce(billinginvoiceitem.FieldInvoiceID)
		}
	}
	if ps := biiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := biiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := biiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := biiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (biiq *BillingInvoiceItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(biiq.driver.Dialect())
	t1 := builder.Table(billinginvoiceitem.Table)
	columns := biiq.ctx.Fields
	if len(columns) == 0 {
		columns = billinginvoiceitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if biiq.sql != nil {
		selector = biiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if biiq.ctx.Unique != nil && *biiq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range biiq.modifiers {
		m(selector)
	}
	for _, p := range biiq.predicates {
		p(selector)
	}
	for _, p := range biiq.order {
		p(selector)
	}
	if offset := biiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := biiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (biiq *BillingInvoiceItemQuery) ForUpdate(opts ...sql.LockOption) *BillingInvoiceItemQuery {
	if biiq.driver.Dialect() == dialect.Postgres {
		biiq.Unique(false)
	}
	biiq.modifiers = append(biiq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return biiq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (biiq *BillingInvoiceItemQuery) ForShare(opts ...sql.LockOption) *BillingInvoiceItemQuery {
	if biiq.driver.Dialect() == dialect.Postgres {
		biiq.Unique(false)
	}
	biiq.modifiers = append(biiq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return biiq
}

// BillingInvoiceItemGroupBy is the group-by builder for BillingInvoiceItem entities.
type BillingInvoiceItemGroupBy struct {
	selector
	build *BillingInvoiceItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (biigb *BillingInvoiceItemGroupBy) Aggregate(fns ...AggregateFunc) *BillingInvoiceItemGroupBy {
	biigb.fns = append(biigb.fns, fns...)
	return biigb
}

// Scan applies the selector query and scans the result into the given value.
func (biigb *BillingInvoiceItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, biigb.build.ctx, ent.OpQueryGroupBy)
	if err := biigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingInvoiceItemQuery, *BillingInvoiceItemGroupBy](ctx, biigb.build, biigb, biigb.build.inters, v)
}

func (biigb *BillingInvoiceItemGroupBy) sqlScan(ctx context.Context, root *BillingInvoiceItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(biigb.fns))
	for _, fn := range biigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*biigb.flds)+len(biigb.fns))
		for _, f := range *biigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*biigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := biigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillingInvoiceItemSelect is the builder for selecting fields of BillingInvoiceItem entities.
type BillingInvoiceItemSelect struct {
	*BillingInvoiceItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (biis *BillingInvoiceItemSelect) Aggregate(fns ...AggregateFunc) *BillingInvoiceItemSelect {
	biis.fns = append(biis.fns, fns...)
	return biis
}

// Scan applies the selector query and scans the result into the given value.
func (biis *BillingInvoiceItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, biis.ctx, ent.OpQuerySelect)
	if err := biis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingInvoiceItemQuery, *BillingInvoiceItemSelect](ctx, biis.BillingInvoiceItemQuery, biis, biis.inters, v)
}

func (biis *BillingInvoiceItemSelect) sqlScan(ctx context.Context, root *BillingInvoiceItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(biis.fns))
	for _, fn := range biis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*biis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := biis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
