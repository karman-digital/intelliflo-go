package credentials

func NewTenantCredentials(apiKey string, clientId string, clientSecret string) TenantCredentials {
	var creds TenantCredentials
	creds.SetApiKey(apiKey)
	creds.SetClientId(clientId)
	creds.SetClientSecret(clientSecret)
	return creds
}
