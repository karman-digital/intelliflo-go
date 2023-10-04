package intelliflo

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hatch-models/intelliflomodels"
)

type credentials struct {
	Client      *retryablehttp.Client
	AccessToken string
	ExpiresAt   time.Time
	APIKey      string
}

type IntellfloAPI interface {
	RetrieveAPIKey() string
	RetrieveAccessToken() string
	RetrieveExpiresAt() time.Time
	SetAccessToken(accessToken string)
	SetAPIKey(apiKey string)
	SetExpiresAt(expiresAt time.Time)
	GetUserById(id int) (intelliflomodels.User, error)
	GetUsersByEmail(email string) (intelliflomodels.Users, error)
	GetAdvisersByUserId(userId int) (intelliflomodels.Advisers, error)
	GetAddresses(clientId string) (intelliflomodels.Addresses, error)
}
