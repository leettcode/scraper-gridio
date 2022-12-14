// Code generated by entc, DO NOT EDIT.

package codegen

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alexgtn/go-linkshort/tools/ent/codegen/nordpooldayahead"
	"github.com/alexgtn/go-linkshort/tools/ent/codegen/predicate"
)

// NordPoolDayAheadQuery is the builder for querying NordPoolDayAhead entities.
type NordPoolDayAheadQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NordPoolDayAhead
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NordPoolDayAheadQuery builder.
func (npdaq *NordPoolDayAheadQuery) Where(ps ...predicate.NordPoolDayAhead) *NordPoolDayAheadQuery {
	npdaq.predicates = append(npdaq.predicates, ps...)
	return npdaq
}

// Limit adds a limit step to the query.
func (npdaq *NordPoolDayAheadQuery) Limit(limit int) *NordPoolDayAheadQuery {
	npdaq.limit = &limit
	return npdaq
}

// Offset adds an offset step to the query.
func (npdaq *NordPoolDayAheadQuery) Offset(offset int) *NordPoolDayAheadQuery {
	npdaq.offset = &offset
	return npdaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (npdaq *NordPoolDayAheadQuery) Unique(unique bool) *NordPoolDayAheadQuery {
	npdaq.unique = &unique
	return npdaq
}

// Order adds an order step to the query.
func (npdaq *NordPoolDayAheadQuery) Order(o ...OrderFunc) *NordPoolDayAheadQuery {
	npdaq.order = append(npdaq.order, o...)
	return npdaq
}

// First returns the first NordPoolDayAhead entity from the query.
// Returns a *NotFoundError when no NordPoolDayAhead was found.
func (npdaq *NordPoolDayAheadQuery) First(ctx context.Context) (*NordPoolDayAhead, error) {
	nodes, err := npdaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nordpooldayahead.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (npdaq *NordPoolDayAheadQuery) FirstX(ctx context.Context) *NordPoolDayAhead {
	node, err := npdaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NordPoolDayAhead ID from the query.
// Returns a *NotFoundError when no NordPoolDayAhead ID was found.
func (npdaq *NordPoolDayAheadQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = npdaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nordpooldayahead.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (npdaq *NordPoolDayAheadQuery) FirstIDX(ctx context.Context) int {
	id, err := npdaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NordPoolDayAhead entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NordPoolDayAhead entity is found.
// Returns a *NotFoundError when no NordPoolDayAhead entities are found.
func (npdaq *NordPoolDayAheadQuery) Only(ctx context.Context) (*NordPoolDayAhead, error) {
	nodes, err := npdaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nordpooldayahead.Label}
	default:
		return nil, &NotSingularError{nordpooldayahead.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (npdaq *NordPoolDayAheadQuery) OnlyX(ctx context.Context) *NordPoolDayAhead {
	node, err := npdaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NordPoolDayAhead ID in the query.
// Returns a *NotSingularError when more than one NordPoolDayAhead ID is found.
// Returns a *NotFoundError when no entities are found.
func (npdaq *NordPoolDayAheadQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = npdaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nordpooldayahead.Label}
	default:
		err = &NotSingularError{nordpooldayahead.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (npdaq *NordPoolDayAheadQuery) OnlyIDX(ctx context.Context) int {
	id, err := npdaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NordPoolDayAheads.
func (npdaq *NordPoolDayAheadQuery) All(ctx context.Context) ([]*NordPoolDayAhead, error) {
	if err := npdaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return npdaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (npdaq *NordPoolDayAheadQuery) AllX(ctx context.Context) []*NordPoolDayAhead {
	nodes, err := npdaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NordPoolDayAhead IDs.
func (npdaq *NordPoolDayAheadQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := npdaq.Select(nordpooldayahead.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (npdaq *NordPoolDayAheadQuery) IDsX(ctx context.Context) []int {
	ids, err := npdaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (npdaq *NordPoolDayAheadQuery) Count(ctx context.Context) (int, error) {
	if err := npdaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return npdaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (npdaq *NordPoolDayAheadQuery) CountX(ctx context.Context) int {
	count, err := npdaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (npdaq *NordPoolDayAheadQuery) Exist(ctx context.Context) (bool, error) {
	if err := npdaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return npdaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (npdaq *NordPoolDayAheadQuery) ExistX(ctx context.Context) bool {
	exist, err := npdaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NordPoolDayAheadQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (npdaq *NordPoolDayAheadQuery) Clone() *NordPoolDayAheadQuery {
	if npdaq == nil {
		return nil
	}
	return &NordPoolDayAheadQuery{
		config:     npdaq.config,
		limit:      npdaq.limit,
		offset:     npdaq.offset,
		order:      append([]OrderFunc{}, npdaq.order...),
		predicates: append([]predicate.NordPoolDayAhead{}, npdaq.predicates...),
		// clone intermediate query.
		sql:    npdaq.sql.Clone(),
		path:   npdaq.path,
		unique: npdaq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Date string `json:"date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NordPoolDayAhead.Query().
//		GroupBy(nordpooldayahead.FieldDate).
//		Aggregate(codegen.Count()).
//		Scan(ctx, &v)
//
func (npdaq *NordPoolDayAheadQuery) GroupBy(field string, fields ...string) *NordPoolDayAheadGroupBy {
	grbuild := &NordPoolDayAheadGroupBy{config: npdaq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := npdaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return npdaq.sqlQuery(ctx), nil
	}
	grbuild.label = nordpooldayahead.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Date string `json:"date,omitempty"`
//	}
//
//	client.NordPoolDayAhead.Query().
//		Select(nordpooldayahead.FieldDate).
//		Scan(ctx, &v)
//
func (npdaq *NordPoolDayAheadQuery) Select(fields ...string) *NordPoolDayAheadSelect {
	npdaq.fields = append(npdaq.fields, fields...)
	selbuild := &NordPoolDayAheadSelect{NordPoolDayAheadQuery: npdaq}
	selbuild.label = nordpooldayahead.Label
	selbuild.flds, selbuild.scan = &npdaq.fields, selbuild.Scan
	return selbuild
}

func (npdaq *NordPoolDayAheadQuery) prepareQuery(ctx context.Context) error {
	for _, f := range npdaq.fields {
		if !nordpooldayahead.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
		}
	}
	if npdaq.path != nil {
		prev, err := npdaq.path(ctx)
		if err != nil {
			return err
		}
		npdaq.sql = prev
	}
	return nil
}

func (npdaq *NordPoolDayAheadQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NordPoolDayAhead, error) {
	var (
		nodes = []*NordPoolDayAhead{}
		_spec = npdaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*NordPoolDayAhead).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &NordPoolDayAhead{config: npdaq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, npdaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (npdaq *NordPoolDayAheadQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := npdaq.querySpec()
	_spec.Node.Columns = npdaq.fields
	if len(npdaq.fields) > 0 {
		_spec.Unique = npdaq.unique != nil && *npdaq.unique
	}
	return sqlgraph.CountNodes(ctx, npdaq.driver, _spec)
}

func (npdaq *NordPoolDayAheadQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := npdaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("codegen: check existence: %w", err)
	}
	return n > 0, nil
}

func (npdaq *NordPoolDayAheadQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nordpooldayahead.Table,
			Columns: nordpooldayahead.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nordpooldayahead.FieldID,
			},
		},
		From:   npdaq.sql,
		Unique: true,
	}
	if unique := npdaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := npdaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nordpooldayahead.FieldID)
		for i := range fields {
			if fields[i] != nordpooldayahead.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := npdaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := npdaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := npdaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := npdaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (npdaq *NordPoolDayAheadQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(npdaq.driver.Dialect())
	t1 := builder.Table(nordpooldayahead.Table)
	columns := npdaq.fields
	if len(columns) == 0 {
		columns = nordpooldayahead.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if npdaq.sql != nil {
		selector = npdaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if npdaq.unique != nil && *npdaq.unique {
		selector.Distinct()
	}
	for _, p := range npdaq.predicates {
		p(selector)
	}
	for _, p := range npdaq.order {
		p(selector)
	}
	if offset := npdaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := npdaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NordPoolDayAheadGroupBy is the group-by builder for NordPoolDayAhead entities.
type NordPoolDayAheadGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (npdagb *NordPoolDayAheadGroupBy) Aggregate(fns ...AggregateFunc) *NordPoolDayAheadGroupBy {
	npdagb.fns = append(npdagb.fns, fns...)
	return npdagb
}

// Scan applies the group-by query and scans the result into the given value.
func (npdagb *NordPoolDayAheadGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := npdagb.path(ctx)
	if err != nil {
		return err
	}
	npdagb.sql = query
	return npdagb.sqlScan(ctx, v)
}

func (npdagb *NordPoolDayAheadGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range npdagb.fields {
		if !nordpooldayahead.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := npdagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := npdagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (npdagb *NordPoolDayAheadGroupBy) sqlQuery() *sql.Selector {
	selector := npdagb.sql.Select()
	aggregation := make([]string, 0, len(npdagb.fns))
	for _, fn := range npdagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(npdagb.fields)+len(npdagb.fns))
		for _, f := range npdagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(npdagb.fields...)...)
}

// NordPoolDayAheadSelect is the builder for selecting fields of NordPoolDayAhead entities.
type NordPoolDayAheadSelect struct {
	*NordPoolDayAheadQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (npdas *NordPoolDayAheadSelect) Scan(ctx context.Context, v interface{}) error {
	if err := npdas.prepareQuery(ctx); err != nil {
		return err
	}
	npdas.sql = npdas.NordPoolDayAheadQuery.sqlQuery(ctx)
	return npdas.sqlScan(ctx, v)
}

func (npdas *NordPoolDayAheadSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := npdas.sql.Query()
	if err := npdas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
