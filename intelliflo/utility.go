package intelliflo

import (
	"fmt"

	systemiomodels "github.com/karman-digital/hatch-systemio/systemio/models"
)

func (c *credentials) RetrieveAccessToken() AccessToken {
	return c.AccessToken
}

func (c *credentials) RetrieveAPIKey() systemiomodels.APIKey {
	return c.APIKey
}

func (c *credentials) RetrieveExpiresAt() ExpiresAt {
	return c.ExpiresAt
}

func (c *credentials) SetAccessToken(accessToken AccessToken) error {
	if accessToken == "" {
		return fmt.Errorf("access token cannot be empty")
	}
	c.AccessToken = accessToken
	return nil
}

func (c *credentials) SetAPIKey(apiKey systemiomodels.APIKey) error {
	if apiKey == "" {
		return fmt.Errorf("api key cannot be empty")
	}
	c.APIKey = apiKey
	return nil
}

func (c *credentials) SetExpiresAt(expiresAt ExpiresAt) error {
	if expiresAt.Time().IsZero() {
		return fmt.Errorf("expires at cannot be empty")
	}
	c.ExpiresAt = expiresAt
	return nil
}
