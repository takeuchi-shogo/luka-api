package apierrors

func NewError(code, msg string) *ApiError {
	return &ApiError{
		Code:    code,
		Message: msg,
	}
}

func newBadRequest(code, msg string) *ApiError {
	e := NewError(code, msg)
	e.BadRequest()

	return e
}

func newUnauthorized(code, msg string) *ApiError {
	e := NewError(code, msg)
	e.Unauthorized()

	return e
}

func newPaymentRequired(code, msg string) *ApiError {
	e := NewError(code, msg)
	e.PaymentRequired()

	return e
}

func newNotFound(code, msg string) *ApiError {
	e := NewError(code, msg)
	e.NotFound()

	return e
}

func newTooManyRequests(code, msg string) *ApiError {
	e := NewError(code, msg)
	e.TooManyRequests()

	return e
}

func newInternalServerError(code, msg string) *ApiError {
	e := NewError(code, msg)
	e.InternalServerError()

	return e
}

func newServiceUnavailable(code, msg string) *ApiError {
	e := NewError(code, msg)
	e.ServiceUnavailable()

	return e
}
