package errors

import (
	goErrors "errors"
	"fmt"

	"git.corout.in/golibs/fields"
)

// ContextError - интерфейс контекстной (анотированной полями) ошибки
type ContextError interface {
	Just(err error) error
	New(reason string) error
	Newf(format string, args ...interface{}) error
	Wrap(err error, reason string) error
	Wrapf(err error, format string, args ...interface{}) error
	With(flds ...fields.Field) ContextError
	Pos(depth int) ContextError

	Bool(key string, value bool) ContextError
	Int(key string, value int) ContextError
	Int8(key string, value int8) ContextError
	Int16(key string, value int16) ContextError
	Int32(key string, value int32) ContextError
	Int64(key string, value int64) ContextError
	Uint(key string, value uint) ContextError
	Uint8(key string, value uint8) ContextError
	Uint16(key string, value uint16) ContextError
	Uint32(key string, value uint32) ContextError
	Uint64(key string, value uint64) ContextError
	Float32(key string, value float32) ContextError
	Float64(key string, value float64) ContextError
	Str(key string, value string) ContextError
	Strings(key string, values []string) ContextError
	Stringer(key string, value fmt.Stringer) ContextError
	Any(key string, value interface{}) ContextError
	Extract(out fields.FieldExtractor)
}

// Ctx - конструтор контекстной ошибки для определния полей
// в блоках условий для последующей обертки или формирования как есть
func Ctx() ContextError {
	return &wrappedError{}
}

func New(reason string) error {
	return newerr().New(reason)
}

func Wrap(err error, message string) error {
	return newerr().position(1).Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return newerr().position(1).Wrapf(err, format, args...)
}

func Is(err, target error) bool {
	return goErrors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return goErrors.As(err, target)
}

// And собирает ошибки в цепочку
func And(err1, err2 error) error {
	if err2 == nil {
		return err1
	}

	if err1 == nil {
		return err2
	}

	var ch1, ch2 Chain

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
				err:     e,
				fields:  w.fields,
			}
		}

		return newChain
	}

	if !As(err, &chain) && err != nil {
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
