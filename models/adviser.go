package intelliflomodels

type Adviser struct {
	ID                 int32  `json:"id"`
	Href               string `json:"href"`
	Name               string `json:"name"`
	AuthorisedOn       string `json:"authorisedOn"`
	FcaRefNo           string `json:"fcaRefNo"`
	Person             Person `json:"person"`
	JoinedFirmOn       string `json:"joinedFirmOn"`
	LeftFirmOn         string `json:"leftFirmOn"`
	MigrationRef       string `json:"migrationRef"`
	ExternalRef1       string `json:"externalRef1"`
	ExternalRef2       string `json:"externalRef2"`
	AddressesHref      string `json:"addresses_href"`
	ContactDetailsHref string `json:"contactDetails_href"`
	TandCCoach         struct {
		ID   int32  `json:"id"`
		Href string `json:"href"`
	} `json:"tandCCoach"`
	User struct {
		ID   int32  `json:"id"`
		Href string `json:"href"`
	} `json:"user"`
	Group struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
		Href string `json:"href"`
	} `json:"group"`
	CanSellOffPanel bool `json:"canSellOffPanel"`
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
