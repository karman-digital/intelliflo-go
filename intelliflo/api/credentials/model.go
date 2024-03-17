package credentials

import (
	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
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
	AccessToken() intelliflomodels.AccessToken
	ApiKey() intelliflomodels.APIKey
	ClientSecret() intelliflomodels.ClientSecret
	ClientId() intelliflomodels.ClientId
}
