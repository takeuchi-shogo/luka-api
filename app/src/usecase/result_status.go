package usecase

import "errors"

type ResultStatus struct {
	ErrorMessage error
	StatusCode   int // ステータスコード
}

func NewResultStatus(code int, errorMessage string) *ResultStatus {

	resultStatus := new(ResultStatus)

	resultStatus.StatusCode = code
	resultStatus.ErrorMessage = errors.New(errorMessage)

	return resultStatus
}
