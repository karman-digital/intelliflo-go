package marketingpreferences

import "github.com/karman-digital/intelliflo/intelliflo/api/credentials"

func NewMarketingPreferencesService(creds credentials.Credentials) *MarketingPreferencesService {
	return &MarketingPreferencesService{
		creds,
	}
}
