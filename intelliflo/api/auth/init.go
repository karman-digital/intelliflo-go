package auth

import "github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"

func NewTenantAuthService(creds *credentials.TenantCredentials) *AuthTenantService {
	return &AuthTenantService{
		creds,
	}
}
