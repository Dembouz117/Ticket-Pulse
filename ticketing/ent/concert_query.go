// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"ticketing/ent/concert"
	"ticketing/ent/concertsession"
	"ticketing/ent/genre"
	"ticketing/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ConcertQuery is the builder for querying Concert entities.
type ConcertQuery struct {
	config
	ctx                    *QueryContext
	order                  []concert.OrderOption
	inters                 []Interceptor
	predicates             []predicate.Concert
	withHasConcertSessions *ConcertSessionQuery
	withGenres             *GenreQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConcertQuery builder.
func (cq *ConcertQuery) Where(ps ...predicate.Concert) *ConcertQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ConcertQuery) Limit(limit int) *ConcertQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ConcertQuery) Offset(offset int) *ConcertQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ConcertQuery) Unique(unique bool) *ConcertQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ConcertQuery) Order(o ...concert.OrderOption) *ConcertQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryHasConcertSessions chains the current query on the "hasConcertSessions" edge.
func (cq *ConcertQuery) QueryHasConcertSessions() *ConcertSessionQuery {
	query := (&ConcertSessionClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(concert.Table, concert.FieldID, selector),
			sqlgraph.To(concertsession.Table, concertsession.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, concert.HasConcertSessionsTable, concert.HasConcertSessionsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGenres chains the current query on the "genres" edge.
func (cq *ConcertQuery) QueryGenres() *GenreQuery {
	query := (&GenreClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(concert.Table, concert.FieldID, selector),
			sqlgraph.To(genre.Table, genre.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, concert.GenresTable, concert.GenresPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Concert entity from the query.
// Returns a *NotFoundError when no Concert was found.
func (cq *ConcertQuery) First(ctx context.Context) (*Concert, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{concert.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ConcertQuery) FirstX(ctx context.Context) *Concert {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Concert ID from the query.
// Returns a *NotFoundError when no Concert ID was found.
func (cq *ConcertQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{concert.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ConcertQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Concert entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Concert entity is found.
// Returns a *NotFoundError when no Concert entities are found.
func (cq *ConcertQuery) Only(ctx context.Context) (*Concert, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{concert.Label}
	default:
		return nil, &NotSingularError{concert.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ConcertQuery) OnlyX(ctx context.Context) *Concert {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Concert ID in the query.
// Returns a *NotSingularError when more than one Concert ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ConcertQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{concert.Label}
	default:
		err = &NotSingularError{concert.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ConcertQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Concerts.
func (cq *ConcertQuery) All(ctx context.Context) ([]*Concert, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Concert, *ConcertQuery]()
	return withInterceptors[[]*Concert](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ConcertQuery) AllX(ctx context.Context) []*Concert {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Concert IDs.
func (cq *ConcertQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(concert.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ConcertQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ConcertQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ConcertQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ConcertQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ConcertQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ConcertQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConcertQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ConcertQuery) Clone() *ConcertQuery {
	if cq == nil {
		return nil
	}
	return &ConcertQuery{
		config:                 cq.config,
		ctx:                    cq.ctx.Clone(),
		order:                  append([]concert.OrderOption{}, cq.order...),
		inters:                 append([]Interceptor{}, cq.inters...),
		predicates:             append([]predicate.Concert{}, cq.predicates...),
		withHasConcertSessions: cq.withHasConcertSessions.Clone(),
		withGenres:             cq.withGenres.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithHasConcertSessions tells the query-builder to eager-load the nodes that are connected to
// the "hasConcertSessions" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConcertQuery) WithHasConcertSessions(opts ...func(*ConcertSessionQuery)) *ConcertQuery {
	query := (&ConcertSessionClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withHasConcertSessions = query
	return cq
}

// WithGenres tells the query-builder to eager-load the nodes that are connected to
// the "genres" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConcertQuery) WithGenres(opts ...func(*GenreQuery)) *ConcertQuery {
	query := (&GenreClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withGenres = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Concert.Query().
//		GroupBy(concert.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ConcertQuery) GroupBy(field string, fields ...string) *ConcertGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ConcertGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = concert.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Concert.Query().
//		Select(concert.FieldTitle).
//		Scan(ctx, &v)
func (cq *ConcertQuery) Select(fields ...string) *ConcertSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ConcertSelect{ConcertQuery: cq}
	sbuild.label = concert.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ConcertSelect configured with the given aggregations.
func (cq *ConcertQuery) Aggregate(fns ...AggregateFunc) *ConcertSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ConcertQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !concert.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ConcertQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Concert, error) {
	var (
		nodes       = []*Concert{}
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withHasConcertSessions != nil,
			cq.withGenres != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Concert).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Concert{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withHasConcertSessions; query != nil {
		if err := cq.loadHasConcertSessions(ctx, query, nodes,
			func(n *Concert) { n.Edges.HasConcertSessions = []*ConcertSession{} },
			func(n *Concert, e *ConcertSession) {
				n.Edges.HasConcertSessions = append(n.Edges.HasConcertSessions, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := cq.withGenres; query != nil {
		if err := cq.loadGenres(ctx, query, nodes,
			func(n *Concert) { n.Edges.Genres = []*Genre{} },
			func(n *Concert, e *Genre) { n.Edges.Genres = append(n.Edges.Genres, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ConcertQuery) loadHasConcertSessions(ctx context.Context, query *ConcertSessionQuery, nodes []*Concert, init func(*Concert), assign func(*Concert, *ConcertSession)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Concert)
	nids := make(map[uuid.UUID]map[*Concert]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(concert.HasConcertSessionsTable)
		s.Join(joinT).On(s.C(concertsession.FieldID), joinT.C(concert.HasConcertSessionsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(concert.HasConcertSessionsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(concert.HasConcertSessionsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Concert]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*ConcertSession](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "hasConcertSessions" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *ConcertQuery) loadGenres(ctx context.Context, query *GenreQuery, nodes []*Concert, init func(*Concert), assign func(*Concert, *Genre)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Concert)
	nids := make(map[uuid.UUID]map[*Concert]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(concert.GenresTable)
		s.Join(joinT).On(s.C(genre.FieldID), joinT.C(concert.GenresPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(concert.GenresPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(concert.GenresPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Concert]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Genre](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "genres" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (cq *ConcertQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ConcertQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(concert.Table, concert.Columns, sqlgraph.NewFieldSpec(concert.FieldID, field.TypeUUID))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, concert.FieldID)
		for i := range fields {
			if fields[i] != concert.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ConcertQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(concert.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = concert.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ConcertGroupBy is the group-by builder for Concert entities.
type ConcertGroupBy struct {
	selector
	build *ConcertQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ConcertGroupBy) Aggregate(fns ...AggregateFunc) *ConcertGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ConcertGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConcertQuery, *ConcertGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ConcertGroupBy) sqlScan(ctx context.Context, root *ConcertQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ConcertSelect is the builder for selecting fields of Concert entities.
type ConcertSelect struct {
	*ConcertQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ConcertSelect) Aggregate(fns ...AggregateFunc) *ConcertSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ConcertSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConcertQuery, *ConcertSelect](ctx, cs.ConcertQuery, cs, cs.inters, v)
}

func (cs *ConcertSelect) sqlScan(ctx context.Context, root *ConcertQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
