package errors

import (
	"bytes"
	goErrors "errors"
	"fmt"
	"runtime"
	"strings"

	"git.eth4.dev/golibs/fields"
)

var (
	_ error        = (*wrappedError)(nil)
	_ ContextError = (*wrappedError)(nil)
)

type position struct {
	file  string
	stack string
	line  int
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

// Just - возвращает обернутую ошибку в контекст без доавления мессаджа
func (we *wrappedError) Just(err error) error {
	if err == nil {
		panic("wrapping error must be not nil")
	}

	var werr *wrappedError

	if As(err, &werr) {
		werr.fields = append(werr.fields, we.fields...)

		return werr
	}

	if we.pos == nil {
		we.position(1)
	}

	we.err = err

	return we
}

// New - конструктор новой контекстной ошибки
func (we *wrappedError) New(reason string) error {
	if we.pos == nil {
		we.position(1)
	}

	we.err = goErrors.New(reason)

	return we
}

// Newf - конструктор новой контекстной ошибки с форматируемым мессаджем
func (we *wrappedError) Newf(format string, args ...any) error {
	if we.pos == nil {
		we.position(1)
	}

	we.err = fmt.Errorf(format, args...)

	return we
}

// Wrap - оборачивает ошибку предыдущего уровня
func (we *wrappedError) Wrap(err error, reason string) error {
	if err == nil {
		panic("wrapping error must be not nil")
	}

	var werr *wrappedError

	if As(err, &werr) {
		werr.reasons = append(werr.reasons, reason)
		werr.fields = append(werr.fields, we.fields...)

		return werr
	}

	if we.pos == nil {
		we.position(1)
	}

	we.err = err
	we.reasons = append(we.reasons, reason)

	return we
}

// Wrapf - оборачивает ошибку предыдущего уровня с форматированным мессаджем
func (we *wrappedError) Wrapf(err error, format string, args ...any) error {
	if we.pos == nil {
		we.position(1)
	}

	return we.Wrap(err, fmt.Sprintf(format, args...))
}

// Error - имплементация интерфейсу error
func (we *wrappedError) Error() string {
	var buf bytes.Buffer

	for i := len(we.reasons) - 1; i >= 0; i-- {
		buf.WriteString(we.reasons[i] + ": ")
	}

	buf.WriteString(we.err.Error())

	return buf.String()
}

// As - выражает обернутую ошиюку конкретному типу
func (we *wrappedError) As(target interface{}) bool {
	switch val := target.(type) {
	case **wrappedError:
		*val = we

		return true
	default:
		return false
	}
}

func (we *wrappedError) Pos(depth int) ContextError {
	return we.position(depth)
}

// With - добавляет поля контекста в ошибку
func (we *wrappedError) With(flds ...fields.Field) ContextError {
	we.fields = append(we.fields, flds...)

	return we
}

func (we *wrappedError) Unwrap() error {
	return we.err
}

// принудительно устанавливает/меняет позицию (место возникновения)
// ошибки на указанную
// nolint: unparam
func (we *wrappedError) position(depth int) ContextError {
	pc, file, line, _ := runtime.Caller(depth + 1) // nolint: gomnd, dogsled

	if we.pos == nil {
		we.pos = &position{line: line}
	}

	buf := make([]byte, 1024)

	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			we.pos.stack = formatStack(buf[:n], runtime.FuncForPC(pc).Name())
			break
		}

		buf = make([]byte, 2*len(buf))
	}

	if fp := strings.Index(file, "/src/"); fp >= 0 {
		we.pos.file = file[fp+5:]
	} else if fp = strings.Index(file, "/pkg/mod/"); fp >= 0 {
		we.pos.file = file[fp+9:]
	}

	return we
}

func formatStack(buf []byte, caller string) string {
	lines := bytes.Split(bytes.ReplaceAll(buf, []byte("\t"), []byte("")), []byte("\n"))
	result := make([][]byte, 0, len(lines))

	write := false
	for i := range lines {
		if !write && bytes.Contains(lines[i], []byte(caller)) {
			write = true
		}

		if write && !bytes.EqualFold(lines[i], []byte("")) {
			result = append(result, lines[i])
		}
	}

	return string(bytes.Join(result, []byte(" ")))
}
