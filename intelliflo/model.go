package intelliflo

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hatch-models/intelliflomodels"
	systemiomodels "github.com/karman-digital/hatch-systemio/systemio/models"
)

type AccessToken string
type ExpiresAt time.Time

type credentials struct {
	Client      *retryablehttp.Client
	AccessToken AccessToken
	ExpiresAt   ExpiresAt
	APIKey      systemiomodels.APIKey
}

type IntellfloAPI interface {
	RetrieveAPIKey() systemiomodels.APIKey
	RetrieveAccessToken() AccessToken
	RetrieveExpiresAt() ExpiresAt
	SetAccessToken(accessToken AccessToken) error
	SetAPIKey(apiKey systemiomodels.APIKey) error
	SetExpiresAt(expiresAt ExpiresAt) error
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

func (e ExpiresAt) Time() time.Time {
	return time.Time(e)
}
