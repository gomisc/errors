package errors

import (
	"git.corout.in/golibs/fields"
)

type CustomError interface {
	Just(err error) error
	New(reason string) error
	Newf(format string, args ...interface{}) error
	Wrap(err error, reason string) error
	Wrapf(err error, format string, args ...interface{}) error
	With(fields ...fields.Field) CustomError
	fields.FieldContainer
}
