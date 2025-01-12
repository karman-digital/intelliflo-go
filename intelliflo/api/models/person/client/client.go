package clientmodels

import (
	personmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
)

type ServiceStatus struct {
	Name      string `json:"name,omitempty"`
	StartedOn string `json:"startedOn,omitempty"`
}

type Campaign struct {
	Type      string `json:"type,omitempty"`
	Source    string `json:"source,omitempty"`
	Reference string `json:"reference,omitempty"`
}

type Client struct {
	ID                     int                       `json:"id,omitempty"`
	Href                   string                    `json:"href,omitempty"`
	Name                   string                    `json:"name,omitempty"`
	CreatedAt              string                    `json:"createdAt,omitempty"`
	Campaign               *Campaign                 `json:"campaign,omitempty"`
	Category               string                    `json:"category,omitempty"`
	MigrationReference     string                    `json:"migrationReference,omitempty"`
	ExternalReference      string                    `json:"externalReference,omitempty"`
	SecondaryReference     string                    `json:"secondaryReference,omitempty"`
	OriginalAdviser        *sharedmodels.IOSubObject `json:"originalAdviser,omitempty"`
	CurrentAdviser         *sharedmodels.IOSubObject `json:"currentAdviser,omitempty"`
	Type                   string                    `json:"type,omitempty"`
	PartyType              string                    `json:"partyType,omitempty"`
	Person                 *personmodels.Person      `json:"person,omitempty"`
	Corporate              *personmodels.Corporate   `json:"corporate,omitempty"`
	Trust                  *personmodels.Trust       `json:"trust,omitempty"`
	AddressesHref          string                    `json:"addresses_href,omitempty"`
	ContactDetailsHref     string                    `json:"contactdetails_href,omitempty"`
	PlansHref              string                    `json:"plans_href,omitempty"`
	RelationshipsHref      string                    `json:"relationships_href,omitempty"`
	ServiceCasesHref       string                    `json:"servicecases_href,omitempty"`
	DependantsHref         string                    `json:"dependants_href,omitempty"`
	IsHeadOfFamilyGroup    bool                      `json:"isHeadOfFamilyGroup,omitempty"`
	ServicingAdministrator *sharedmodels.IOSubObject `json:"servicingAdministrator,omitempty"`
	Paraplanner            *sharedmodels.IOSubObject `json:"paraplanner,omitempty"`
	Tags                   []string                  `json:"tags,omitempty"`
	TaxReferenceNumber     string                    `json:"taxReferenceNumber,omitempty"`
	User                   *sharedmodels.IOSubObject `json:"user,omitempty"`
	ServiceStatus          *ServiceStatus            `json:"serviceStatus,omitempty"`
	Group                  *sharedmodels.IOSubObject `json:"group,omitempty"`
	Household              *sharedmodels.IOSubObject `json:"household,omitempty"`
}

type Clients struct {
	Href      string   `json:"href"`
	FirstHref string   `json:"first_href"`
	LastHref  string   `json:"last_href"`
	NextHref  string   `json:"next_href"`
	PrevHref  string   `json:"prev_href"`
	Items     []Client `json:"items"`
	Count     int      `json:"count"`
}
