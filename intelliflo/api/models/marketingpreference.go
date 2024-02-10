package intelliflomodels

import "time"

type PrefClient struct {
	ID   int32  `json:"id,omitempty"`
	Href string `json:"href,omitempty"`
}

type Preferences struct {
	Href                                   string     `json:"href,omitempty"`
	Client                                 PrefClient `json:"client"`
	ID                                     int64      `json:"id,omitempty"`
	AllowCompanyContactByTelephone         *bool      `json:"allowCompanyContactByTelephone,omitempty"`
	AllowCompanyContactByMail              *bool      `json:"allowCompanyContactByMail,omitempty"`
	AllowCompanyContactByEmail             *bool      `json:"allowCompanyContactByEmail,omitempty"`
	AllowCompanyContactBySms               *bool      `json:"allowCompanyContactBySms,omitempty"`
	AllowCompanyContactBySocialMedia       *bool      `json:"allowCompanyContactBySocialMedia,omitempty"`
	AllowCompanyContactByAutomatedCalls    *bool      `json:"allowCompanyContactByAutomatedCalls,omitempty"`
	AllowCompanyContactByPfp               *bool      `json:"allowCompanyContactByPfp,omitempty"`
	AllowThirdPartyContactByTelephone      *bool      `json:"allowThirdPartyContactByTelephone,omitempty"`
	AllowThirdPartyContactByMail           *bool      `json:"allowThirdPartyContactByMail,omitempty"`
	AllowThirdPartyContactByEmail          *bool      `json:"allowThirdPartyContactByEmail,omitempty"`
	AllowThirdPartyContactBySms            *bool      `json:"allowThirdPartyContactBySms,omitempty"`
	AllowThirdPartyContactBySocialMedia    *bool      `json:"allowThirdPartyContactBySocialMedia,omitempty"`
	AllowThirdPartyContactByAutomatedCalls *bool      `json:"allowThirdPartyContactByAutomatedCalls,omitempty"`
	AllowThirdPartyContactByPfp            *bool      `json:"allowThirdPartyContactByPfp,omitempty"`
	CanContactForMarketingPurposes         string     `json:"canContactForMarketingPurposes"`
	ConsentedAt                            time.Time  `json:"consentedAt,omitempty"`
	DeliveryMethod                         string     `json:"deliveryMethod,omitempty"`
	AccessibleFormat                       string     `json:"accessibleFormat,omitempty"`
}
