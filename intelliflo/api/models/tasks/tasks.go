package taskmodels

import sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"

type TasksResponse struct {
	Href      string `json:"href"`
	FirstHref string `json:"first_href"`
	LastHref  string `json:"last_href"`
	NextHref  string `json:"next_href"`
	PrevHref  string `json:"prev_href"`
	Items     []Task `json:"items"`
	Count     int    `json:"count"`
}

type TaskType struct {
	ID       int                      `json:"id"`
	Href     string                   `json:"href,omitempty"`
	Name     string                   `json:"name,omitempty"`
	Category sharedmodels.IOSubObject `json:"category"`
}

type Task struct {
	ID            int                       `json:"id,omitempty"`
	Href          string                    `json:"href,omitempty"`
	Subject       string                    `json:"subject"`
	Description   string                    `json:"description,omitempty"`
	ActivityType  TaskType                  `json:"activityType"`
	Priority      *sharedmodels.IOSubObject `json:"priority,omitempty"`
	Status        string                    `json:"status"`
	ShownInDiary  bool                      `json:"shownInDiary"`
	ShownInPortal bool                      `json:"shownInPortal"`
	Reference     string                    `json:"reference"`
	Completion    *TaskCompletion           `json:"completion,omitempty"`
	RelatedTo     []RelatedEntity           `json:"relatedTo,omitempty"`
	LinkedEntity  *LinkedEntity             `json:"linkedEntity,omitempty"`
	AssignedBy    *TaskUser                 `json:"assignedBy,omitempty"`
	AssignedTo    TaskAssignment            `json:"assignedTo"`
	StartsAt      string                    `json:"startsAt,omitempty"`
	DueAt         string                    `json:"dueAt"`
	Duration      *TaskDuration             `json:"duration,omitempty"`
	Billing       *TaskBilling              `json:"billing,omitempty"`
	Cost          *TaskCost                 `json:"cost,omitempty"`
	Recurrence    *TaskRecurrence           `json:"recurrence,omitempty"`
	CreatedByUser TaskUser                  `json:"createdByUser"`
	CreatedAt     string                    `json:"createdAt,omitempty"`
	UpdatedByUser *TaskUser                 `json:"updatedByUser,omitempty"`
	UpdatedAt     string                    `json:"updatedAt,omitempty"`
	Properties    map[string]string         `json:"properties,omitempty"`
	Workflow      *sharedmodels.IOSubObject `json:"workflow,omitempty"`
}

type TaskCompletion struct {
	Percentage  int                      `json:"percentage"`
	CompletedBy TaskUser                 `json:"completedBy"`
	CompletedAt string                   `json:"completedAt"`
	Outcome     sharedmodels.IOSubObject `json:"outcome"`
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
	User TaskUser                  `json:"user"`
	Role *sharedmodels.IOSubObject `json:"role,omitempty"`
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
