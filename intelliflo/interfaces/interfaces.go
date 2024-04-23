package interfaces

import intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"

type Auth interface {
	RefreshToken(tenantId intelliflomodels.TenantId, scope []string) error
	ValidateToken() error
}

type Client interface {
	GetClient(clientId int) (intelliflomodels.Client, error)
	GetClients(options ...intelliflomodels.GetOptions) (intelliflomodels.Clients, error)
	PostClient(clientObj intelliflomodels.Client) (intelliflomodels.Client, error)
	PutClient(clientId int, client intelliflomodels.Client) (intelliflomodels.Client, error)
}

type Address interface {
	GetAddresses(int) (intelliflomodels.Addresses, error)
	PostAddress(entityId int, address intelliflomodels.Residence) (intelliflomodels.Residence, error)
	PutAddress(entityId int, addressId int, address intelliflomodels.Residence) (intelliflomodels.Residence, error)
}

type ContactDetail interface {
	GetContactDetails(entityId int) (intelliflomodels.ContactDetails, error)
	PostContactDetail(entityId int, contactDetail intelliflomodels.ContactDetail) (intelliflomodels.ContactDetail, error)
	PutContactDetail(entityId int, contactDetailId int, contactDetail intelliflomodels.ContactDetail) (intelliflomodels.ContactDetail, error)
}

type Plan interface {
	GetPlans(clientId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Plans, error)
}

type Holding interface {
	GetHoldings(clientId int, planId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Holdings, error)
}

type Relationship interface {
	PostRelationship(cliendId int, postBody intelliflomodels.Relationship) (intelliflomodels.Relationship, error)
	GetRelationships(clientId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Relationships, error)
}

type MarketingPreference interface {
	GetMarketingPreference(clientId int) (intelliflomodels.Preferences, error)
	PutMarketingPreference(clientId int, body intelliflomodels.Preferences) (intelliflomodels.Preferences, error)
}

type User interface {
	GetUserById(id int) (intelliflomodels.User, error)
	GetUsersByEmail(email string) (intelliflomodels.Users, error)
}

type Adviser interface {
	GetAdvisersByUserId(userId int) (intelliflomodels.Advisers, error)
	GetAdvisers(options ...intelliflomodels.GetOptions) (intelliflomodels.Advisers, error)
}

type Webhook interface {
	GetActiveWebhooks() (intelliflomodels.Webhooks, error)
	GetWebhook(id int) (intelliflomodels.Webhook, error)
	PostWebhookSubscription(postBody intelliflomodels.WebhookSubscriptionRequest) (intelliflomodels.Webhook, error)
}
