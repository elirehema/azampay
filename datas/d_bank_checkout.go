package datas

type BankCheckout struct {
	AdditionalProperties  Properties `json:"additionalProperties" binding:"required"`
	Amount                string     `json:"amount" binding:"required"`
	CurrencyCode          string     `json:"currencyCode" binding:"required"`
	MerchantAccountNumber string     `json:"merchantAccountNumber" binding:"required"`
	MerchantMobileNumber  string     `json:"merchantMobileNumber" binding:"required"`
	MerchantName          string     `json:"merchantName" binding:"required"`
	ExternalId            string     `json:"externalId" binding:"required"`
	Otp                   string     `json:"otp" binding:"required"`
	Provider              string     `json:"provider" binding:"required"`
	ReferenceID           string     `json:"referenceId" binding:"required"`
}

func (b *BankCheckout) SetPassword(password string) {
	b.Otp = password
}
