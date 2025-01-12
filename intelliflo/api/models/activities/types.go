package activitiesmodels

type ActivityTypeResponse struct {
	Href      string         `json:"href"`
	FirstHref string         `json:"first_href"`
	LastHref  string         `json:"last_href"`
	NextHref  string         `json:"next_href"`
	PrevHref  string         `json:"prev_href"`
	Items     []ActivityType `json:"items"`
	Count     int            `json:"count"`
}

type ActivityType struct {
	ID                   int              `json:"id,omitempty"`
	Href                 string           `json:"href,omitempty"`
	Name                 string           `json:"name"`
	Description          string           `json:"description"`
	ActivityEventType    string           `json:"activityEventType"`
	IsClientRelated      bool             `json:"isClientRelated"`
	IsPlanRelated        bool             `json:"isPlanRelated"`
	IsFeeRelated         bool             `json:"isFeeRelated"`
	IsRetainerRelated    bool             `json:"isRetainerRelated"`
	IsOpportunityRelated bool             `json:"isOpportunityRelated"`
	IsAdviserRelated     bool             `json:"isAdviserRelated"`
	IsOutcomeMandatory   bool             `json:"isOutcomeMandatory"`
	IsArchived           bool             `json:"isArchived"`
	IsDeletable          bool             `json:"isDeletable"`
	IsHidden             bool             `json:"isHidden"`
	Category             ActivityCategory `json:"category"`
	Group                ActivityGroup    `json:"group"`
	IncludeSubgroups     bool             `json:"includeSubgroups"`
	Priority             ActivityPriority `json:"priority"`
	TaskBillingRate      int              `json:"taskBillingRate"`
	EstimatedTime        EstimatedTime    `json:"estimatedTime"`
	ActivityTypeEvents   []ActivityEvent  `json:"activityTypeEvents"`
}

type ActivityGroup struct {
	ID   int    `json:"id,omitempty"`
	Href string `json:"href,omitempty"`
	Name string `json:"name"`
}

type ActivityPriority struct {
	ID   int    `json:"id"`
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}

type EstimatedTime struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

type ActivityEvent struct {
	ID   int    `json:"id"`
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}

type ActivityEventTypeResponse struct {
	Href      string          `json:"href"`
	FirstHref string          `json:"first_href"`
	LastHref  string          `json:"last_href"`
	NextHref  string          `json:"next_href"`
	PrevHref  string          `json:"prev_href"`
	Items     []ActivityEvent `json:"items"`
	Count     int             `json:"count"`
}
