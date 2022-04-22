package errors

import (
	"fmt"
	"runtime"
)

type ErrRecover struct {
	reason interface{}
	stack  string
}

func (e ErrRecover) Error() string {
	return fmt.Sprintf("%s\n%s", e.reason, e.stack)
}

func RecoverError(reason interface{}) error {
	err := &ErrRecover{reason: reason}
	buf := make([]byte, 1024)

	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			err.stack = string(buf[:n])

			return err
		}

		buf = make([]byte, 2*len(buf))
	}
}
