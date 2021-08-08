package errors

import (
	"errors"

	"git.corout.in/golibs/fields"
)

func Builder(fields ...fields.Field) CustomError {
	return &wrappedError{
		fields: fields,
	}
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

// And собирает ошибки в цепочку
func And(err1, err2 error) error {
	if err2 == nil {
		return err1
	}

	if err1 == nil {
		return err2
	}

	var (
		ch1, ch2 Chain
	)

	if !As(err1, &ch1) {
		ch1 = make(Chain, 0, 4)
		ch1 = append(ch1, err1)
	}

	if As(err2, &ch2) {
		ch1 = append(ch1, ch2...)
	} else {
		ch1 = append(ch1, err2)
	}

	return ch1
}

// AsChain ищет тип `errors.Chain` во вложенных ошибках и возвращает его.
// Если цепочки не найдено, то создаётся новая, состоящая из текущей ошибки.
func AsChain(err error) Chain {
	var (
		w     *wrappedError
		chain Chain
	)

	if As(err, &w) && As(w, &chain) {
		newChain := make(Chain, len(chain))
		for i, e := range chain {
			newChain[i] = &wrappedError{
				reasons: w.reasons,
				err:  e,
				fields: w.fields,
			}
		}

		return newChain
	}

	if !As(err, &chain) {
		return Chain{err}
	}

	return chain
}

func Extract(err error) []fields.Field {
	var we *wrappedError

	if As(err, &we) {
		return we.fields
	}

	return []fields.Field{fields.Error(err)}
}
