package auth

import "github.com/karman-digital/intelliflo/intelliflo/api/credentials"

func NewTenantAuthService(creds credentials.TenantCredentials) *AuthTenantService {
	return &AuthTenantService{
		creds,
	}
}
