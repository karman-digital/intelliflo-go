package taskmodels

type TasksResponse struct {
	Href      string `json:"href"`
	FirstHref string `json:"first_href"`
	LastHref  string `json:"last_href"`
	NextHref  string `json:"next_href"`
	PrevHref  string `json:"prev_href"`
	Items     []Task `json:"items"`
	Count     int    `json:"count"`
}

type Task struct {
	ID           int    `json:"id,omitempty"`
	Href         string `json:"href,omitempty"`
	Subject      string `json:"subject"`
	Description  string `json:"description,omitempty"`
	ActivityType struct {
		ID       int    `json:"id"`
		Href     string `json:"href,omitempty"`
		Name     string `json:"name,omitempty"`
		Category struct {
			ID   int    `json:"id"`
			Href string `json:"href,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"category"`
	} `json:"activityType"`
	Priority struct {
		ID   int    `json:"id"`
		Href string `json:"href,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"priority"`
	Status        string            `json:"status"`
	ShownInDiary  bool              `json:"shownInDiary"`
	ShownInPortal bool              `json:"shownInPortal"`
	Reference     string            `json:"reference"`
	Completion    TaskCompletion    `json:"completion"`
	RelatedTo     []RelatedEntity   `json:"relatedTo"`
	LinkedEntity  LinkedEntity      `json:"linkedEntity"`
	AssignedBy    TaskUser          `json:"assignedBy"`
	AssignedTo    TaskAssignment    `json:"assignedTo"`
	StartsAt      string            `json:"startsAt"`
	DueAt         string            `json:"dueAt"`
	Duration      TaskDuration      `json:"duration"`
	Billing       TaskBilling       `json:"billing"`
	Cost          TaskCost          `json:"cost"`
	Recurrence    TaskRecurrence    `json:"recurrence"`
	CreatedByUser TaskUser          `json:"createdByUser"`
	CreatedAt     string            `json:"createdAt"`
	UpdatedByUser TaskUser          `json:"updatedByUser"`
	UpdatedAt     string            `json:"updatedAt"`
	Properties    map[string]string `json:"properties"`
	Workflow      struct {
		Name string `json:"name"`
	} `json:"workflow"`
}

type TaskCompletion struct {
	Percentage  int      `json:"percentage"`
	CompletedBy TaskUser `json:"completedBy"`
	CompletedAt string   `json:"completedAt"`
	Outcome     struct {
		ID   int    `json:"id"`
		Href string `json:"href,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"outcome"`
}

type RelatedEntity struct {
	ID   int    `json:"id"`
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}

type LinkedEntity struct {
	ID   int    `json:"id"`
	Href string `json:"href,omitempty"`
	Type string `json:"type"`
}

type TaskUser struct {
	ID   int    `json:"id"`
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}

type TaskAssignment struct {
	User TaskUser `json:"user"`
	Role struct {
		ID   int    `json:"id"`
		Href string `json:"href,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"role"`
}

type TaskDuration struct {
	Estimated TimeSpan `json:"estimated"`
	Actual    TimeSpan `json:"actual"`
	Spent     TimeSpan `json:"spent"`
}

type TimeSpan struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

type TaskBilling struct {
	HourRate Money `json:"hourRate"`
}

type TaskCost struct {
	Estimated Money `json:"estimated"`
	Actual    Money `json:"actual"`
}

type Money struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type TaskRecurrence struct {
	Pattern struct {
		RFCRule string `json:"rfcRule"`
	} `json:"pattern"`
	StartsOn                string `json:"startsOn"`
	EndsOn                  string `json:"endsOn"`
	IsBasedOnCompletionDate bool   `json:"isBasedOnCompletionDate"`
}
