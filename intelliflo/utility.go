package intelliflo

import "time"

func (c *Credentials) RetrieveAccessToken() string {
	return c.AccessToken
}

func (c *Credentials) RetrieveAPIKey() string {
	return c.APIKey
}

func (c *Credentials) RetrieveExpiresAt() time.Time {
	return c.ExpiresAt
}

func (c *Credentials) SetAccessToken(accessToken string) {
	c.AccessToken = accessToken
}

func (c *Credentials) SetAPIKey(apiKey string) {
	c.APIKey = apiKey
}

func (c *Credentials) SetExpiresAt(expiresAt time.Time) {
	c.ExpiresAt = expiresAt
}
