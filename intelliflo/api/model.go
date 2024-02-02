package intelliflo

import (
	"time"

	"github.com/hashicorp/go-retryablehttp"
	apptypes "github.com/karman-digital/integrations/types"
	intelliflomodels "github.com/karman-digital/intelliflo/models"
)

type credentials struct {
	client      *retryablehttp.Client
	accessToken apptypes.AccessToken
	expiresAt   apptypes.ExpiresAt
	apiKey      apptypes.APIKey
	tenantId    apptypes.TenantId
}

type IntellfloAPI interface {
	RetrieveAPIKey() apptypes.APIKey
	RetrieveAccessToken() apptypes.AccessToken
	RetrieveExpiresAt() apptypes.ExpiresAt
	RetrieveTenantId() apptypes.TenantId
	SetAccessToken(accessToken string) error
	SetAPIKey(apiKey apptypes.APIKey)
	SetExpiresAt(expiresAt time.Time) error
	SetTenantId(tenantId apptypes.TenantId)
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
}
