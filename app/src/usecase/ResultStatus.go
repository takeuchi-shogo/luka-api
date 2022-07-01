package usecase

import "github.com/takeuchi-shogo/luka-api/src/domain"

type ResultStatus struct {
	ErrorMessage *domain.Error
	StatusCode   int // ステータスコード
}

func NewResultStatus(code int, errorCode int) *ResultStatus {

	resultStatus := new(ResultStatus)

	resultStatus.StatusCode = code
	resultStatus.ErrorMessage = domain.NewError(errorCode)

	return resultStatus
}
