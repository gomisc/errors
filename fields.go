package errors

import (
	"fmt"

	"git.corout.in/golibs/fields"
)

const (
	positionKey = "error-position"
)

func (we *wrappedError) Bool(key string, value bool) fields.FieldContainer {
	we.fields = append(we.fields, fields.Bool(key, value))

	return we
}

func (we *wrappedError) Int(key string, value int) fields.FieldContainer {
	we.fields = append(we.fields, fields.Int(key, value))

	return we
}

func (we *wrappedError) Int8(key string, value int8) fields.FieldContainer {
	we.fields = append(we.fields, fields.Int8(key, value))

	return we
}

func (we *wrappedError) Int16(key string, value int16) fields.FieldContainer {
	we.fields = append(we.fields, fields.Int16(key, value))

	return we
}

func (we *wrappedError) Int32(key string, value int32) fields.FieldContainer {
	we.fields = append(we.fields, fields.Int32(key, value))

	return we
}

func (we *wrappedError) Int64(key string, value int64) fields.FieldContainer {
	we.fields = append(we.fields, fields.Int64(key, value))

	return we
}

func (we *wrappedError) Uint(key string, value uint) fields.FieldContainer {
	we.fields = append(we.fields, fields.Uint(key, value))

	return we
}

func (we *wrappedError) Uint8(key string, value uint8) fields.FieldContainer {
	we.fields = append(we.fields, fields.Uint8(key, value))

	return we
}

func (we *wrappedError) Uint16(key string, value uint16) fields.FieldContainer {
	we.fields = append(we.fields, fields.Uint16(key, value))

	return we
}

func (we *wrappedError) Uint32(key string, value uint32) fields.FieldContainer {
	we.fields = append(we.fields, fields.Uint32(key, value))

	return we
}

func (we *wrappedError) Uint64(key string, value uint64) fields.FieldContainer {
	we.fields = append(we.fields, fields.Uint64(key, value))

	return we
}

func (we *wrappedError) Float32(key string, value float32) fields.FieldContainer {
	we.fields = append(we.fields, fields.Float32(key, value))

	return we
}

func (we *wrappedError) Float64(key string, value float64) fields.FieldContainer {
	we.fields = append(we.fields, fields.Float64(key, value))

	return we
}

func (we *wrappedError) Str(key string, value string) fields.FieldContainer {
	we.fields = append(we.fields, fields.Str(key, value))

	return we
}

func (we *wrappedError) Strings(key string, values []string) fields.FieldContainer {
	we.fields = append(we.fields, fields.Strings(key, values))

	return we
}

func (we *wrappedError) Stringer(key string, value fmt.Stringer) fields.FieldContainer {
	we.fields = append(we.fields, fields.Stringer(key, value))

	return we
}

func (we *wrappedError) Any(key string, value interface{}) fields.FieldContainer {
	we.fields = append(we.fields, fields.Any(key, value))

	return we
}

func (we *wrappedError) Extract(out fields.FieldExtractor) {
	if we.pos != nil {
		out.Str(positionKey, we.pos.String())
	}

	for i := 0; i < len(we.fields); i++ {
		key := we.fields[i].Key()
		we.fields[i].Value().Extract(key, out)
	}
}