package planmodels

import sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"

type PlanStatusChangePayload struct {
	ID          int                  `json:"id"`
	Status      string               `json:"status"`
	EffectiveOn sharedmodels.IntellifloDateTime `json:"effectiveOn"`
	CreatedAt   sharedmodels.IntellifloDateTime `json:"createdAt"`
	Plan        PlanStatusChangePlan `json:"plan"`
}

type PlanStatusChangePlan struct {
	PolicyNumber string `json:"policyNumber"`
	ID           int    `json:"id"`
	Href         string `json:"href"`
}
