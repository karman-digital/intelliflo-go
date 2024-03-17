package intellifloapp

import (
	"github.com/karman-digital/intelliflo/intelliflo/api/advisers"
	"github.com/karman-digital/intelliflo/intelliflo/api/auth"
	"github.com/karman-digital/intelliflo/intelliflo/api/clients"
	"github.com/karman-digital/intelliflo/intelliflo/api/credentials"
	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
	"github.com/karman-digital/intelliflo/intelliflo/api/users"
)

func InitTenantAuthIntelliflo(tenantId intelliflomodels.TenantId) *TenantIntelliflo {
	return &TenantIntelliflo{
		TenantId: tenantId,
	}
}

func (i *TenantIntelliflo) InitClient(creds credentials.TenantCredentials) {
	i.TenantCredentials = creds
	i.ApiClient = *NewTenantApiClient(creds)
}

func NewTenantApiClient(creds credentials.TenantCredentials) *ApiClient {
	return &ApiClient{
		Auth:     auth.NewTenantAuthService(creds),
		Clients:  *clients.NewClientObject(creds),
		Advisers: *advisers.NewAdviserObject(creds),
		Users:    *users.NewUser(creds),
	}
}
