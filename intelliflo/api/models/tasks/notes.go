package taskmodels

type TaskNotesResponse struct {
	Href      string     `json:"href"`
	FirstHref string     `json:"first_href"`
	LastHref  string     `json:"last_href"`
	NextHref  string     `json:"next_href"`
	PrevHref  string     `json:"prev_href"`
	Items     []TaskNote `json:"items"`
	Count     int        `json:"count"`
}

type TaskNote struct {
	ID                 int      `json:"id,omitempty"`
	Href               string   `json:"href,omitempty"`
	Notes              string   `json:"notes"`
	ShowOnClientPortal bool     `json:"showOnClientPortal"`
	CreatedByUser      TaskUser `json:"createdByUser"`
	CreatedAt          string   `json:"createdAt"`
	UpdatedByUser      TaskUser `json:"updatedByUser"`
	UpdatedAt          string   `json:"updatedAt"`
}
