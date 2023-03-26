# Azampay Golang package (Beta Version)

# NOTE:
**This package was not developed, owned or maintained by Azam Pay Development Team**
The package was created as a open source project for Golang developers who are looking to consume [Azam Pay API's](https://developerdocs.azampay.co.tz/redoc). 

## Now What ?
### What is Azam Pay
---
> AzamPay is specialized in the development of end-to-end online payment management solutions for companies operating in East Africa. Our range of digital solutions and services are carefully designed not only to streamline your payment and collection processes, but to also allow easy integration with your current Accounting or Enterprise Resource Planning (ERP) systems thus leaving you time to focus on your customers. AzamPay offers bespoke solutions that guarantee optimal business performance and efficiency whether you are transacting locally, regionally, or internationally.

## Usage

Put the azampay package on your `import` statement:

```go
import "github.com/elirehema/azampay"
```

Download and modify the [.env file](https://github.com/elirehema/azampay/blob/master/.env) with the Azam Client information provided in you Azam Mechant application. Put the `.env` file in the root directory of your golang app. With right information provided in `.env` file the package will take care or creating and generating the access token required to access the Azampay public endpoints. 

Your `.env` file will looks like this
```.env
# AZAMPAY CLIENT
APP_NAME={{YOUR-APPLICATION-NAME}}
BASE_URL=https://sandbox.azampay.co.tz
AUTHENTICATOR_URL=https://authenticator-sandbox.azampay.co.tz
CLIENT_ID={{YOUR-CLIENT-ID}}
CLIENT_SECRET={{YOUR-CLIENT-SECRET}}
CLIENT_TOKEN={{YOUR-CLIENT-TOKEN}}

```

## Payment Partners
This function will return the registered partners of provided mechant informations 

```go
    import "github.com/elirehema/azampay"
    func RetrieveRegisteredPartners(){
        clients := azampay.GetPaymentPartners()
        print(clients)
    }

```
This will return array list of partners of [azampay.datas.PaymentPartner struct](https://github.com/elirehema/azampay/blob/master/datas/d_partner.go)

<details>
    <summary> Sample Response example </summary>

```json
[
   {
      "id": "6ebafc56-6d4d-4265-a8d4-c0e1e7806c19",
      "logoUrl": "https://azampay-sarafutest.s3.eu-central-1.amazonaws.com/azampesa.png",
      "partnerName": "Azampesa",
      "provider": 5,
      "vendorName": "your@mechane.email_com",
      "paymentVendorId": "your-payment-vendor-id",
      "paymentPartnerId": "08d997ae-1961-4c32-8b2f-e00f53003b00",
      "paymentAcknowledgmentRoute": "https://your/callback/endpoint",
      "currency": "TZS",
      "status": "1",
      "vendorType": "seller"
   },
   {
      "id": "bbb6121c-b158-4078-aa09-67a584100746",
      "logoUrl": "https://pg-vnext-banners.s3.eu-central-1.amazonaws.com/vnext-images/pgvnext-payment-images/halopesa.svg",
      "partnerName": "HaloPesa",
      "provider": 4,
      "vendorName": "your@mechane.email_com",
      "paymentVendorId": "your-payment-vendor-id",
      "paymentPartnerId": "08d99549-2d49-4694-8320-f60b4e76be6a",
      "paymentAcknowledgmentRoute": "https://your/callback/endpoint",
      "currency": "TZS",
      "status": "1",
      "vendorType": "seller"
   },
   {
      "id": "bf61faec-1421-4e93-bc5b-41090fde3aa9",
      "logoUrl": "https://pg-vnext-banners.s3.eu-central-1.amazonaws.com/vnext-images/pgvnext-payment-images/tigopesa.svg",
      "partnerName": "Tigopesa",
      "provider": 3,
      "vendorName": "your@mechane.email_com",
      "paymentVendorId": "your-payment-vendor-id",
      "paymentPartnerId": "8f2b5341-78c1-4aa2-a8a4-0e1fbe263f1c",
      "paymentAcknowledgmentRoute": "https://your/callback/endpoint",
      "currency": "TZS",
      "status": "1",
      "vendorType": "seller"
   },
   {
      "id": "e21d2941-804f-42e6-a547-66bbbfc7d533",
      "logoUrl": "https://pg-vnext-banners.s3.eu-central-1.amazonaws.com/vnext-images/pgvnext-payment-images/airtel.svg",
      "partnerName": "Airtel",
      "provider": 2,
      "vendorName": "your@mechane.email_com",
      "paymentVendorId": "your-payment-vendor-id",
      "paymentPartnerId": "08d9945a-c9df-4834-876e-04b2df375d8e",
      "paymentAcknowledgmentRoute": "https://your/callback/endpoint",
      "currency": "TZS",
      "status": "1",
      "vendorType": "seller"
   }
]

```


</detail>