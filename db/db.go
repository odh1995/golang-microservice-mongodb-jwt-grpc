package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type conn struct {
	client   *mongo.Client
	ctx      *context.Context
	cancel   context.CancelFunc
	database *mongo.Database
}

// This is a user defined method that returns mongo.Client,
// context.Context, context.CancelFunc and error.
// mongo.Client will be used for further database operation.
// context.Context will be used set deadlines for process.
// context.CancelFunc will be used to cancel context and
// resource associated with it.

func NewConnection(cfg Config) (Connection, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Dsn()))
	if err != nil {
		panic(err)
	}
	return &conn{client: client, ctx: &ctx, cancel: cancel, database: client.Database(cfg.DbName())}, nil
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func Ping(c *conn) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := c.client.Ping(*c.ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

// Thisb is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func (c *conn) Close() {

	// CancelFunc to cancel to context
	defer c.cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := c.client.Disconnect(*c.ctx); err != nil {
			panic(err)
		}
	}()
}

func (c *conn) DB() *mongo.Database {
	return c.database
}
