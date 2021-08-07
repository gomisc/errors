package errors

import (
	"errors"
)

func Builder() CustomError {
	return &wrappedError{}
}

func New(reason string) error {
	return newerr().position(1).New(reason)
}

func Newf(format string, args ...interface{}) error {
	return newerr().position(1).Newf(format, args...)
}

func Wrap(err error, reason string) error {
	return newerr().position(1).Wrap(err, reason)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return newerr().position(1).Wrapf(err, format, args...)
}

// Is stdlib errors.Is wrapper
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As stdlib errors.As wrapper
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
