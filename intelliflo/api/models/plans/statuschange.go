package planmodels

type PlanStatusChangePayload struct {
	ID          int                  `json:"id"`
	Status      string               `json:"status"`
	EffectiveOn string               `json:"effectiveOn"`
	CreatedAt   string               `json:"createdAt"`
	Plan        PlanStatusChangePlan `json:"plan"`
}

type PlanStatusChangePlan struct {
	PolicyNumber string `json:"policyNumber"`
	ID           int    `json:"id"`
	Href         string `json:"href"`
}
