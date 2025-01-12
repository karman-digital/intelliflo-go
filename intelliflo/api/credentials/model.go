package credentials

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models"
)

type TenantCredentials struct {
	client       *retryablehttp.Client
	accessToken  intelliflomodels.AccessToken
	expiresAt    intelliflomodels.ExpiresAt
	apiKey       intelliflomodels.APIKey
	clientSecret intelliflomodels.ClientSecret
	cliendId     intelliflomodels.ClientId
}

type Credentials interface {
	Client() *retryablehttp.Client
	AccessToken() *intelliflomodels.AccessToken
	ApiKey() intelliflomodels.APIKey
	ClientSecret() intelliflomodels.ClientSecret
	ClientId() intelliflomodels.ClientId
	SendRequest(method, path string, body []byte, opts ...intelliflomodels.GetOptions) (*http.Response, error)
}
