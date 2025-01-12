package advisermodels

import personmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person"

type Adviser struct {
	ID                 int32               `json:"id"`
	Href               string              `json:"href"`
	Name               string              `json:"name"`
	AuthorisedOn       string              `json:"authorisedOn"`
	FcaRefNo           string              `json:"fcaRefNo"`
	Person             personmodels.Person `json:"person"`
	JoinedFirmOn       string              `json:"joinedFirmOn"`
	LeftFirmOn         string              `json:"leftFirmOn"`
	MigrationRef       string              `json:"migrationRef"`
	ExternalRef1       string              `json:"externalRef1"`
	ExternalRef2       string              `json:"externalRef2"`
	AddressesHref      string              `json:"addresses_href"`
	ContactDetailsHref string              `json:"contactDetails_href"`
	TandCCoach         AdviserUser         `json:"tandCCoach"`
	User               AdviserUser         `json:"user"`
	Group              Group               `json:"group"`
	CanSellOffPanel    bool                `json:"canSellOffPanel"`
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

type AdviserUser struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

type Advisers struct {
	Href      string    `json:"href"`
	FirstHref string    `json:"first_href"`
	LastHref  string    `json:"last_href"`
	NextHref  string    `json:"next_href"`
	PrevHref  string    `json:"prev_href"`
	Items     []Adviser `json:"items"`
	Count     int32     `json:"count"`
}
