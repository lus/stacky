package stacks

import "go.mongodb.org/mongo-driver/bson/primitive"

// Stack represents a stack of API instances
type Stack struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Hosts []string           `bson:"hosts"`
}
