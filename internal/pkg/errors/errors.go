package internal_errors

import "errors"

var (
	ErrNotFound   = errors.New("not found")
	ErrNotCreated = errors.New("unable to create")
	ErrNotRemoved = errors.New("unable to remove")

	ErrNotImplementedMethod = errors.New("method is not implemented")
)
