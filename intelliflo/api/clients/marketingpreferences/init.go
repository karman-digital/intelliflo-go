package marketingpreferences

import "github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"

func NewMarketingPreferencesService(creds credentials.Credentials) *MarketingPreferencesService {
	return &MarketingPreferencesService{
		creds,
	}
}
