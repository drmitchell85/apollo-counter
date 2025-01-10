package utils

import "errors"

var (
	// general error messages
	ErrMissingFields     = errors.New("missing fields")
	ErrInvalidJsonFormat = errors.New("invalid JSON format")

	// user error messages
	ErrUserNotFound   = errors.New("user not found")
	ErrDuplicateEmail = errors.New("email already exists")

	// event error messages
	ErrDuplicateTitle = errors.New("title already exists")
)
