package personmodels

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

type LegalEntity struct {
	Identifier string `json:"identifier,omitempty"`
	ExpiresOn  string `json:"expiresOn,omitempty"`
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
	CountryOfResidence     *Country  `json:"countryOfResidence,omitempty"`
	CountryOfDomicile      *Country  `json:"countryOfDomicile,omitempty"`
	CountryOfOrigin        *Country  `json:"countryOfOrigin,omitempty"`
	PlaceOfBirth           string    `json:"placeOfBirth,omitempty"`
	CountriesOfCitizenship []Country `json:"countriesOfCitizenship,omitempty"`
}

type HealthProfile struct {
	IsSmoker             string `json:"isSmoker,omitempty"`
	SmokedInLast12Months bool   `json:"smokedInLast12Months,omitempty"`
	InGoodHealth         bool   `json:"inGoodHealth,omitempty"`
}

type Country struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}
