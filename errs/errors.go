package errs

import "errors"

var (
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrFailedToCreateUser    = errors.New("failed to create user")
	ErrFailedToHashPassword  = errors.New("failed to hash password")
	ErrFailedToGenerateToken = errors.New("failed to generate token")
	ErrFailedToSignToken     = errors.New("failed to signed token")
)

var (
	ErrUserNotFound = errors.New("user not found")
)
