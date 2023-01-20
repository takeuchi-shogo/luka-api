package controllers

type ErrorResponse struct {
	Error struct {
		Message     string `json:"message"`
		InfoMessage string `json:"infoMessage"`
	} `json:"error"`
}

func NewErrorResponse(err error, msg string) *ErrorResponse {
	res := new(ErrorResponse)
	res.Error.Message = err.Error()
	res.Error.InfoMessage = msg

	return res
}
