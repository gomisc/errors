package errors

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

	return ch.Error()
}