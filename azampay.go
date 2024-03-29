package azampay

import (
	"github.com/elirehema/azampay/clients"
	"github.com/elirehema/azampay/datas"
)

/**
* Checkout and make payment to requested MNO provider
* Ref:  https://developerdocs.azampay.co.tz/redoc#tag/Checkout-API/operation/Mno%20Checkout
 */
func MNOCheckout(request datas.CheckoutRequest) datas.ResponseData {
	result := datas.TransactionResponse{}
	body := request.CreateMnoRequest()
	_, err := clients.AzamClient.R().SetResult(&result).SetBody(body).Post("/azampay/mno/checkout")
	if err != nil {
		return datas.ReturnErrorMessage(err.Error())
	}
	return &result

}

/**
* Checkout and make payment to requested bank provider
* Ref:  https://developerdocs.azampay.co.tz/redoc#tag/Checkout-API/operation/Bank%20Checkout
 */
func BankCheckout(request datas.CheckoutRequest) datas.ResponseData {
	result := datas.TransactionResponse{}
	body := request.CreateBankRequest()
	_, err := clients.AzamClient.R().SetResult(&result).SetBody(body).Post("/azampay/bank/checkout")
	if err != nil {
		return datas.ReturnErrorMessage(err.Error())
	}
	return &result
}

/**
* This end point will respond back with the URL of your payments.
* Merchant Application can open this url in a new window to continue with the checkout process of the transaction
* Ref: https://developerdocs.azampay.co.tz/redoc#tag/Checkout-API/operation/Bank%20Checkout
**/
func PostCheckout(request datas.PostCheckoutRequest) string {
	result := ""
	res, err := clients.AzamClient.R().SetResult(&result).SetBody(request).Post("/Partner/PostCheckout")
	if err != nil {
		print("STATUS CODE {}:", res.StatusCode(), res.Request.URL)
		return ""
	}
	return result
}

func GetPaymentPartners() []datas.PaymentPartner {
	results := []datas.PaymentPartner{}
	res, err := clients.AzamClient.R().SetResult(&results).Get("/api/v1/Partner/GetPaymentPartners")
	if err != nil {
		print("STATUS CODE {}:", res.StatusCode(), res.Request.URL)
	}
	return results

}
