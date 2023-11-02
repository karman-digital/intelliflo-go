package intelliflo_models

import "time"

type Address struct {
	Line1      string           `json:"line1,omitempty"`
	Line2      string           `json:"line2,omitempty"`
	Line3      string           `json:"line3,omitempty"`
	Line4      string           `json:"line4,omitempty"`
	Locality   string           `json:"locality,omitempty"`
	PostalCode string           `json:"postalCode,omitempty"`
	Country    *ISOLocationCode `json:"country,omitempty"`
	County     *ISOLocationCode `json:"county,omitempty"`
}

type ISOLocationCode struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type NewBuild struct {
	IsCoveredByNHBCCertificate   *bool  `json:"isCoveredByNHBCCertificate,omitempty"`
	IsCoveredByOtherCertificates *bool  `json:"isCoveredByOtherCertificates,omitempty"`
	OtherCertificateNotes        string `json:"otherCertificateNotes,omitempty"`
	NameOfBuilder                string `json:"nameOfBuilder,omitempty"`
}

type Property struct {
	Type                 string     `json:"type,omitempty"`
	TypeDetail           string     `json:"typeDetail,omitempty"`
	TenureType           string     `json:"tenureType,omitempty"`
	LeaseholdEndsOn      *time.Time `json:"leaseholdEndsOn,omitempty"`
	Status               string     `json:"status,omitempty"`
	Construction         string     `json:"construction,omitempty"`
	ConstructionNotes    string     `json:"constructionNotes,omitempty"`
	RoofConstruction     string     `json:"roofConstruction,omitempty"`
	NumberOfBedrooms     int        `json:"numberOfBedrooms,omitempty"`
	YearBuilt            int        `json:"yearBuilt,omitempty"`
	IsExLocalAuthority   *bool      `json:"isExLocalAuthority,omitempty"`
	NumberOfOutbuildings *int       `json:"numberOfOutbuildings,omitempty"`
	NewBuild             *NewBuild  `json:"newBuild,omitempty"`
}

type Residence struct {
	ID                          int        `json:"id,omitempty"`
	Href                        string     `json:"href,omitempty"`
	ResidencyStatus             string     `json:"residencyStatus,omitempty"`
	Type                        string     `json:"type,omitempty"`
	ResidentFrom                *time.Time `json:"residentFrom,omitempty"`
	ResidentTo                  *time.Time `json:"residentTo,omitempty"`
	Status                      string     `json:"status,omitempty"`
	IsDefault                   bool       `json:"isDefault,omitempty"`
	Address                     *Address   `json:"address,omitempty"`
	IsRegisteredOnElectoralRoll *bool      `json:"isRegisteredOnElectoralRoll,omitempty"`
	IsPotentialMortgage         *bool      `json:"isPotentialMortgage,omitempty"`
	Property                    *Property  `json:"property,omitempty"`
}

type Addresses struct {
	Href      string      `json:"href"`
	FirstHref string      `json:"first_href"`
	LastHref  string      `json:"last_href"`
	NextHref  string      `json:"next_href"`
	PrevHref  string      `json:"prev_href"`
	Items     []Residence `json:"items"`
	Count     int32       `json:"count"`
}
