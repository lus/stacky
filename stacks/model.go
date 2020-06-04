package stacks

import "go.mongodb.org/mongo-driver/bson/primitive"

// Stack represents a stack of API instances
type Stack struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Hosts []string           `bson:"hosts" json:"hosts"`
}
