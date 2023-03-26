package datas

type TransactionResponse struct {
	Success       bool             `json:"success"`
	TransactionId string           `json:"transactionId,omitempty"`
	Message       string           `json:"message"`
	Data          *TransactionData `json:"data,omitempty"`
}

func (tr TransactionResponse) success() bool {
	return true
}
