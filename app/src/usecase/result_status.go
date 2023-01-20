package usecase

type ResultStatus struct {
	Error      error  // error message
	StatusCode int    // status code
	Message    string // info error message
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
