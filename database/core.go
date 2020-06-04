package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// CurrentClient defines the current MongoDB client
var CurrentClient *mongo.Client

// Connect established a connection to the given MongoDB host
func Connect(connectionURI string) error {
	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB host
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		return err
	}

	// Ping the MongoDB host
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	CurrentClient = client
	return nil
}
