package datas

type ErrorMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (tr ErrorMessage) success() bool {
	return false
}

func ReturnErrorMessage(message string) *ErrorMessage {
	return &ErrorMessage{
		Status:  "Failied",
		Message: message,
	}
}
