package datas

type CheckoutRequest struct {
	Amount          string `json:"amount" binding:"required"`
	Provider        string `json:"provider" binding:"required"`
	AccountNumber   string `json:"accountNo,omitempty"`
	MechantName     string `json:"mechantName"`
	CurrencyCode    string `json:"currency"`
	OneTimePassword string `json:"otp"`
	ExternalId      string `json:"externalId"`
	MobileNumber    string `json:"mobileNumber"  binding:"required"`
	Properties      Properties
}

func (r *CheckoutRequest) CreateMnoRequest() MNOCheckout {
	return MNOCheckout{
		Currency:             r.CurrencyCode,
		ExternalID:           r.ExternalId,
		Provider:             r.Provider,
		AccountNumber:        r.MobileNumber,
		Amount:               r.Amount,
		AdditionalProperties: r.Properties,
	}
}

func (b *CheckoutRequest) CreateBankRequest() BankCheckout {
	return BankCheckout{
		Amount:                b.Amount,
		CurrencyCode:          b.CurrencyCode,
		MerchantAccountNumber: b.AccountNumber,
		MerchantMobileNumber:  b.MobileNumber,
		MerchantName:          b.MechantName,
		Provider:              b.Provider,
		Otp:                   b.OneTimePassword,
		ReferenceID:           b.ExternalId,
		AdditionalProperties:  b.Properties,
	}
}
