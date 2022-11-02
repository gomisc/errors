package errors

import (
	"fmt"

	"git.eth4.dev/golibs/fields"
)

const (
	positionKey = "position"
	stackKey    = "stack"
)

// Bool имплементация ContextError
func (we *wrappedError) Bool(key string, value bool) ContextError {
	we.fields = append(we.fields, fields.Bool(key, value))

	return we
}

// Int имплементация ContextError
func (we *wrappedError) Int(key string, value int) ContextError {
	we.fields = append(we.fields, fields.Int(key, value))

	return we
}

// Int8 имплементация ContextError
func (we *wrappedError) Int8(key string, value int8) ContextError {
	we.fields = append(we.fields, fields.Int8(key, value))

	return we
}

// Int16 имплементация ContextError
func (we *wrappedError) Int16(key string, value int16) ContextError {
	we.fields = append(we.fields, fields.Int16(key, value))

	return we
}

// Int32 имплементация ContextError
func (we *wrappedError) Int32(key string, value int32) ContextError {
	we.fields = append(we.fields, fields.Int32(key, value))

	return we
}

// Int64 имплементация ContextError
func (we *wrappedError) Int64(key string, value int64) ContextError {
	we.fields = append(we.fields, fields.Int64(key, value))

	return we
}

// Uint имплементация ContextError
func (we *wrappedError) Uint(key string, value uint) ContextError {
	we.fields = append(we.fields, fields.Uint(key, value))

	return we
}

// Uint8 имплементация ContextError
func (we *wrappedError) Uint8(key string, value uint8) ContextError {
	we.fields = append(we.fields, fields.Uint8(key, value))

	return we
}

// Uint16 имплементация ContextError
func (we *wrappedError) Uint16(key string, value uint16) ContextError {
	we.fields = append(we.fields, fields.Uint16(key, value))

	return we
}

// Uint32 имплементация ContextError
func (we *wrappedError) Uint32(key string, value uint32) ContextError {
	we.fields = append(we.fields, fields.Uint32(key, value))

	return we
}

// Uint64 имплементация ContextError
func (we *wrappedError) Uint64(key string, value uint64) ContextError {
	we.fields = append(we.fields, fields.Uint64(key, value))

	return we
}

// Float32 имплементация ContextError
func (we *wrappedError) Float32(key string, value float32) ContextError {
	we.fields = append(we.fields, fields.Float32(key, value))

	return we
}

// Float64 имплементация ContextError
func (we *wrappedError) Float64(key string, value float64) ContextError {
	we.fields = append(we.fields, fields.Float64(key, value))

	return we
}

// Str имплементация ContextError
func (we *wrappedError) Str(key, value string) ContextError {
	we.fields = append(we.fields, fields.Str(key, value))

	return we
}

// Strings имплементация ContextError
func (we *wrappedError) Strings(key string, values []string) ContextError {
	we.fields = append(we.fields, fields.Strings(key, values))

	return we
}

// Stringer имплементация ContextError
func (we *wrappedError) Stringer(key string, value fmt.Stringer) ContextError {
	we.fields = append(we.fields, fields.Stringer(key, value))

	return we
}

// Any имплементация ContextError
func (we *wrappedError) Any(key string, value interface{}) ContextError {
	we.fields = append(we.fields, fields.Any(key, value))

	return we
}

func (we *wrappedError) WithStack() ContextError {
	if we.pos != nil {
		we.fields = append(
			we.fields,
			fields.Str(stackKey, we.pos.stack),
		)
	}

	return we
}

// Extract имплементация fields.FieldExtractor
func (we *wrappedError) Extract(out fields.FieldExtractor) {
	for i := 0; i < len(we.fields); i++ {
		key := we.fields[i].Key()
		we.fields[i].Value().Extract(key, out)
	}

	if we.pos != nil {
		out.Str(positionKey, fmt.Sprintf("%s:%d", we.pos.file, we.pos.line))
	}
}
