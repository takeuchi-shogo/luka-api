package usecase

type ResultStatus struct {
	Error      error
	StatusCode int // ステータスコード
	Message    string
}

func NewResultStatus(code int, err error, message string) *ResultStatus {

	resultStatus := new(ResultStatus)

	resultStatus.StatusCode = code
	resultStatus.Message = message
	if err != nil {
		resultStatus.Error = err
	} else {
		resultStatus.Error = nil
	}

	return resultStatus
}
