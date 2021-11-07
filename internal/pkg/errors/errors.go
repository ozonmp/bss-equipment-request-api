package internal_errors

import "errors"

var (
	// ErrNotFound is a item not found error
	ErrNotFound = errors.New("not found")
	// ErrNotCreated is a item not created error
	ErrNotCreated = errors.New("unable to create")
	// ErrNotRemoved is a item not removed error
	ErrNotRemoved = errors.New("unable to remove")

	// ErrNotImplementedMethod is a function not implemented error
	ErrNotImplementedMethod = errors.New("method is not implemented")
)
