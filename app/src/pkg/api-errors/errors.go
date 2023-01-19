package apierrors

type ApiError struct {
	// Return a typical error
	Next        error  // err
	Message     string // invalid string value: 'asdf'.
	InfoMessage string // 入力項目が無効です
	StatusCode  int    // 400
	Code        string // invalid_parameter

	// Errors []ErrorItems
	level string
}

func (e ApiError) New(err error, msg string) ApiError {
	// err := errors.New(msg)
	e.Next = err
	e.InfoMessage = msg

	return e
}
