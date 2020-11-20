package domain

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")

	// ErrNotFound will throw if the requested item does not exist
	ErrNotFound = errors.New("Your requested Item is not found")

	ErrRecordNotFound = errors.New("Record not found")

	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your item already exist")

	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given param is not valid")
)
