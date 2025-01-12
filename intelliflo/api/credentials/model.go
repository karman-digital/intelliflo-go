package credentials

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
)

type TenantCredentials struct {
	client       *retryablehttp.Client
	accessToken  sharedmodels.AccessToken
	expiresAt    sharedmodels.ExpiresAt
	apiKey       sharedmodels.APIKey
	clientSecret sharedmodels.ClientSecret
	cliendId     sharedmodels.ClientId
}

type Credentials interface {
	Client() *retryablehttp.Client
	AccessToken() *sharedmodels.AccessToken
	ApiKey() sharedmodels.APIKey
	ClientSecret() sharedmodels.ClientSecret
	ClientId() sharedmodels.ClientId
	SendRequest(method, path string, body []byte, opts ...sharedmodels.GetOptions) (*http.Response, error)
}
