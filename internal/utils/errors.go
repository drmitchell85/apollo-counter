package utils

import "errors"

// user error messages
var (
	ErrMissingFields     = errors.New("missing fields")
	ErrUserNotFound      = errors.New("user not found")
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrInvalidJsonFormat = errors.New("invalid JSON format")
)
