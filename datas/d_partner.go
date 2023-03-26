package datas

type PaymentPartner struct {
	ID                         string `json:"id"`
	LogoUrl                    string `json:"logoUrl"`
	PartnerName                string `json:"partnerName"`
	Provider                   int    `json:"provider"`
	VendorName                 string `json:"vendorName"`
	PaymentVendorID            string `json:"paymentVendorId"`
	PaymentPartnerID           string `json:"paymentPartnerId"`
	PaymentAcknowledgmentRoute string `json:"paymentAcknowledgmentRoute"`
	Currency                   string `json:"currency"`
	Status                     string `json:"status"`
	VendorType                 string `json:"vendorType"`
}

func (p *PaymentPartner) success() bool {
	return true
}
