package activitiesmodels

type ActivityPriorityResponse struct {
	Href      string             `json:"href"`
	FirstHref string             `json:"first_href"`
	LastHref  string             `json:"last_href"`
	NextHref  string             `json:"next_href"`
	PrevHref  string             `json:"prev_href"`
	Items     []ActivityPriority `json:"items"`
	Count     int                `json:"count"`
}
