package intelliflomodels

type ContactDetails struct {
	Href      string          `json:"href"`
	FirstHref string          `json:"first_href"`
	LastHref  string          `json:"last_href"`
	NextHref  string          `json:"next_href"`
	PrevHref  string          `json:"prev_href"`
	Items     []ContactDetail `json:"items"`
	Count     int             `json:"count"`
}

type ContactDetail struct {
	ID        int    `json:"id,omitempty"`
	Href      string `json:"href,omitempty"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	IsDefault bool   `json:"isDefault,omitempty"`
	Note      string `json:"note,omitempty"`
}
