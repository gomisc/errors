package errors

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"strings"

	"git.corout.in/golibs/fields"
)

var (
	_ error       = (*wrappedError)(nil)
	_ CustomError = (*wrappedError)(nil)
)

type position struct {
	fname string
	file  string
	line  int
}

func (p *position) String() string {
	return fmt.Sprintf("%s %s:%d", p.fname, p.file, p.line)
}

type wrappedError struct {
	err     error
	pos     *position
	reasons []string
	fields  []fields.Field
}

func newerr() *wrappedError {
	return &wrappedError{}
}

func (we *wrappedError) Just(err error) error {
	if err == nil {
		panic("wrapping error must be not nil")
	}

	var werr *wrappedError

	if As(err, &werr) {
		werr.fields = append(werr.fields, we.fields...)

		return werr
	} else {
		if we.pos == nil {
			we.position(1)
		}

		we.err = err

		return we
	}
}

func (we *wrappedError) New(reason string) error {
	if we.pos == nil {
		we.position(1)
	}

	we.err = errors.New(reason)

	return we
}

func (we *wrappedError) Newf(format string, args ...interface{}) error {
	if we.pos == nil {
		we.position(1)
	}

	we.err = fmt.Errorf(format, args...)

	return we
}

func (we *wrappedError) Wrap(err error, reason string) error {
	if err == nil {
		panic("wrapping error must be not nil")
	}

	var werr *wrappedError

	if As(err, werr) {
		werr.reasons = append(werr.reasons, reason)
		werr.fields = append(werr.fields, we.fields...)

		return werr
	} else {
		if we.pos == nil {
			we.position(1)
		}

		we.err = err
		we.reasons = append(we.reasons, reason)

		return we
	}
}

func (we *wrappedError) Wrapf(err error, format string, args ...interface{}) error {
	if we.pos == nil {
		we.position(1)
	}

	return we.Wrap(err, fmt.Sprintf(format, args...))
}

func (we *wrappedError) Error() string {
	var buf bytes.Buffer

	for i := len(we.reasons) - 1; i >= 0; i-- {
		buf.WriteString(we.reasons[i] + ": ")
	}

	buf.WriteString(we.err.Error())

	return buf.String()
}

func (we *wrappedError) As(target interface{}) bool {
	switch val := target.(type) {
	case **wrappedError:
		*val = we

		return true
	default:
		return false
	}
}

func (we *wrappedError) position(depth int) CustomError {
	pc, file, line, _ := runtime.Caller(depth + 1) // nolint: gomnd, dogsled

	if we.pos == nil {
		we.pos = &position{line: line}
	}

	fname := runtime.FuncForPC(pc).Name()
	we.pos.fname = fname[strings.Index(fname, ".")+1:]

	if fp := strings.Index(file, "/src/"); fp >= 0 {
		we.pos.file = file[fp+5:]
	} else if fp = strings.Index(file, "/pkg/mod/"); fp >= 0 {
		we.pos.file = file[fp+9:]
	}
	return we
}
