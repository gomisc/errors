package errors

import (
	"testing"
)

func TestRecoverError(t *testing.T) {
	const testPanic = Const("test panic")

	tests := []struct {
		name      string
		panicFunc func(ch chan error)
		wantErr   bool
	}{
		{
			name: "panic",
			panicFunc: func(ch chan error) {
				defer func() {
					if r := recover(); r != nil {
						ch <- RecoverError(r)
					}
				}()

				panic(testPanic)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan error, 1)
			defer close(ch)

			tt.panicFunc(ch)

			if err := <-ch; err != nil {
				t.Error(err)
			} else {
				t.Error("expect panic error")
			}
		})
	}
}
