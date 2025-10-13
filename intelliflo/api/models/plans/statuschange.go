package planmodels

import "time"

type PlanStatusChangePayload struct {
	ID          int                  `json:"id"`
	Status      string               `json:"status"`
	EffectiveOn time.Time            `json:"effectiveOn"`
	CreatedAt   time.Time            `json:"createdAt"`
	Plan        PlanStatusChangePlan `json:"plan"`
}

type PlanStatusChangePlan struct {
	PolicyNumber string `json:"policyNumber"`
	ID           int    `json:"id"`
	Href         string `json:"href"`
}
