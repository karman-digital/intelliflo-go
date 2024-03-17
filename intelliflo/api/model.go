package intelliflo

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
)

type credentials struct {
	client      *retryablehttp.Client
	accessToken intelliflomodels.AccessToken
	expiresAt   intelliflomodels.ExpiresAt
	apiKey      intelliflomodels.APIKey
	tenantId    intelliflomodels.TenantId
}

type IntellifloAPI interface {
	RetrieveAPIKey() intelliflomodels.APIKey
	RetrieveAccessToken() intelliflomodels.AccessToken
	RetrieveExpiresAt() intelliflomodels.ExpiresAt
	RetrieveTenantId() intelliflomodels.TenantId
	SetAccessToken(accessToken string) error
	SetAPIKey(apiKey intelliflomodels.APIKey)
	SetExpiresAt(expiresAt time.Time) error
	SetTenantId(tenantId intelliflomodels.TenantId)
	GetUserById(id int) (intelliflomodels.User, error)
	GetUsersByEmail(email string) (intelliflomodels.Users, error)
	GetAdvisersByUserId(userId int) (intelliflomodels.Advisers, error)
	GetAddresses(clientId int) (intelliflomodels.Addresses, error)
	PostAddress(clientId int, address intelliflomodels.Residence) (intelliflomodels.Residence, error)
	PutAddress(clientId int, addressId int, address intelliflomodels.Residence) (intelliflomodels.Residence, error)
	GetContactDetails(clientId int) (intelliflomodels.ContactDetails, error)
	PostContactDetail(clientId int, contactDetail intelliflomodels.ContactDetail) (intelliflomodels.ContactDetail, error)
	PutContactDetail(clientId int, contactDetailId int, contactDetail intelliflomodels.ContactDetail) (intelliflomodels.ContactDetail, error)
	GetClient(clientId int) (intelliflomodels.Client, error)
	GetClients(options ...intelliflomodels.GetOptions) (intelliflomodels.Clients, error)
	GetPlans(clientId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Plans, error)
	GetHoldings(clientId int, planId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Holdings, error)
	PutClient(clientId int, client intelliflomodels.Client) (intelliflomodels.Client, error)
	PostRelationship(cliendId int, postBody intelliflomodels.Relationship) (intelliflomodels.Relationship, error)
	GetRelationships(clientId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Relationships, error)
	GetClientMarketingPreference(clientId int) (intelliflomodels.Preferences, error)
	PutClientMarketingPreference(clientId int, body intelliflomodels.Preferences) (intelliflomodels.Preferences, error)
	PostClient(clientObj intelliflomodels.Client) (intelliflomodels.Client, error)
	GenerateAccessTokenScopes(clientSecret string, clientId string, scope []string) (intelliflomodels.TenantTokenResponse, error)
}
