package marketingpreferencemodels

import (
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
)

type PrefClient struct {
	ID   int32  `json:"id,omitempty"`
	Href string `json:"href,omitempty"`
}

type Preferences struct {
	Href                                   string                          `json:"href,omitempty"`
	Client                                 PrefClient                      `json:"client"`
	ID                                     int64                           `json:"id,omitempty"`
	AllowCompanyContactByTelephone         bool                            `json:"allowCompanyContactByTelephone"`
	AllowCompanyContactByMail              bool                            `json:"allowCompanyContactByMail"`
	AllowCompanyContactByEmail             bool                            `json:"allowCompanyContactByEmail"`
	AllowCompanyContactBySms               bool                            `json:"allowCompanyContactBySms"`
	AllowCompanyContactBySocialMedia       bool                            `json:"allowCompanyContactBySocialMedia"`
	AllowCompanyContactByAutomatedCalls    bool                            `json:"allowCompanyContactByAutomatedCalls"`
	AllowCompanyContactByPfp               bool                            `json:"allowCompanyContactByPfp"`
	AllowThirdPartyContactByTelephone      bool                            `json:"allowThirdPartyContactByTelephone"`
	AllowThirdPartyContactByMail           bool                            `json:"allowThirdPartyContactByMail"`
	AllowThirdPartyContactByEmail          bool                            `json:"allowThirdPartyContactByEmail"`
	AllowThirdPartyContactBySms            bool                            `json:"allowThirdPartyContactBySms"`
	AllowThirdPartyContactBySocialMedia    bool                            `json:"allowThirdPartyContactBySocialMedia"`
	AllowThirdPartyContactByAutomatedCalls bool                            `json:"allowThirdPartyContactByAutomatedCalls"`
	AllowThirdPartyContactByPfp            bool                            `json:"allowThirdPartyContactByPfp"`
	CanContactForMarketingPurposes         string                          `json:"canContactForMarketingPurposes"`
	ConsentedAt                            sharedmodels.IntellifloDateTime `json:"consentedAt,omitempty"`
	DeliveryMethod                         string                          `json:"deliveryMethod,omitempty"`
	AccessibleFormat                       string                          `json:"accessibleFormat,omitempty"`
}
