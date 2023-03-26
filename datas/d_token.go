package datas

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type TokenRequest struct {
	ApplicationName string `json:"appName"`
	ClientId        string `json:"clientId"`
	ClientSecret    string `json:"clientSecret"`
}

type Token struct {
	Message    string `json:"message"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Data       struct {
		AccessToken string    `json:"accessToken"`
		Expire      time.Time `json:"expire"`
	} `json:"data"`
}

func CreateTokenRequest() TokenRequest {
	godotenv.Load()
	APPNAME := os.Getenv("APP_NAME")
	CLIENTID := os.Getenv("CLIENT_ID")
	CLIENTSECRET := os.Getenv("CLIENT_SECRET")
	return TokenRequest{
		ApplicationName: APPNAME,
		ClientId:        CLIENTID,
		ClientSecret:    CLIENTSECRET,
	}

}
