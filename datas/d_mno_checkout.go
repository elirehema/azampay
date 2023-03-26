package datas

type MNOCheckout struct {
	AccountNumber        string     `json:"accountNumber" binding:"required"`
	AdditionalProperties Properties `json:"additionalProperties" binding:"required"`
	Amount               string     `json:"amount" binding:"required"`
	Currency             string     `json:"currency" binding:"required"`
	ExternalID           string     `json:"externalId" binding:"required"`
	Provider             string     `json:"provider" binding:"required"`
}
