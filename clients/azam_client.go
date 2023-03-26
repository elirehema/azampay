package clients

import (
	"log"
	"os"
	"time"

	"github.com/elirehema/azampay/datas"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var AzamClient = initHttpClient()

func initHttpClient() *resty.Client {
	client := resty.New()
	godotenv.Load()
	godotenv.Load("key.env")
	BASEURL := os.Getenv("BASE_URL")

	client.SetBaseURL(BASEURL)
	client.SetHeader("Accept", "*/*")
	client.SetHeader("Content-Type", "application/json")
	client.SetTimeout(1 * time.Minute)
	client.SetDebug(true)
	AUTHENTICATOR_URL := os.Getenv("AUTHENTICATOR_URL")

	client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		request := datas.CreateTokenRequest()
		result := &datas.Token{}
		if r.URL != "/AppRegistration/GenerateToken" {
			res, err := c.SetBaseURL(AUTHENTICATOR_URL).R().SetResult(&result).SetBody(request).Post("/AppRegistration/GenerateToken")
			if err != nil {
				print("STATUS CODE {}:", res.StatusCode(), res.Request.URL)
			} else {
				r.SetAuthToken(result.Data.AccessToken)
				c.SetBaseURL(BASEURL)
			}
		}
		return nil // if its success otherwise return error
	})
	client.OnError(func(req *resty.Request, err error) {
		if v, ok := err.(*resty.ResponseError); ok {
		} else {
			log.Println("Logging to custom fileeeeeeee: ", v.Error())
		}

		// Log the error, increment a metric, etc...
	})
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {

		return nil // if its success otherwise return error
	})

	client.SetDebug(true)

	return client
}
