package webhooks

import (
	"github.com/karman-digital/intelliflo/intelliflo/api/credentials"
	"github.com/karman-digital/intelliflo/intelliflo/interfaces"
)

type Webhook struct {
	interfaces.Webhook
}

type WebhookService struct {
	credentials.Credentials
}
