package webhookmodels

import feesmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/fees"

type FeesWebhookBody struct {
	WebhookBody
	Payload feesmodels.Fee `json:"payload"`
}
