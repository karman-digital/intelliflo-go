package intelliflo

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hatch-models/intelliflomodels"
	apptypes "github.com/karman-digital/hatch-shared/types"
)

type credentials struct {
	client      *retryablehttp.Client
	accessToken apptypes.AccessToken
	expiresAt   apptypes.ExpiresAt
	apiKey      apptypes.APIKey
}

type IntellfloAPI interface {
	RetrieveAPIKey() apptypes.APIKey
	RetrieveAccessToken() apptypes.AccessToken
	RetrieveExpiresAt() apptypes.ExpiresAt
	SetAccessToken(accessToken apptypes.AccessToken) error
	SetAPIKey(apiKey apptypes.APIKey) error
	SetExpiresAt(expiresAt apptypes.ExpiresAt) error
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
