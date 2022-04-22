package errors

import (
	"bytes"
)

var _ error = Chain{}

// Chain - цепочка ошибок
type Chain []error

func (ch Chain) Error() string {
	chainSz := len(ch)

	if chainSz == 0 {
		panic("empty errors chain can't be use. don't forget to check it")
	}

	if chainSz == 1 {
		return ch[0].Error()
	}

	var buf bytes.Buffer

	for i := range ch {
		if i > 0 {
			buf.WriteString(";")
		}

		buf.WriteString(ch[i].Error())
	}

	return buf.String()
}

func (ch Chain) As(target interface{}) bool {
	for i := range ch {
		if As(ch[i], target) {
			return true
		}
	}

	return false
}

func (ch Chain) Is(err error) bool {
	for i := range ch {
		if Is(ch[i], err) {
			return true
		}
	}

	return false
}
