package intellifloapp

import (
	"github.com/karman-digital/intelliflo/intelliflo/api/advisers"
	"github.com/karman-digital/intelliflo/intelliflo/api/auth"
	"github.com/karman-digital/intelliflo/intelliflo/api/clients"
	"github.com/karman-digital/intelliflo/intelliflo/api/credentials"
	"github.com/karman-digital/intelliflo/intelliflo/api/users"
	"github.com/karman-digital/intelliflo/intelliflo/api/webhooks"
)

func InitTenantAuthIntelliflo() *TenantIntelliflo {
	return &TenantIntelliflo{}
}

func (i *TenantIntelliflo) InitClient(creds *credentials.TenantCredentials) {
	i.TenantCredentials = creds
	i.ApiClient = *NewTenantApiClient(creds)
}

func NewTenantApiClient(creds *credentials.TenantCredentials) *ApiClient {
	return &ApiClient{
		Auth:     auth.NewTenantAuthService(creds),
		Clients:  *clients.NewClientObject(creds),
		Advisers: *advisers.NewAdviserObject(creds),
		Users:    *users.NewUser(creds),
		Webhooks: *webhooks.NewWebhook(creds),
	}
}
