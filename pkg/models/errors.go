package models

import (
	"errors"
	"fmt"
)

// ErrUserAlreadyExist represents a "user already exists" error.
type ErrUserAlreadyExist struct {
	UID   int64
	Email string
}

// IsErrUserAlreadyExist checks if an error is a ErrUserAlreadyExists.
func IsErrUserAlreadyExist(err error) bool {
	return errors.As(err, &ErrUserAlreadyExist{})
}

func (err ErrUserAlreadyExist) Error() string {
	return fmt.Sprintf("user already exists [email: %s]", err.Email)
}

// ErrUserNotExist represents a "UserNotExist" kind of error.
type ErrUserNotExist struct {
	UID   int64
	Email string
}

// IsErrUserNotExist checks if an error is a ErrUserNotExist.
func IsErrUserNotExist(err error) bool {
	return errors.As(err, &ErrUserNotExist{})
}

func (err ErrUserNotExist) Error() string {
	return fmt.Sprintf("user does not exist [uid: %d, email: %s]", err.UID, err.Email)
}
