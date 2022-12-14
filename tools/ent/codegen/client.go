// Code generated by entc, DO NOT EDIT.

package codegen

import (
	"context"
	"fmt"
	"log"

	"github.com/alexgtn/go-linkshort/tools/ent/codegen/migrate"

	"github.com/alexgtn/go-linkshort/tools/ent/codegen/nordpooldayahead"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// NordPoolDayAhead is the client for interacting with the NordPoolDayAhead builders.
	NordPoolDayAhead *NordPoolDayAheadClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.NordPoolDayAhead = NewNordPoolDayAheadClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("codegen: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("codegen: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		NordPoolDayAhead: NewNordPoolDayAheadClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:              ctx,
		config:           cfg,
		NordPoolDayAhead: NewNordPoolDayAheadClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		NordPoolDayAhead.
//		Query().
//		Count(ctx)
//
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
	c.NordPoolDayAhead.Use(hooks...)
}

// NordPoolDayAheadClient is a client for the NordPoolDayAhead schema.
type NordPoolDayAheadClient struct {
	config
}

// NewNordPoolDayAheadClient returns a client for the NordPoolDayAhead from the given config.
func NewNordPoolDayAheadClient(c config) *NordPoolDayAheadClient {
	return &NordPoolDayAheadClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `nordpooldayahead.Hooks(f(g(h())))`.
func (c *NordPoolDayAheadClient) Use(hooks ...Hook) {
	c.hooks.NordPoolDayAhead = append(c.hooks.NordPoolDayAhead, hooks...)
}

// Create returns a create builder for NordPoolDayAhead.
func (c *NordPoolDayAheadClient) Create() *NordPoolDayAheadCreate {
	mutation := newNordPoolDayAheadMutation(c.config, OpCreate)
	return &NordPoolDayAheadCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of NordPoolDayAhead entities.
func (c *NordPoolDayAheadClient) CreateBulk(builders ...*NordPoolDayAheadCreate) *NordPoolDayAheadCreateBulk {
	return &NordPoolDayAheadCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for NordPoolDayAhead.
func (c *NordPoolDayAheadClient) Update() *NordPoolDayAheadUpdate {
	mutation := newNordPoolDayAheadMutation(c.config, OpUpdate)
	return &NordPoolDayAheadUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NordPoolDayAheadClient) UpdateOne(npda *NordPoolDayAhead) *NordPoolDayAheadUpdateOne {
	mutation := newNordPoolDayAheadMutation(c.config, OpUpdateOne, withNordPoolDayAhead(npda))
	return &NordPoolDayAheadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NordPoolDayAheadClient) UpdateOneID(id int) *NordPoolDayAheadUpdateOne {
	mutation := newNordPoolDayAheadMutation(c.config, OpUpdateOne, withNordPoolDayAheadID(id))
	return &NordPoolDayAheadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for NordPoolDayAhead.
func (c *NordPoolDayAheadClient) Delete() *NordPoolDayAheadDelete {
	mutation := newNordPoolDayAheadMutation(c.config, OpDelete)
	return &NordPoolDayAheadDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NordPoolDayAheadClient) DeleteOne(npda *NordPoolDayAhead) *NordPoolDayAheadDeleteOne {
	return c.DeleteOneID(npda.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NordPoolDayAheadClient) DeleteOneID(id int) *NordPoolDayAheadDeleteOne {
	builder := c.Delete().Where(nordpooldayahead.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NordPoolDayAheadDeleteOne{builder}
}

// Query returns a query builder for NordPoolDayAhead.
func (c *NordPoolDayAheadClient) Query() *NordPoolDayAheadQuery {
	return &NordPoolDayAheadQuery{
		config: c.config,
	}
}

// Get returns a NordPoolDayAhead entity by its id.
func (c *NordPoolDayAheadClient) Get(ctx context.Context, id int) (*NordPoolDayAhead, error) {
	return c.Query().Where(nordpooldayahead.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NordPoolDayAheadClient) GetX(ctx context.Context, id int) *NordPoolDayAhead {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NordPoolDayAheadClient) Hooks() []Hook {
	return c.hooks.NordPoolDayAhead
}
