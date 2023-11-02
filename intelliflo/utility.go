package intelliflo

import (
	"fmt"

	apptypes "github.com/karman-digital/hatch-shared/types"
)

func (c *credentials) RetrieveAccessToken() apptypes.AccessToken {
	return c.accessToken
}

func (c *credentials) RetrieveAPIKey() apptypes.APIKey {
	return c.apiKey
}

func (c *credentials) RetrieveExpiresAt() apptypes.ExpiresAt {
	return c.expiresAt
}

func (c *credentials) SetAccessToken(accessToken apptypes.AccessToken) error {
	if accessToken == "" {
		return fmt.Errorf("access token cannot be empty")
	}
	c.accessToken = accessToken
	return nil
}

func (c *credentials) SetAPIKey(apiKey apptypes.APIKey) error {
	if apiKey == "" {
		return fmt.Errorf("api key cannot be empty")
	}
	c.apiKey = apiKey
	return nil
}

func (c *credentials) SetExpiresAt(expiresAt apptypes.ExpiresAt) error {
	if expiresAt.Time().IsZero() {
		return fmt.Errorf("expires at cannot be empty")
	}
	c.expiresAt = expiresAt
	return nil
}
