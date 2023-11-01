package intelliflo

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hatch-models/intelliflomodels"
)

type AccessToken string
type ExpiresAt time.Time
type APIKey string

type credentials struct {
	Client      *retryablehttp.Client
	AccessToken AccessToken
	ExpiresAt   ExpiresAt
	APIKey      APIKey
}

type IntellfloAPI interface {
	RetrieveAPIKey() APIKey
	RetrieveAccessToken() AccessToken
	RetrieveExpiresAt() ExpiresAt
	SetAccessToken(accessToken string) error
	SetAPIKey(apiKey string) error
	SetExpiresAt(expiresAt time.Time) error
	GetUserById(id int) (intelliflomodels.User, error)
	GetUsersByEmail(email string) (intelliflomodels.Users, error)
	GetAdvisersByUserId(userId int) (intelliflomodels.Advisers, error)
	GetAddresses(clientId int) (intelliflomodels.Addresses, error)
	PostAddress(clientId int, address intelliflomodels.Residence) (intelliflomodels.Residence, error)
	PutAddress(clientId int, addressId int, address intelliflomodels.Residence) (intelliflomodels.Residence, error)
	GenerateAccessToken(clientSecret string, clientId string, tenantId int) (intelliflomodels.TokenResponse, error)
	GenerateAccessTokenScopes(clientSecret string, clientId string, tenantId int, scope []string) (intelliflomodels.TokenResponse, error)
	GetContactDetails(clientId int) (intelliflomodels.ContactDetails, error)
	PostContactDetail(clientId int, contactDetail intelliflomodels.ContactDetail) (intelliflomodels.ContactDetail, error)
	PutContactDetail(clientId int, contactDetailId int, contactDetail intelliflomodels.ContactDetail) (intelliflomodels.ContactDetail, error)
	GetClients(options ...intelliflomodels.GetOptions) (intelliflomodels.Clients, error)
	GetPlans(clientId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Plans, error)
	GetHoldings(clientId int, planId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Holdings, error)
}

func (a AccessToken) String() string {
	return string(a)
}

func (a APIKey) String() string {
	return string(a)
}

func (e ExpiresAt) Time() time.Time {
	return time.Time(e)
}
