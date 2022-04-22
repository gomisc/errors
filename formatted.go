package errors

import (
	"encoding/json"
	"fmt"

	"git.corout.in/golibs/fields"
)

// FormattedError форматированная ошибка для вывода в результатах тестов
type FormattedError struct {
	message string
	fields  map[string]interface{}
}

// Formatted конструктор форматированной аннотированной полями ошибки
func Formatted(err error, args ...interface{}) error {
	if err == nil {
		return nil
	}

	message := err.Error()

	if len(args) != 0 {
		message = fmt.Sprintf(args[0].(string), args[1:]...)
	}

	fmtErr := &FormattedError{
		message: message,
		fields: map[string]interface{}{
			"message": err.Error(),
		},
	}

	if ctxErr, ok := err.(fields.Extractor); ok {
		ctxErr.Extract(fmtErr)
	}

	return fmtErr
}

func (f *FormattedError) Error() string {
	message := f.message

	b, err := json.MarshalIndent(f.fields, "", "  ")
	if err == nil {
		message += ": " + string(b)
	}

	return message
}

func (f *FormattedError) Bool(key string, value bool) {
	f.fields[key] = value
}

func (f *FormattedError) Int(key string, value int) {
	f.fields[key] = value
}

func (f *FormattedError) Int8(key string, value int8) {
	f.fields[key] = value
}

func (f *FormattedError) Int16(key string, value int16) {
	f.fields[key] = value
}

func (f *FormattedError) Int32(key string, value int32) {
	f.fields[key] = value
}

func (f *FormattedError) Int64(key string, value int64) {
	f.fields[key] = value
}

func (f *FormattedError) Uint(key string, value uint) {
	f.fields[key] = value
}

func (f *FormattedError) Uint8(key string, value uint8) {
	f.fields[key] = value
}

func (f *FormattedError) Uint16(key string, value uint16) {
	f.fields[key] = value
}

func (f *FormattedError) Uint32(key string, value uint32) {
	f.fields[key] = value
}

func (f *FormattedError) Uint64(key string, value uint64) {
	f.fields[key] = value
}

func (f *FormattedError) Float32(key string, value float32) {
	f.fields[key] = value
}

func (f *FormattedError) Float64(key string, value float64) {
	f.fields[key] = value
}

func (f *FormattedError) Str(key, value string) {
	f.fields[key] = value
}

func (f *FormattedError) Strings(key string, values []string) {
	f.fields[key] = values
}

func (f *FormattedError) Any(key string, value interface{}) {
	f.fields[key] = value
}
