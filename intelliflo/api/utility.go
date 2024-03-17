package intelliflo

import (
	"fmt"
	"time"

	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
)

func (c *credentials) RetrieveAccessToken() intelliflomodels.AccessToken {
	return c.accessToken
}

func (c *credentials) RetrieveAPIKey() intelliflomodels.APIKey {
	return c.apiKey
}

func (c *credentials) RetrieveExpiresAt() intelliflomodels.ExpiresAt {
	return c.expiresAt
}

func (c *credentials) RetrieveTenantId() intelliflomodels.TenantId {
	return c.tenantId
}

func (c *credentials) SetAccessToken(accessToken string) error {
	if accessToken == "" {
		return fmt.Errorf("access token cannot be empty")
	}
	c.accessToken = intelliflomodels.AccessToken(accessToken)
	return nil
}

func (c *credentials) SetAPIKey(apiKey intelliflomodels.APIKey) {
	c.apiKey = apiKey
}

func (c *credentials) SetExpiresAt(expiresAt time.Time) error {
	if expiresAt.IsZero() {
		return fmt.Errorf("expires at cannot be empty")
	}
	c.expiresAt = intelliflomodels.ExpiresAt(expiresAt)
	return nil
}

func (c *credentials) SetTenantId(tenantId intelliflomodels.TenantId) {
	c.tenantId = tenantId
}
