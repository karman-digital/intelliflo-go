package intelliflo

import (
	"fmt"
	"time"

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

func (c *credentials) SetAccessToken(accessToken string) error {
	if accessToken == "" {
		return fmt.Errorf("access token cannot be empty")
	}
	c.accessToken = apptypes.AccessToken(accessToken)
	return nil
}

func (c *credentials) SetAPIKey(apiKey apptypes.APIKey) {
	c.apiKey = apiKey
}

func (c *credentials) SetExpiresAt(expiresAt time.Time) error {
	if expiresAt.IsZero() {
		return fmt.Errorf("expires at cannot be empty")
	}
	c.expiresAt = apptypes.ExpiresAt(expiresAt)
	return nil
}
