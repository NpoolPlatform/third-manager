// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/appsmstemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppSMSTemplateQuery is the builder for querying AppSMSTemplate entities.
type AppSMSTemplateQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppSMSTemplate
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppSMSTemplateQuery builder.
func (astq *AppSMSTemplateQuery) Where(ps ...predicate.AppSMSTemplate) *AppSMSTemplateQuery {
	astq.predicates = append(astq.predicates, ps...)
	return astq
}

// Limit adds a limit step to the query.
func (astq *AppSMSTemplateQuery) Limit(limit int) *AppSMSTemplateQuery {
	astq.limit = &limit
	return astq
}

// Offset adds an offset step to the query.
func (astq *AppSMSTemplateQuery) Offset(offset int) *AppSMSTemplateQuery {
	astq.offset = &offset
	return astq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (astq *AppSMSTemplateQuery) Unique(unique bool) *AppSMSTemplateQuery {
	astq.unique = &unique
	return astq
}

// Order adds an order step to the query.
func (astq *AppSMSTemplateQuery) Order(o ...OrderFunc) *AppSMSTemplateQuery {
	astq.order = append(astq.order, o...)
	return astq
}

// First returns the first AppSMSTemplate entity from the query.
// Returns a *NotFoundError when no AppSMSTemplate was found.
func (astq *AppSMSTemplateQuery) First(ctx context.Context) (*AppSMSTemplate, error) {
	nodes, err := astq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appsmstemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (astq *AppSMSTemplateQuery) FirstX(ctx context.Context) *AppSMSTemplate {
	node, err := astq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppSMSTemplate ID from the query.
// Returns a *NotFoundError when no AppSMSTemplate ID was found.
func (astq *AppSMSTemplateQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = astq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appsmstemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (astq *AppSMSTemplateQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := astq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppSMSTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppSMSTemplate entity is found.
// Returns a *NotFoundError when no AppSMSTemplate entities are found.
func (astq *AppSMSTemplateQuery) Only(ctx context.Context) (*AppSMSTemplate, error) {
	nodes, err := astq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appsmstemplate.Label}
	default:
		return nil, &NotSingularError{appsmstemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (astq *AppSMSTemplateQuery) OnlyX(ctx context.Context) *AppSMSTemplate {
	node, err := astq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppSMSTemplate ID in the query.
// Returns a *NotSingularError when more than one AppSMSTemplate ID is found.
// Returns a *NotFoundError when no entities are found.
func (astq *AppSMSTemplateQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = astq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appsmstemplate.Label}
	default:
		err = &NotSingularError{appsmstemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (astq *AppSMSTemplateQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := astq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppSMSTemplates.
func (astq *AppSMSTemplateQuery) All(ctx context.Context) ([]*AppSMSTemplate, error) {
	if err := astq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return astq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (astq *AppSMSTemplateQuery) AllX(ctx context.Context) []*AppSMSTemplate {
	nodes, err := astq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppSMSTemplate IDs.
func (astq *AppSMSTemplateQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := astq.Select(appsmstemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (astq *AppSMSTemplateQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := astq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (astq *AppSMSTemplateQuery) Count(ctx context.Context) (int, error) {
	if err := astq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return astq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (astq *AppSMSTemplateQuery) CountX(ctx context.Context) int {
	count, err := astq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (astq *AppSMSTemplateQuery) Exist(ctx context.Context) (bool, error) {
	if err := astq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return astq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (astq *AppSMSTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := astq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppSMSTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (astq *AppSMSTemplateQuery) Clone() *AppSMSTemplateQuery {
	if astq == nil {
		return nil
	}
	return &AppSMSTemplateQuery{
		config:     astq.config,
		limit:      astq.limit,
		offset:     astq.offset,
		order:      append([]OrderFunc{}, astq.order...),
		predicates: append([]predicate.AppSMSTemplate{}, astq.predicates...),
		// clone intermediate query.
		sql:    astq.sql.Clone(),
		path:   astq.path,
		unique: astq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppSMSTemplate.Query().
//		GroupBy(appsmstemplate.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (astq *AppSMSTemplateQuery) GroupBy(field string, fields ...string) *AppSMSTemplateGroupBy {
	grbuild := &AppSMSTemplateGroupBy{config: astq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := astq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return astq.sqlQuery(ctx), nil
	}
	grbuild.label = appsmstemplate.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.AppSMSTemplate.Query().
//		Select(appsmstemplate.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (astq *AppSMSTemplateQuery) Select(fields ...string) *AppSMSTemplateSelect {
	astq.fields = append(astq.fields, fields...)
	selbuild := &AppSMSTemplateSelect{AppSMSTemplateQuery: astq}
	selbuild.label = appsmstemplate.Label
	selbuild.flds, selbuild.scan = &astq.fields, selbuild.Scan
	return selbuild
}

func (astq *AppSMSTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, f := range astq.fields {
		if !appsmstemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if astq.path != nil {
		prev, err := astq.path(ctx)
		if err != nil {
			return err
		}
		astq.sql = prev
	}
	if appsmstemplate.Policy == nil {
		return errors.New("ent: uninitialized appsmstemplate.Policy (forgotten import ent/runtime?)")
	}
	if err := appsmstemplate.Policy.EvalQuery(ctx, astq); err != nil {
		return err
	}
	return nil
}

func (astq *AppSMSTemplateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppSMSTemplate, error) {
	var (
		nodes = []*AppSMSTemplate{}
		_spec = astq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppSMSTemplate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppSMSTemplate{config: astq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(astq.modifiers) > 0 {
		_spec.Modifiers = astq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, astq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (astq *AppSMSTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := astq.querySpec()
	if len(astq.modifiers) > 0 {
		_spec.Modifiers = astq.modifiers
	}
	_spec.Node.Columns = astq.fields
	if len(astq.fields) > 0 {
		_spec.Unique = astq.unique != nil && *astq.unique
	}
	return sqlgraph.CountNodes(ctx, astq.driver, _spec)
}

func (astq *AppSMSTemplateQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := astq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (astq *AppSMSTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appsmstemplate.Table,
			Columns: appsmstemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appsmstemplate.FieldID,
			},
		},
		From:   astq.sql,
		Unique: true,
	}
	if unique := astq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := astq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appsmstemplate.FieldID)
		for i := range fields {
			if fields[i] != appsmstemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := astq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := astq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := astq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := astq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (astq *AppSMSTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(astq.driver.Dialect())
	t1 := builder.Table(appsmstemplate.Table)
	columns := astq.fields
	if len(columns) == 0 {
		columns = appsmstemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if astq.sql != nil {
		selector = astq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if astq.unique != nil && *astq.unique {
		selector.Distinct()
	}
	for _, m := range astq.modifiers {
		m(selector)
	}
	for _, p := range astq.predicates {
		p(selector)
	}
	for _, p := range astq.order {
		p(selector)
	}
	if offset := astq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := astq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (astq *AppSMSTemplateQuery) ForUpdate(opts ...sql.LockOption) *AppSMSTemplateQuery {
	if astq.driver.Dialect() == dialect.Postgres {
		astq.Unique(false)
	}
	astq.modifiers = append(astq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return astq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (astq *AppSMSTemplateQuery) ForShare(opts ...sql.LockOption) *AppSMSTemplateQuery {
	if astq.driver.Dialect() == dialect.Postgres {
		astq.Unique(false)
	}
	astq.modifiers = append(astq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return astq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (astq *AppSMSTemplateQuery) Modify(modifiers ...func(s *sql.Selector)) *AppSMSTemplateSelect {
	astq.modifiers = append(astq.modifiers, modifiers...)
	return astq.Select()
}

// AppSMSTemplateGroupBy is the group-by builder for AppSMSTemplate entities.
type AppSMSTemplateGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (astgb *AppSMSTemplateGroupBy) Aggregate(fns ...AggregateFunc) *AppSMSTemplateGroupBy {
	astgb.fns = append(astgb.fns, fns...)
	return astgb
}

// Scan applies the group-by query and scans the result into the given value.
func (astgb *AppSMSTemplateGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := astgb.path(ctx)
	if err != nil {
		return err
	}
	astgb.sql = query
	return astgb.sqlScan(ctx, v)
}

func (astgb *AppSMSTemplateGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range astgb.fields {
		if !appsmstemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := astgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := astgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (astgb *AppSMSTemplateGroupBy) sqlQuery() *sql.Selector {
	selector := astgb.sql.Select()
	aggregation := make([]string, 0, len(astgb.fns))
	for _, fn := range astgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(astgb.fields)+len(astgb.fns))
		for _, f := range astgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(astgb.fields...)...)
}

// AppSMSTemplateSelect is the builder for selecting fields of AppSMSTemplate entities.
type AppSMSTemplateSelect struct {
	*AppSMSTemplateQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (asts *AppSMSTemplateSelect) Scan(ctx context.Context, v interface{}) error {
	if err := asts.prepareQuery(ctx); err != nil {
		return err
	}
	asts.sql = asts.AppSMSTemplateQuery.sqlQuery(ctx)
	return asts.sqlScan(ctx, v)
}

func (asts *AppSMSTemplateSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := asts.sql.Query()
	if err := asts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (asts *AppSMSTemplateSelect) Modify(modifiers ...func(s *sql.Selector)) *AppSMSTemplateSelect {
	asts.modifiers = append(asts.modifiers, modifiers...)
	return asts
}
