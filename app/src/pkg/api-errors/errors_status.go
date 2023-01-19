package apierrors

import "net/http"

func (e *ApiError) BadRequest() {
	e.StatusCode = http.StatusBadRequest
	// return e
}

func (e *ApiError) Unauthorized() {
	e.StatusCode = http.StatusUnauthorized
	// return e
}

func (e *ApiError) PaymentRequired() {
	e.StatusCode = http.StatusPaymentRequired
	// return e
}

func (e *ApiError) NotFound() {
	e.StatusCode = http.StatusNotFound
	// return e
}

func (e *ApiError) TooManyRequests() {
	e.StatusCode = http.StatusTooManyRequests
	// return e
}

func (e *ApiError) InternalServerError() {
	e.StatusCode = http.StatusInternalServerError
	// return e
}

func (e *ApiError) ServiceUnavailable() {
	e.StatusCode = http.StatusServiceUnavailable
	// return e
}
