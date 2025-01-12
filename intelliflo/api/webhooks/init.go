package webhooks

import "github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"

func NewWebhook(creds credentials.Credentials) *Webhook {
	return &Webhook{
		NewWebhookService(creds),
	}
}

func NewWebhookService(creds credentials.Credentials) *WebhookService {
	return &WebhookService{
		creds,
	}
}
