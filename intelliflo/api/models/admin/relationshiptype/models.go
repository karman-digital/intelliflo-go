package relationshiptypemodels

type RelationshipTypeResponse struct {
	Href      string `json:"href"`
	FirstHref string `json:"first_href"`
	LastHref  string `json:"last_href"`
	NextHref  string `json:"next_href"`
	PrevHref  string `json:"prev_href"`
	Items     []Item `json:"items"`
	Count     int    `json:"count"`
}

type Item struct {
	Name               string   `json:"name"`
	CorrespondingTypes []string `json:"correspondingTypes"`
	AppliesTo          []string `json:"appliesTo"`
}
