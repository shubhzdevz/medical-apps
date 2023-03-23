package db

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// DB is a struct representing a database
type DB struct {
    client *mongo.Client
}

// NewDB returns a new instance of DB
func NewDB(ctx context.Context, uri string) (*DB, error) {
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }

    return &DB{
        client: client,
    }, nil
}

// Close closes the database connection
func (d *DB) Close(ctx context.Context) error {
    return d.client.Disconnect(ctx)
}

// Collection returns a collection from the database
func (d *DB) Collection(database, collection string) *mongo.Collection {
    return d.client.Database(database).Collection(collection)
}
