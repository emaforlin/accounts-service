package types

import "errors"

var (
	ErrInvalidToken       = errors.New("invalid token")
	ErrEmptyAuthorization = errors.New("received empty authorization from the client")

	ErrUnauthorized = errors.New("user unauthorized")
	ErrUserNotFound = errors.New("user not found")

	ErrFailedCreateAccount = errors.New("create new account failed")
)
