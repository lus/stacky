package stacks

import (
	"context"
	"time"

	"github.com/Lukaesebrot/stacky/config"
	"github.com/Lukaesebrot/stacky/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create creates a new stack
func Create(name string, hosts ...string) (*Stack, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(config.CurrentConfig.MongoDBDatabase).Collection("stacks")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create the stack structure
	stack := &Stack{
		Name:  name,
		Hosts: hosts,
	}

	// Create the stack document
	result, err := collection.InsertOne(ctx, stack)
	if err != nil {
		return nil, err
	}

	// Define the object ID and return the stack structure
	stack.ID = result.InsertedID.(primitive.ObjectID)
	return stack, err
}

// AddHost adds the given host to the current stack
func AddHost(host string) error {
	// TODO: Implement host addition logic
	return nil
}

// RemoveHost removes the given host from the current stack
func RemoveHost(host string) error {
	// TODO: Implement host removing logic
	return nil
}

// Update writes the current local variables of the current stack into the database
func (stack *Stack) Update() error {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(config.CurrentConfig.MongoDBDatabase).Collection("stacks")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update the MongoDB document
	filter := bson.M{"_id": stack.ID}
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": stack})
	return err
}

// Delete deletes the current stack
func (stack *Stack) Delete() error {
	// TODO: Implement stack deletion logic
	return nil
}
