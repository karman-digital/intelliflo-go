package intelliflo

import (
	"fmt"
	"time"
)

func (c *credentials) RetrieveAccessToken() AccessToken {
	return c.AccessToken
}

func (c *credentials) RetrieveAPIKey() APIKey {
	return c.APIKey
}

func (c *credentials) RetrieveExpiresAt() ExpiresAt {
	return c.ExpiresAt
}

func (c *credentials) SetAccessToken(accessToken string) error {
	if accessToken == "" {
		return fmt.Errorf("access token cannot be empty")
	}
	c.AccessToken = AccessToken(accessToken)
	return nil
}

func (c *credentials) SetAPIKey(apiKey string) error {
	if apiKey == "" {
		return fmt.Errorf("api key cannot be empty")
	}
	c.APIKey = APIKey(apiKey)
	return nil
}

func (c *credentials) SetExpiresAt(expiresAt time.Time) error {
	if expiresAt.IsZero() {
		return fmt.Errorf("expires at cannot be empty")
	}
	c.ExpiresAt = ExpiresAt(expiresAt)
	return nil
}
