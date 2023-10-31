package models

type Country struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type NationalityCountry struct {
	Name    string `json:"name,omitempty"`
	IsoCode string `json:"isoCode,omitempty"`
}

type TerritorialProfile struct {
	UkResident             bool      `json:"ukResident,omitempty"`
	CountryOfBirth         string    `json:"countryOfBirth,omitempty"`
	UkDomicile             bool      `json:"ukDomicile,omitempty"`
	Expatriate             bool      `json:"expatriate,omitempty"`
	CountryOfResidence     Country   `json:"countryOfResidence,omitempty"`
	CountryOfDomicile      Country   `json:"countryOfDomicile,omitempty"`
	CountryOfOrigin        Country   `json:"countryOfOrigin,omitempty"`
	PlaceOfBirth           string    `json:"placeOfBirth,omitempty"`
	CountriesOfCitizenship []Country `json:"countriesOfCitizenship,omitempty"`
}

type HealthProfile struct {
	IsSmoker             string `json:"isSmoker,omitempty"`
	SmokedInLast12Months bool   `json:"smokedInLast12Months,omitempty"`
	InGoodHealth         bool   `json:"inGoodHealth,omitempty"`
}

type Person struct {
	Salutation               string              `json:"salutation,omitempty"`
	Title                    string              `json:"title,omitempty"`
	FirstName                string              `json:"firstName,omitempty"`
	MiddleName               string              `json:"middleName,omitempty"`
	LastName                 string              `json:"lastName,omitempty"`
	DateOfBirth              string              `json:"dateOfBirth,omitempty"`
	Gender                   string              `json:"gender,omitempty"`
	MaidenName               string              `json:"maidenName,omitempty"`
	NiNumber                 string              `json:"niNumber,omitempty"`
	MaritalStatus            string              `json:"maritalStatus,omitempty"`
	MaritalStatusSince       string              `json:"maritalStatusSince,omitempty"`
	Nationality              string              `json:"nationality,omitempty"`
	NationalityCountry       *NationalityCountry `json:"nationalityCountry,omitempty"`
	IsDeceased               bool                `json:"isDeceased,omitempty"`
	DeceasedOn               string              `json:"deceasedOn,omitempty"`
	TerritorialProfile       *TerritorialProfile `json:"territorialProfile,omitempty"`
	HealthProfile            *HealthProfile      `json:"healthProfile,omitempty"`
	HasWill                  bool                `json:"hasWill,omitempty"`
	IsWillUptoDate           bool                `json:"isWillUptoDate,omitempty"`
	IsPowerOfAttorneyGranted bool                `json:"isPowerOfAttorneyGranted,omitempty"`
	NationalClientIdentifier string              `json:"nationalClientIdentifier,omitempty"`
}

type Campaign struct {
	Type      string `json:"type,omitempty"`
	Source    string `json:"source,omitempty"`
	Reference string `json:"reference,omitempty"`
}

type IOSubObject struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
}

type LegalEntity struct {
	Identifier string `json:"identifier,omitempty"`
	ExpiresOn  string `json:"expiresOn,omitempty"`
}

type Corporate struct {
	Name                        string       `json:"name,omitempty"`
	Type                        string       `json:"type,omitempty"`
	BusinessType                string       `json:"businessType,omitempty"`
	CompanyRegistrationNumber   string       `json:"companyRegistrationNumber,omitempty"`
	EstablishedOrIncorporatedOn string       `json:"establishedOrIncorporatedOn,omitempty"`
	IsVatRegistered             bool         `json:"isVatRegistered,omitempty"`
	VatRegistrationNumber       string       `json:"vatRegistrationNumber,omitempty"`
	LegalEntity                 *LegalEntity `json:"legalEntity,omitempty"`
	BusinessRegistrationNumber  string       `json:"businessRegistrationNumber,omitempty"`
}

type Trust struct {
	Name                        string       `json:"name,omitempty"`
	Type                        string       `json:"type,omitempty"`
	EstablishedOrIncorporatedOn string       `json:"establishedOrIncorporatedOn,omitempty"`
	RegistrationNumber          string       `json:"registrationNumber,omitempty"`
	RegisteredOn                string       `json:"registeredOn,omitempty"`
	LegalEntity                 *LegalEntity `json:"legalEntity,omitempty"`
	Instrument                  string       `json:"instrument,omitempty"`
	BusinessRegistrationNumber  string       `json:"businessRegistrationNumber,omitempty"`
	NatureOfTrust               string       `json:"natureOfTrust,omitempty"`
	VatRegistrationNumber       string       `json:"vatRegistrationNumber,omitempty"`
}

type Client struct {
	ID                     int          `json:"id,omitempty"`
	Href                   string       `json:"href,omitempty"`
	Name                   string       `json:"name,omitempty"`
	CreatedAt              string       `json:"createdAt,omitempty"`
	Campaign               *Campaign    `json:"campaign,omitempty"`
	Category               string       `json:"category,omitempty"`
	MigrationReference     string       `json:"migrationReference,omitempty"`
	ExternalReference      string       `json:"externalReference,omitempty"`
	SecondaryReference     string       `json:"secondaryReference,omitempty"`
	OriginalAdviser        *IOSubObject `json:"originalAdviser,omitempty"`
	CurrentAdviser         *IOSubObject `json:"currentAdviser,omitempty"`
	Type                   string       `json:"type,omitempty"`
	PartyType              string       `json:"partyType,omitempty"`
	Person                 *Person      `json:"person,omitempty"`
	Corporate              *Corporate   `json:"corporate,omitempty"`
	Trust                  *Trust       `json:"trust,omitempty"`
	AddressesHref          string       `json:"addresses_href,omitempty"`
	ContactDetailsHref     string       `json:"contactdetails_href,omitempty"`
	PlansHref              string       `json:"plans_href,omitempty"`
	RelationshipsHref      string       `json:"relationships_href,omitempty"`
	ServiceCasesHref       string       `json:"servicecases_href,omitempty"`
	DependantsHref         string       `json:"dependants_href,omitempty"`
	IsHeadOfFamilyGroup    bool         `json:"isHeadOfFamilyGroup,omitempty"`
	ServicingAdministrator *IOSubObject `json:"servicingAdministrator,omitempty"`
	Paraplanner            *IOSubObject `json:"paraplanner,omitempty"`
	Tags                   []string     `json:"tags,omitempty"`
	TaxReferenceNumber     string       `json:"taxReferenceNumber,omitempty"`
	User                   *IOSubObject `json:"user,omitempty"`
	ServiceStatus          *struct {
		Name      string `json:"name,omitempty"`
		StartedOn string `json:"startedOn,omitempty"`
	} `json:"serviceStatus,omitempty"`
	Group     *IOSubObject `json:"group,omitempty"`
	Household *IOSubObject `json:"household,omitempty"`
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
