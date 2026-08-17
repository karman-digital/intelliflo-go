package servicecasemodels

import sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"

type ServiceCases struct {
	Href      string        `json:"href"`
	FirstHref string        `json:"first_href"`
	LastHref  string        `json:"last_href"`
	NextHref  string        `json:"next_href"`
	PrevHref  string        `json:"prev_href"`
	Items     []ServiceCase `json:"items"`
	Count     int           `json:"count"`
}

type ServiceCase struct {
	ID          int                        `json:"id"`
	Href        string                     `json:"href"`
	Reference   string                     `json:"reference"`
	Plans       []sharedmodels.IOSubObject `json:"plans"`
	Opportunity sharedmodels.IOSubObject   `json:"opportunity"`
}
