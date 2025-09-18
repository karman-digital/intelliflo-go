package intellifloapp

import (
	"github.com/karman-digital/intelliflo-go/intelliflo/api/activities"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/advisers"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/clients"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/opportunities"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/users"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/webhooks"
	"github.com/karman-digital/intelliflo-go/intelliflo/interfaces"
)

type TenantIntelliflo struct {
	*credentials.TenantCredentials
	ApiClient
	sharedmodels.TenantId
}

type ApiClient struct {
	interfaces.Auth
	Clients       clients.Client
	Advisers      advisers.Adviser
	Users         users.User
	Webhooks      webhooks.Webhook
	Activities    activities.Activity
	Opportunities opportunities.Opportunity
}
