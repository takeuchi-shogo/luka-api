package apierrors

type AppError interface {
	BadRequest() AppError
	InvalidParameter() ApiError
	Unauthorized() AppError
	PaymentRequired() AppError
	NotFound() AppError
	TooManyRequests() AppError
	InternalServerError() AppError
	ServiceUnavailable() AppError

	New(msg ...string) AppError
}
