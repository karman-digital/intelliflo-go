package intellifloapp

import (
	"github.com/karman-digital/intelliflo/intelliflo/api/advisers"
	"github.com/karman-digital/intelliflo/intelliflo/api/clients"
	"github.com/karman-digital/intelliflo/intelliflo/api/credentials"
	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
	"github.com/karman-digital/intelliflo/intelliflo/api/users"
	"github.com/karman-digital/intelliflo/intelliflo/interfaces"
)

type TenantIntelliflo struct {
	credentials.TenantCredentials
	ApiClient
	intelliflomodels.TenantId
}

type ApiClient struct {
	interfaces.Auth
	Clients  clients.Client
	Advisers advisers.Adviser
	Users    users.User
}
