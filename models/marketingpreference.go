package intelliflomodels

import "time"

type PrefClient struct {
	ID   int32  `json:"id"`
	Href string `json:"href"`
}

type Preferences struct {
	Href                                   string     `json:"href"`
	Client                                 PrefClient `json:"client"`
	ID                                     int64      `json:"id"`
	AllowCompanyContactByTelephone         bool       `json:"allowCompanyContactByTelephone"`
	AllowCompanyContactByMail              bool       `json:"allowCompanyContactByMail"`
	AllowCompanyContactByEmail             bool       `json:"allowCompanyContactByEmail"`
	AllowCompanyContactBySms               bool       `json:"allowCompanyContactBySms"`
	AllowCompanyContactBySocialMedia       bool       `json:"allowCompanyContactBySocialMedia"`
	AllowCompanyContactByAutomatedCalls    bool       `json:"allowCompanyContactByAutomatedCalls"`
	AllowCompanyContactByPfp               bool       `json:"allowCompanyContactByPfp"`
	AllowThirdPartyContactByTelephone      bool       `json:"allowThirdPartyContactByTelephone"`
	AllowThirdPartyContactByMail           bool       `json:"allowThirdPartyContactByMail"`
	AllowThirdPartyContactByEmail          bool       `json:"allowThirdPartyContactByEmail"`
	AllowThirdPartyContactBySms            bool       `json:"allowThirdPartyContactBySms"`
	AllowThirdPartyContactBySocialMedia    bool       `json:"allowThirdPartyContactBySocialMedia"`
	AllowThirdPartyContactByAutomatedCalls bool       `json:"allowThirdPartyContactByAutomatedCalls"`
	AllowThirdPartyContactByPfp            bool       `json:"allowThirdPartyContactByPfp"`
	CanContactForMarketingPurposes         string     `json:"canContactForMarketingPurposes"`
	ConsentedAt                            time.Time  `json:"consentedAt"`
	DeliveryMethod                         string     `json:"deliveryMethod"`
	AccessibleFormat                       string     `json:"accessibleFormat"`
}
