package activitiesmodels

type ActivityOutcomeResponse struct {
	Href      string            `json:"href"`
	FirstHref string            `json:"first_href"`
	LastHref  string            `json:"last_href"`
	NextHref  string            `json:"next_href"`
	PrevHref  string            `json:"prev_href"`
	Items     []ActivityOutcome `json:"items"`
	Count     int               `json:"count"`
}

type ActivityOutcome struct {
	ID           int                 `json:"id,omitempty"`
	Href         string              `json:"href,omitempty"`
	Name         string              `json:"name"`
	ActivityType ActivityOutcomeType `json:"activityType"`
	Group        ActivityGroup       `json:"group"`
	IsDeletable  bool                `json:"isDeletable,omitempty"`
	IsHidden     bool                `json:"isHidden,omitempty"`
	IsArchived   bool                `json:"isArchived"`
}

type ActivityOutcomeType struct {
	ID int `json:"id"`
}
