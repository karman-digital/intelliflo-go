package credentials

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
)

func (t TenantCredentials) AccessToken() *sharedmodels.AccessToken {
	return &t.accessToken
}

func (t TenantCredentials) ExpiresAt() sharedmodels.ExpiresAt {
	return t.expiresAt
}

func (t TenantCredentials) ApiKey() sharedmodels.APIKey {
	return t.apiKey
}

func (t *TenantCredentials) SetAccessToken(token string) {
	t.accessToken = sharedmodels.AccessToken(token)
}

func (t *TenantCredentials) SetExpiresAt(expiresAt time.Time) {
	t.expiresAt = sharedmodels.ExpiresAt(expiresAt)
}

func (t *TenantCredentials) SetApiKey(apiKey string) {
	t.apiKey = sharedmodels.APIKey(apiKey)
}

func (t *TenantCredentials) SetClient(client *retryablehttp.Client) {
	t.client = client
}

func (t *TenantCredentials) SetClientSecret(clientSecret string) {
	t.clientSecret = sharedmodels.ClientSecret(clientSecret)
}

func (t *TenantCredentials) SetClientId(clientId string) {
	t.cliendId = sharedmodels.ClientId(clientId)
}

func (t TenantCredentials) Client() *retryablehttp.Client {
	return t.client
}

func (t TenantCredentials) ClientSecret() sharedmodels.ClientSecret {
	return t.clientSecret
}

func (t TenantCredentials) ClientId() sharedmodels.ClientId {
	return t.cliendId
}
