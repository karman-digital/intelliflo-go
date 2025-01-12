package webhooks

import (
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	"github.com/karman-digital/intelliflo-go/intelliflo/interfaces"
)

type Webhook struct {
	interfaces.Webhook
}

type WebhookService struct {
	credentials.Credentials
}
