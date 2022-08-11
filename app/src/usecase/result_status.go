package usecase

import "errors"

type ResultStatus struct {
	ErrorMessage error
	StatusCode   int // ステータスコード
}

func NewResultStatus(code int, errorMessage string) *ResultStatus {

	resultStatus := new(ResultStatus)

	resultStatus.StatusCode = code
	if errorMessage != "" {
		resultStatus.ErrorMessage = errors.New(errorMessage)
	} else {
		resultStatus.ErrorMessage = nil
	}

	return resultStatus
}
