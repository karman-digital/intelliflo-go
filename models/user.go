package intelliflo_models

type User struct {
	ID              int    `json:"id"`
	Href            string `json:"href"`
	UserName        string `json:"userName"`
	Email           string `json:"email"`
	Subject         string `json:"subject"`
	ApplicationType string `json:"applicationType"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	TimeZone        string `json:"timeZone"`
	Status          string `json:"status"`
	References      struct {
		ExternalReference  string `json:"externalReference"`
		MigrationReference string `json:"migrationReference"`
	} `json:"references"`
	Tenant struct {
		ID   int    `json:"id"`
		Href string `json:"href"`
		Name string `json:"name"`
	} `json:"tenant"`
	Group struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Href string `json:"href"`
	} `json:"group"`
	LastLoginAt string `json:"lastLoginAt"`
}

type Users struct {
	Href      string `json:"href"`
	FirstHref string `json:"first_href"`
	LastHref  string `json:"last_href"`
	NextHref  string `json:"next_href"`
	PrevHref  string `json:"prev_href"`
	Items     []User `json:"items"`
	Count     int32  `json:"count"`
}
