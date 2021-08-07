package errors

// Константные ошибки

// Const тип используемый для константных ошибок, позволяет избегать возможных мутаций значений ошибок.
type Const string

func (e Const) Error() string {
	return string(e)
}
