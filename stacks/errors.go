package stacks

import "errors"

var (
	// ErrHostAlreadyExists occurs when a given host which should get added already exists in the current stack
	ErrHostAlreadyExists = errors.New("the given host already exists")

	// ErrHostDoesNotExist occurs when a given host does not exist in the current stack
	ErrHostDoesNotExist = errors.New("the given host does not exist")
)
