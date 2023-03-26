package datas

import (
	"fmt"

	"github.com/google/uuid"
)

type CheckoutRequest struct {
	Amount         string `json:"amount" binding:"required"`
	Provider       string `json:"provider" binding:"required"`
	MechantAccount string `json:"mechantAccountNo,omitempty"`
	Password       string `json:"password"`
	TenantId       string `json:"tenantId"  binding:"required"`
	TenantName     string `json:"tenantName"  binding:"required"`
	MobileNumber   string `json:"mobileNumber"  binding:"required"`
	LoanId         string `json:"loanId"  binding:"required"`
	ClientId       string `json:"clientId" binding:"required"`
}

func (r *CheckoutRequest) CreateMnoRequest() MNOCheckout {
	return MNOCheckout{
		Currency:      "TZS",
		ExternalID:    uuid.New().String(),
		Provider:      r.Provider,
		AccountNumber: r.MobileNumber,
		Amount:        r.Amount,
		AdditionalProperties: Properties{
			PropertyOne: fmt.Sprintf("Tenant-Name:%v", r.TenantName),
			PropertyTwo: fmt.Sprintf("Tenant-Id:%v", r.TenantId),
		},
	}
}

func (b *CheckoutRequest) CreateBankRequest() BankCheckout {
	return BankCheckout{
		Amount:                b.Amount,
		CurrencyCode:          "TZS",
		MerchantAccountNumber: b.MechantAccount,
		MerchantMobileNumber:  b.MobileNumber,
		MerchantName:          "Amala Core banking",
		Provider:              b.Provider,
		ReferenceID:           uuid.New().String(),

		AdditionalProperties: Properties{
			PropertyOne: "Tenant-ID",
			PropertyTwo: "Tenant-Name",
		},
	}
}
