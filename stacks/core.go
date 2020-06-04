package stacks

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Lukaesebrot/stacky/config"
	"github.com/Lukaesebrot/stacky/database"
	"github.com/Lukaesebrot/stacky/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Retrieve retrieves the stack with the given name
func Retrieve(name string) (*Stack, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(config.CurrentConfig.MongoDBDatabase).Collection("stacks")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the stack with the given name
	filter := bson.M{"name": name}
	result := collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		return nil, err
	}

	// Decode the document into a stack structure
	stack := new(Stack)
	err = result.Decode(stack)
	return stack, err
}

// Create creates a new stack
func Create(name string, hosts ...string) (*Stack, error) {
	// Validate the stack name
	stack, err := Retrieve(name)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	if stack != nil {
		return nil, ErrStackAlreadyExists
	}

	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(config.CurrentConfig.MongoDBDatabase).Collection("stacks")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create the stack structure
	stack = &Stack{
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
func (stack *Stack) AddHost(host string) error {
	for _, hst := range stack.Hosts {
		if hst == host {
			return ErrHostAlreadyExists
		}
	}
	stack.Hosts = append(stack.Hosts, host)
	return stack.Update()
}

// RemoveHost removes the given host from the current stack
func (stack *Stack) RemoveHost(host string) error {
	index, contains := utils.StringArrayContains(stack.Hosts, host)
	if !contains {
		return ErrHostDoesNotExist
	}
	stack.Hosts = append(stack.Hosts[:index], stack.Hosts[index+1:]...)
	return stack.Update()
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
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(config.CurrentConfig.MongoDBDatabase).Collection("stacks")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Delete the document
	filter := bson.M{"_id": stack.ID}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
