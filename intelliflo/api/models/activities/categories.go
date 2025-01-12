package activitiesmodels

type ActivityCategoryResponse struct {
	Href      string             `json:"href"`
	FirstHref string             `json:"first_href"`
	LastHref  string             `json:"last_href"`
	NextHref  string             `json:"next_href"`
	PrevHref  string             `json:"prev_href"`
	Items     []ActivityCategory `json:"items"`
	Count     int                `json:"count"`
}

type ActivityCategory struct {
	ID          int    `json:"id,omitempty"`
	Href        string `json:"href,omitempty"`
	Name        string `json:"name"`
	IsArchived  bool   `json:"isArchived,omitempty"`
	IsDeletable bool   `json:"isDeletable,omitempty"`
}
