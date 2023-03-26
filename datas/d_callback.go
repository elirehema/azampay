package datas

type CallBackResponse struct {
	Message           string      `json:"message"`
	User              string      `json:"user"`
	Password          string      `json:"password"`
	ClientID          string      `json:"clientId"`
	Transactionstatus string      `json:"transactionstatus"`
	Operator          string      `json:"operator"`
	Reference         string      `json:"reference"`
	Externalreference string      `json:"externalreference"`
	Utilityref        string      `json:"utilityref"`
	Amount            string      `json:"amount"`
	Transid           string      `json:"transid"`
	Msisdn            string      `json:"msisdn"`
	SubmerchantAcc    string      `json:"submerchantAcc"`
	Properties        *Properties `json:"additionalProperties"`
}
