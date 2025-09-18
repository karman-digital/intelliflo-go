package opportunitiesmodels

import (
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
)

type Opportunity struct {
	ID                     int                      `json:"id"`
	Href                   string                   `json:"href"`
	Status                 OpportunityStatus        `json:"status"`
	Type                   OpportunityType          `json:"type"`
	Prospects              []sharedmodels.Prospect  `json:"prospects"`
	Adviser                sharedmodels.IOSubObject `json:"adviser"`
	PotentialInitialIncome sharedmodels.Currency    `json:"potentialInitialIncome"`
	PotentialOngoingIncome sharedmodels.Currency    `json:"potentialOngoingIncome"`
	AssetValue             sharedmodels.Currency    `json:"assetValue"`
	Proposition            Proposition              `json:"proposition"`
	Campaign               Campaign                 `json:"campaign"`
	Reference              string                   `json:"reference"`
	ServiceCase            ServiceCase              `json:"serviceCase"`
	CreatedOn              string                   `json:"createdOn"`
	CreatedBy              sharedmodels.IOSubObject `json:"createdBy"`
	TargetClosureOn        string                   `json:"targetClosureOn"`
	IsClosed               bool                     `json:"isClosed"`
	ClosedDate             string                   `json:"closedDate"`
}

type OpportunityCreateRequest struct {
	Status                 sharedmodels.IOSubObject       `json:"status"`
	Type                   sharedmodels.IOSubObject       `json:"type"`
	Prospects              []OpportunityProspectRequest   `json:"prospects,omitempty"`
	Adviser                *sharedmodels.IOSubObject      `json:"adviser,omitempty"`
	PotentialInitialIncome *sharedmodels.Currency         `json:"potentialInitialIncome,omitempty"`
	PotentialOngoingIncome *sharedmodels.Currency         `json:"potentialOngoingIncome,omitempty"`
	AssetValue             *sharedmodels.Currency         `json:"assetValue,omitempty"`
	Proposition            *sharedmodels.IOSubObject      `json:"proposition,omitempty"`
	Campaign               *OpportunityCampaignRequest    `json:"campaign,omitempty"`
	Reference              string                         `json:"reference,omitempty"`
	ServiceCase            *OpportunityServiceCaseRequest `json:"serviceCase,omitempty"`
	CreatedOn              string                         `json:"createdOn,omitempty"`
	CreatedBy              *sharedmodels.IOSubObject      `json:"createdBy,omitempty"`
	TargetClosureOn        string                         `json:"targetClosureOn,omitempty"`
	IsClosed               *bool                          `json:"isClosed,omitempty"`
	ClosedDate             string                         `json:"closedDate,omitempty"`
}

type OpportunityProspectRequest struct {
	ID   *int   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type OpportunityCampaignRequest struct {
	Type    string                             `json:"type,omitempty"`
	Name    string                             `json:"name,omitempty"`
	Channel *OpportunityCampaignChannelRequest `json:"channel,omitempty"`
}

type OpportunityCampaignChannelRequest struct {
	Name string `json:"name,omitempty"`
}

type OpportunityServiceCaseRequest struct {
	Reference string `json:"reference,omitempty"`
}

type OpportunityStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OpportunityType struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}

type Proposition struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Campaign struct {
	ID      int             `json:"id"`
	Type    string          `json:"type"`
	Name    string          `json:"name"`
	Channel CampaignChannel `json:"channel"`
}

type CampaignChannel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ServiceCase struct {
	ID        int    `json:"id"`
	Reference string `json:"reference"`
}

type OpportunityCampaignOptionsResponse struct {
	Items []OpportunityCampaignOption `json:"items"`
	Count int                         `json:"count"`
}

type OpportunityCampaignOption struct {
	ID         int                                 `json:"id"`
	Name       string                              `json:"name"`
	Tenant     sharedmodels.IOSubObject            `json:"tenant"`
	Properties OpportunityCampaignOptionProperties `json:"properties"`
}

type OpportunityCampaignOptionProperties struct {
	IsArchived       bool   `json:"isArchived"`
	IsOrganisational bool   `json:"isOrganisational"`
	CampaignType     string `json:"campaignType"`
	Group            string `json:"group"`
}

type OpportunityCampaignTypeOptionsResponse struct {
	Items []OpportunityCampaignTypeOption `json:"items"`
	Count int                             `json:"count"`
}

type OpportunityCampaignTypeOption struct {
	ID         int                                     `json:"id"`
	Name       string                                  `json:"name"`
	Tenant     sharedmodels.IOSubObject                `json:"tenant"`
	Properties OpportunityCampaignTypeOptionProperties `json:"properties"`
}

type OpportunityCampaignTypeOptionProperties struct {
	IsArchived bool `json:"isArchived"`
}

type OpportunityStatusOptionsResponse struct {
	Items []OpportunityStatusOption `json:"items"`
	Count int                       `json:"count"`
}

type OpportunityStatusOption struct {
	ID         int                               `json:"id"`
	Name       string                            `json:"name"`
	Tenant     sharedmodels.IOSubObject          `json:"tenant"`
	Properties OpportunityStatusOptionProperties `json:"properties"`
}

type OpportunityStatusOptionProperties struct {
	IsInitialStatus bool   `json:"isInitialStatus"`
	IsArchived      bool   `json:"isArchived"`
	IsAutoClosed    bool   `json:"isAutoClosed"`
	StatusType      string `json:"statusType"`
}

type OpportunityTypeOptionsResponse struct {
	Href      string                  `json:"href"`
	FirstHref string                  `json:"first_href"`
	LastHref  string                  `json:"last_href"`
	NextHref  string                  `json:"next_href"`
	PrevHref  string                  `json:"prev_href"`
	Items     []OpportunityTypeOption `json:"items"`
	Count     int                     `json:"count"`
}

type OpportunityTypeOption struct {
	ID         int                             `json:"id"`
	Name       string                          `json:"name"`
	Tenant     sharedmodels.IOSubObject        `json:"tenant"`
	Properties OpportunityTypeOptionProperties `json:"properties"`
}

type OpportunityTypeOptionProperties struct {
	IsInvestmentDefault bool   `json:"isInvestmentDefault"`
	IsRetirementDefault bool   `json:"isRetirementDefault"`
	IsSystem            bool   `json:"isSystem"`
	IsArchived          bool   `json:"isArchived"`
	ObjectiveType       string `json:"objectiveType"`
	InUse               bool   `json:"inUse"`
}
