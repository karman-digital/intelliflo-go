package credentials

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
)

func (t TenantCredentials) AccessToken() intelliflomodels.AccessToken {
	return t.accessToken
}

func (t TenantCredentials) ExpiresAt() intelliflomodels.ExpiresAt {
	return t.expiresAt
}

func (t TenantCredentials) ApiKey() intelliflomodels.APIKey {
	return t.apiKey
}

func (t *TenantCredentials) SetAccessToken(token string) {
	t.accessToken = intelliflomodels.AccessToken(token)
}

func (t *TenantCredentials) SetExpiresAt(expiresAt time.Time) {
	t.expiresAt = intelliflomodels.ExpiresAt(expiresAt)
}

func (t *TenantCredentials) SetApiKey(apiKey string) {
	t.apiKey = intelliflomodels.APIKey(apiKey)
}

func (t *TenantCredentials) SetClient(client *retryablehttp.Client) {
	t.client = client
}

func (t *TenantCredentials) SetClientSecret(clientSecret string) {
	t.clientSecret = intelliflomodels.ClientSecret(clientSecret)
}

func (t *TenantCredentials) SetClientId(clientId string) {
	t.cliendId = intelliflomodels.ClientId(clientId)
}

func (t TenantCredentials) Client() *retryablehttp.Client {
	return t.client
}

func (t TenantCredentials) ClientSecret() intelliflomodels.ClientSecret {
	return t.clientSecret
}

func (t TenantCredentials) ClientId() intelliflomodels.ClientId {
	return t.cliendId
}
