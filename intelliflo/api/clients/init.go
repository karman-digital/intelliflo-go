package clients

import (
	"github.com/karman-digital/intelliflo-go/intelliflo/api/clients/addresses"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/clients/contactdetails"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/clients/fees"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/clients/marketingpreferences"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/clients/plans"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/clients/plans/holdings"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/clients/relationships"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/clients/servicecases"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
)

func NewClientObject(creds credentials.Credentials) *Client {
	return &Client{
		ClientEndpoint:       NewClientService(creds),
		Addresses:            addresses.NewAddressService(creds),
		ContactDetails:       contactdetails.NewContactDetailsService(creds),
		Fees:                 fees.NewFeeService(creds),
		Plans:                plans.NewPlansService(creds),
		Holdings:             holdings.NewHoldingService(creds),
		Relationships:        relationships.NewRelationshipService(creds),
		MarketingPreferences: marketingpreferences.NewMarketingPreferencesService(creds),
		ServiceCases:         servicecases.NewServiceCasesService(creds),
	}
}

func NewClientService(creds credentials.Credentials) *ClientService {
	return &ClientService{
		creds,
	}
}
