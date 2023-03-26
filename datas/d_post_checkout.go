package datas

type PostCheckoutRequest struct {
	AppName            string `json:"appName"`
	ClientID           string `json:"clientId"`
	VendorID           string `json:"vendorId"`
	Language           string `json:"language"`
	Currency           string `json:"currency"`
	ExternalID         string `json:"externalId"`
	RequestOrigin      string `json:"requestOrigin"`
	RedirectFailURL    string `json:"redirectFailURL"`
	RedirectSuccessURL string `json:"redirectSuccessURL"`
	VendorName         string `json:"vendorName"`
	Amount             string `json:"amount"`
	Cart               struct {
		Items []struct {
			Name string `json:"name"`
		} `json:"items"`
	} `json:"cart"`
}
