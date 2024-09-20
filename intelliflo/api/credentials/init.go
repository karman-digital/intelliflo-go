package credentials

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func NewTenantCredentials(apiKey string, clientId string, clientSecret string, accessToken string, expiredAt time.Time) TenantCredentials {
	var creds TenantCredentials
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.ErrorHandler = retryablehttp.PassthroughErrorHandler
	creds.SetClient(client)
	creds.SetAccessToken(accessToken)
	creds.SetExpiresAt(expiredAt)
	creds.SetApiKey(apiKey)
	creds.SetClientId(clientId)
	creds.SetClientSecret(clientSecret)
	return creds
}
