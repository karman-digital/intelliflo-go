package intelliflo

import "time"

func (c *credentials) RetrieveAccessToken() string {
	return c.AccessToken
}

func (c *credentials) RetrieveAPIKey() string {
	return c.APIKey
}

func (c *credentials) RetrieveExpiresAt() time.Time {
	return c.ExpiresAt
}

func (c *credentials) SetAccessToken(accessToken string) {
	c.AccessToken = accessToken
}

func (c *credentials) SetAPIKey(apiKey string) {
	c.APIKey = apiKey
}

func (c *credentials) SetExpiresAt(expiresAt time.Time) {
	c.ExpiresAt = expiresAt
}
