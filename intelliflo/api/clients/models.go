package clients

import (
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	"github.com/karman-digital/intelliflo-go/intelliflo/interfaces"
)

type Client struct {
	ClientEndpoint       interfaces.Client
	Addresses            interfaces.Address
	ContactDetails       interfaces.ContactDetail
	Plans                interfaces.Plan
	Holdings             interfaces.Holding
	Relationships        interfaces.Relationship
	MarketingPreferences interfaces.MarketingPreference
}

type ClientService struct {
	credentials.Credentials
}
