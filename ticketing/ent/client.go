// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"ticketing/ent/migrate"

	"ticketing/ent/concert"
	"ticketing/ent/concertsession"
	"ticketing/ent/genre"
	"ticketing/ent/section"
	"ticketing/ent/ticket"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Concert is the client for interacting with the Concert builders.
	Concert *ConcertClient
	// ConcertSession is the client for interacting with the ConcertSession builders.
	ConcertSession *ConcertSessionClient
	// Genre is the client for interacting with the Genre builders.
	Genre *GenreClient
	// Section is the client for interacting with the Section builders.
	Section *SectionClient
	// Ticket is the client for interacting with the Ticket builders.
	Ticket *TicketClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Concert = NewConcertClient(c.config)
	c.ConcertSession = NewConcertSessionClient(c.config)
	c.Genre = NewGenreClient(c.config)
	c.Section = NewSectionClient(c.config)
	c.Ticket = NewTicketClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Concert:        NewConcertClient(cfg),
		ConcertSession: NewConcertSessionClient(cfg),
		Genre:          NewGenreClient(cfg),
		Section:        NewSectionClient(cfg),
		Ticket:         NewTicketClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Concert:        NewConcertClient(cfg),
		ConcertSession: NewConcertSessionClient(cfg),
		Genre:          NewGenreClient(cfg),
		Section:        NewSectionClient(cfg),
		Ticket:         NewTicketClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Concert.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Concert.Use(hooks...)
	c.ConcertSession.Use(hooks...)
	c.Genre.Use(hooks...)
	c.Section.Use(hooks...)
	c.Ticket.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Concert.Intercept(interceptors...)
	c.ConcertSession.Intercept(interceptors...)
	c.Genre.Intercept(interceptors...)
	c.Section.Intercept(interceptors...)
	c.Ticket.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ConcertMutation:
		return c.Concert.mutate(ctx, m)
	case *ConcertSessionMutation:
		return c.ConcertSession.mutate(ctx, m)
	case *GenreMutation:
		return c.Genre.mutate(ctx, m)
	case *SectionMutation:
		return c.Section.mutate(ctx, m)
	case *TicketMutation:
		return c.Ticket.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ConcertClient is a client for the Concert schema.
type ConcertClient struct {
	config
}

// NewConcertClient returns a client for the Concert from the given config.
func NewConcertClient(c config) *ConcertClient {
	return &ConcertClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `concert.Hooks(f(g(h())))`.
func (c *ConcertClient) Use(hooks ...Hook) {
	c.hooks.Concert = append(c.hooks.Concert, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `concert.Intercept(f(g(h())))`.
func (c *ConcertClient) Intercept(interceptors ...Interceptor) {
	c.inters.Concert = append(c.inters.Concert, interceptors...)
}

// Create returns a builder for creating a Concert entity.
func (c *ConcertClient) Create() *ConcertCreate {
	mutation := newConcertMutation(c.config, OpCreate)
	return &ConcertCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Concert entities.
func (c *ConcertClient) CreateBulk(builders ...*ConcertCreate) *ConcertCreateBulk {
	return &ConcertCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ConcertClient) MapCreateBulk(slice any, setFunc func(*ConcertCreate, int)) *ConcertCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ConcertCreateBulk{err: fmt.Errorf("calling to ConcertClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ConcertCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ConcertCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Concert.
func (c *ConcertClient) Update() *ConcertUpdate {
	mutation := newConcertMutation(c.config, OpUpdate)
	return &ConcertUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ConcertClient) UpdateOne(co *Concert) *ConcertUpdateOne {
	mutation := newConcertMutation(c.config, OpUpdateOne, withConcert(co))
	return &ConcertUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ConcertClient) UpdateOneID(id uuid.UUID) *ConcertUpdateOne {
	mutation := newConcertMutation(c.config, OpUpdateOne, withConcertID(id))
	return &ConcertUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Concert.
func (c *ConcertClient) Delete() *ConcertDelete {
	mutation := newConcertMutation(c.config, OpDelete)
	return &ConcertDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ConcertClient) DeleteOne(co *Concert) *ConcertDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ConcertClient) DeleteOneID(id uuid.UUID) *ConcertDeleteOne {
	builder := c.Delete().Where(concert.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ConcertDeleteOne{builder}
}

// Query returns a query builder for Concert.
func (c *ConcertClient) Query() *ConcertQuery {
	return &ConcertQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeConcert},
		inters: c.Interceptors(),
	}
}

// Get returns a Concert entity by its id.
func (c *ConcertClient) Get(ctx context.Context, id uuid.UUID) (*Concert, error) {
	return c.Query().Where(concert.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ConcertClient) GetX(ctx context.Context, id uuid.UUID) *Concert {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHasConcertSessions queries the hasConcertSessions edge of a Concert.
func (c *ConcertClient) QueryHasConcertSessions(co *Concert) *ConcertSessionQuery {
	query := (&ConcertSessionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(concert.Table, concert.FieldID, id),
			sqlgraph.To(concertsession.Table, concertsession.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, concert.HasConcertSessionsTable, concert.HasConcertSessionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGenres queries the genres edge of a Concert.
func (c *ConcertClient) QueryGenres(co *Concert) *GenreQuery {
	query := (&GenreClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(concert.Table, concert.FieldID, id),
			sqlgraph.To(genre.Table, genre.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, concert.GenresTable, concert.GenresPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ConcertClient) Hooks() []Hook {
	return c.hooks.Concert
}

// Interceptors returns the client interceptors.
func (c *ConcertClient) Interceptors() []Interceptor {
	return c.inters.Concert
}

func (c *ConcertClient) mutate(ctx context.Context, m *ConcertMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ConcertCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ConcertUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ConcertUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ConcertDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Concert mutation op: %q", m.Op())
	}
}

// ConcertSessionClient is a client for the ConcertSession schema.
type ConcertSessionClient struct {
	config
}

// NewConcertSessionClient returns a client for the ConcertSession from the given config.
func NewConcertSessionClient(c config) *ConcertSessionClient {
	return &ConcertSessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `concertsession.Hooks(f(g(h())))`.
func (c *ConcertSessionClient) Use(hooks ...Hook) {
	c.hooks.ConcertSession = append(c.hooks.ConcertSession, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `concertsession.Intercept(f(g(h())))`.
func (c *ConcertSessionClient) Intercept(interceptors ...Interceptor) {
	c.inters.ConcertSession = append(c.inters.ConcertSession, interceptors...)
}

// Create returns a builder for creating a ConcertSession entity.
func (c *ConcertSessionClient) Create() *ConcertSessionCreate {
	mutation := newConcertSessionMutation(c.config, OpCreate)
	return &ConcertSessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ConcertSession entities.
func (c *ConcertSessionClient) CreateBulk(builders ...*ConcertSessionCreate) *ConcertSessionCreateBulk {
	return &ConcertSessionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ConcertSessionClient) MapCreateBulk(slice any, setFunc func(*ConcertSessionCreate, int)) *ConcertSessionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ConcertSessionCreateBulk{err: fmt.Errorf("calling to ConcertSessionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ConcertSessionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ConcertSessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ConcertSession.
func (c *ConcertSessionClient) Update() *ConcertSessionUpdate {
	mutation := newConcertSessionMutation(c.config, OpUpdate)
	return &ConcertSessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ConcertSessionClient) UpdateOne(cs *ConcertSession) *ConcertSessionUpdateOne {
	mutation := newConcertSessionMutation(c.config, OpUpdateOne, withConcertSession(cs))
	return &ConcertSessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ConcertSessionClient) UpdateOneID(id uuid.UUID) *ConcertSessionUpdateOne {
	mutation := newConcertSessionMutation(c.config, OpUpdateOne, withConcertSessionID(id))
	return &ConcertSessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ConcertSession.
func (c *ConcertSessionClient) Delete() *ConcertSessionDelete {
	mutation := newConcertSessionMutation(c.config, OpDelete)
	return &ConcertSessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ConcertSessionClient) DeleteOne(cs *ConcertSession) *ConcertSessionDeleteOne {
	return c.DeleteOneID(cs.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ConcertSessionClient) DeleteOneID(id uuid.UUID) *ConcertSessionDeleteOne {
	builder := c.Delete().Where(concertsession.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ConcertSessionDeleteOne{builder}
}

// Query returns a query builder for ConcertSession.
func (c *ConcertSessionClient) Query() *ConcertSessionQuery {
	return &ConcertSessionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeConcertSession},
		inters: c.Interceptors(),
	}
}

// Get returns a ConcertSession entity by its id.
func (c *ConcertSessionClient) Get(ctx context.Context, id uuid.UUID) (*ConcertSession, error) {
	return c.Query().Where(concertsession.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ConcertSessionClient) GetX(ctx context.Context, id uuid.UUID) *ConcertSession {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOfConcert queries the ofConcert edge of a ConcertSession.
func (c *ConcertSessionClient) QueryOfConcert(cs *ConcertSession) *ConcertQuery {
	query := (&ConcertClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cs.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(concertsession.Table, concertsession.FieldID, id),
			sqlgraph.To(concert.Table, concert.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, concertsession.OfConcertTable, concertsession.OfConcertPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cs.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHasSections queries the hasSections edge of a ConcertSession.
func (c *ConcertSessionClient) QueryHasSections(cs *ConcertSession) *SectionQuery {
	query := (&SectionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cs.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(concertsession.Table, concertsession.FieldID, id),
			sqlgraph.To(section.Table, section.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, concertsession.HasSectionsTable, concertsession.HasSectionsColumn),
		)
		fromV = sqlgraph.Neighbors(cs.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ConcertSessionClient) Hooks() []Hook {
	return c.hooks.ConcertSession
}

// Interceptors returns the client interceptors.
func (c *ConcertSessionClient) Interceptors() []Interceptor {
	return c.inters.ConcertSession
}

func (c *ConcertSessionClient) mutate(ctx context.Context, m *ConcertSessionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ConcertSessionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ConcertSessionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ConcertSessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ConcertSessionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ConcertSession mutation op: %q", m.Op())
	}
}

// GenreClient is a client for the Genre schema.
type GenreClient struct {
	config
}

// NewGenreClient returns a client for the Genre from the given config.
func NewGenreClient(c config) *GenreClient {
	return &GenreClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `genre.Hooks(f(g(h())))`.
func (c *GenreClient) Use(hooks ...Hook) {
	c.hooks.Genre = append(c.hooks.Genre, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `genre.Intercept(f(g(h())))`.
func (c *GenreClient) Intercept(interceptors ...Interceptor) {
	c.inters.Genre = append(c.inters.Genre, interceptors...)
}

// Create returns a builder for creating a Genre entity.
func (c *GenreClient) Create() *GenreCreate {
	mutation := newGenreMutation(c.config, OpCreate)
	return &GenreCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Genre entities.
func (c *GenreClient) CreateBulk(builders ...*GenreCreate) *GenreCreateBulk {
	return &GenreCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *GenreClient) MapCreateBulk(slice any, setFunc func(*GenreCreate, int)) *GenreCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &GenreCreateBulk{err: fmt.Errorf("calling to GenreClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*GenreCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &GenreCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Genre.
func (c *GenreClient) Update() *GenreUpdate {
	mutation := newGenreMutation(c.config, OpUpdate)
	return &GenreUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GenreClient) UpdateOne(ge *Genre) *GenreUpdateOne {
	mutation := newGenreMutation(c.config, OpUpdateOne, withGenre(ge))
	return &GenreUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GenreClient) UpdateOneID(id uuid.UUID) *GenreUpdateOne {
	mutation := newGenreMutation(c.config, OpUpdateOne, withGenreID(id))
	return &GenreUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Genre.
func (c *GenreClient) Delete() *GenreDelete {
	mutation := newGenreMutation(c.config, OpDelete)
	return &GenreDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GenreClient) DeleteOne(ge *Genre) *GenreDeleteOne {
	return c.DeleteOneID(ge.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GenreClient) DeleteOneID(id uuid.UUID) *GenreDeleteOne {
	builder := c.Delete().Where(genre.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GenreDeleteOne{builder}
}

// Query returns a query builder for Genre.
func (c *GenreClient) Query() *GenreQuery {
	return &GenreQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGenre},
		inters: c.Interceptors(),
	}
}

// Get returns a Genre entity by its id.
func (c *GenreClient) Get(ctx context.Context, id uuid.UUID) (*Genre, error) {
	return c.Query().Where(genre.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GenreClient) GetX(ctx context.Context, id uuid.UUID) *Genre {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryConcerts queries the concerts edge of a Genre.
func (c *GenreClient) QueryConcerts(ge *Genre) *ConcertQuery {
	query := (&ConcertClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ge.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(genre.Table, genre.FieldID, id),
			sqlgraph.To(concert.Table, concert.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, genre.ConcertsTable, genre.ConcertsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ge.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GenreClient) Hooks() []Hook {
	return c.hooks.Genre
}

// Interceptors returns the client interceptors.
func (c *GenreClient) Interceptors() []Interceptor {
	return c.inters.Genre
}

func (c *GenreClient) mutate(ctx context.Context, m *GenreMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GenreCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GenreUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GenreUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GenreDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Genre mutation op: %q", m.Op())
	}
}

// SectionClient is a client for the Section schema.
type SectionClient struct {
	config
}

// NewSectionClient returns a client for the Section from the given config.
func NewSectionClient(c config) *SectionClient {
	return &SectionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `section.Hooks(f(g(h())))`.
func (c *SectionClient) Use(hooks ...Hook) {
	c.hooks.Section = append(c.hooks.Section, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `section.Intercept(f(g(h())))`.
func (c *SectionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Section = append(c.inters.Section, interceptors...)
}

// Create returns a builder for creating a Section entity.
func (c *SectionClient) Create() *SectionCreate {
	mutation := newSectionMutation(c.config, OpCreate)
	return &SectionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Section entities.
func (c *SectionClient) CreateBulk(builders ...*SectionCreate) *SectionCreateBulk {
	return &SectionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SectionClient) MapCreateBulk(slice any, setFunc func(*SectionCreate, int)) *SectionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SectionCreateBulk{err: fmt.Errorf("calling to SectionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SectionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SectionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Section.
func (c *SectionClient) Update() *SectionUpdate {
	mutation := newSectionMutation(c.config, OpUpdate)
	return &SectionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SectionClient) UpdateOne(s *Section) *SectionUpdateOne {
	mutation := newSectionMutation(c.config, OpUpdateOne, withSection(s))
	return &SectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SectionClient) UpdateOneID(id uuid.UUID) *SectionUpdateOne {
	mutation := newSectionMutation(c.config, OpUpdateOne, withSectionID(id))
	return &SectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Section.
func (c *SectionClient) Delete() *SectionDelete {
	mutation := newSectionMutation(c.config, OpDelete)
	return &SectionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SectionClient) DeleteOne(s *Section) *SectionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SectionClient) DeleteOneID(id uuid.UUID) *SectionDeleteOne {
	builder := c.Delete().Where(section.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SectionDeleteOne{builder}
}

// Query returns a query builder for Section.
func (c *SectionClient) Query() *SectionQuery {
	return &SectionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSection},
		inters: c.Interceptors(),
	}
}

// Get returns a Section entity by its id.
func (c *SectionClient) Get(ctx context.Context, id uuid.UUID) (*Section, error) {
	return c.Query().Where(section.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SectionClient) GetX(ctx context.Context, id uuid.UUID) *Section {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHasTickets queries the hasTickets edge of a Section.
func (c *SectionClient) QueryHasTickets(s *Section) *TicketQuery {
	query := (&TicketClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(section.Table, section.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, section.HasTicketsTable, section.HasTicketsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAtConcertSession queries the atConcertSession edge of a Section.
func (c *SectionClient) QueryAtConcertSession(s *Section) *ConcertSessionQuery {
	query := (&ConcertSessionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(section.Table, section.FieldID, id),
			sqlgraph.To(concertsession.Table, concertsession.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, section.AtConcertSessionTable, section.AtConcertSessionColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SectionClient) Hooks() []Hook {
	return c.hooks.Section
}

// Interceptors returns the client interceptors.
func (c *SectionClient) Interceptors() []Interceptor {
	return c.inters.Section
}

func (c *SectionClient) mutate(ctx context.Context, m *SectionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SectionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SectionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SectionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Section mutation op: %q", m.Op())
	}
}

// TicketClient is a client for the Ticket schema.
type TicketClient struct {
	config
}

// NewTicketClient returns a client for the Ticket from the given config.
func NewTicketClient(c config) *TicketClient {
	return &TicketClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ticket.Hooks(f(g(h())))`.
func (c *TicketClient) Use(hooks ...Hook) {
	c.hooks.Ticket = append(c.hooks.Ticket, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `ticket.Intercept(f(g(h())))`.
func (c *TicketClient) Intercept(interceptors ...Interceptor) {
	c.inters.Ticket = append(c.inters.Ticket, interceptors...)
}

// Create returns a builder for creating a Ticket entity.
func (c *TicketClient) Create() *TicketCreate {
	mutation := newTicketMutation(c.config, OpCreate)
	return &TicketCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ticket entities.
func (c *TicketClient) CreateBulk(builders ...*TicketCreate) *TicketCreateBulk {
	return &TicketCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TicketClient) MapCreateBulk(slice any, setFunc func(*TicketCreate, int)) *TicketCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TicketCreateBulk{err: fmt.Errorf("calling to TicketClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TicketCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TicketCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ticket.
func (c *TicketClient) Update() *TicketUpdate {
	mutation := newTicketMutation(c.config, OpUpdate)
	return &TicketUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TicketClient) UpdateOne(t *Ticket) *TicketUpdateOne {
	mutation := newTicketMutation(c.config, OpUpdateOne, withTicket(t))
	return &TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TicketClient) UpdateOneID(id uuid.UUID) *TicketUpdateOne {
	mutation := newTicketMutation(c.config, OpUpdateOne, withTicketID(id))
	return &TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ticket.
func (c *TicketClient) Delete() *TicketDelete {
	mutation := newTicketMutation(c.config, OpDelete)
	return &TicketDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TicketClient) DeleteOne(t *Ticket) *TicketDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TicketClient) DeleteOneID(id uuid.UUID) *TicketDeleteOne {
	builder := c.Delete().Where(ticket.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TicketDeleteOne{builder}
}

// Query returns a query builder for Ticket.
func (c *TicketClient) Query() *TicketQuery {
	return &TicketQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTicket},
		inters: c.Interceptors(),
	}
}

// Get returns a Ticket entity by its id.
func (c *TicketClient) Get(ctx context.Context, id uuid.UUID) (*Ticket, error) {
	return c.Query().Where(ticket.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TicketClient) GetX(ctx context.Context, id uuid.UUID) *Ticket {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWithinSection queries the withinSection edge of a Ticket.
func (c *TicketClient) QueryWithinSection(t *Ticket) *SectionQuery {
	query := (&SectionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(section.Table, section.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticket.WithinSectionTable, ticket.WithinSectionColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TicketClient) Hooks() []Hook {
	return c.hooks.Ticket
}

// Interceptors returns the client interceptors.
func (c *TicketClient) Interceptors() []Interceptor {
	return c.inters.Ticket
}

func (c *TicketClient) mutate(ctx context.Context, m *TicketMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TicketCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TicketUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TicketDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Ticket mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Concert, ConcertSession, Genre, Section, Ticket []ent.Hook
	}
	inters struct {
		Concert, ConcertSession, Genre, Section, Ticket []ent.Interceptor
	}
)
