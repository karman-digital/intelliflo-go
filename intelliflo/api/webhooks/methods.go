package webhooks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	webhookmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/webhooks"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (w *WebhookService) GetWebhook(id int) (webhookmodels.Webhook, error) {
	var webhook webhookmodels.Webhook
	resp, err := w.SendRequest("GET", fmt.Sprintf("hub/webhooks/%d", id), nil)
	if err != nil {
		return webhook, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return webhook, fmt.Errorf("error handling response code: %v", err)
	}
	err = json.Unmarshal(respBody, &webhook)
	if err != nil {
		return webhook, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return webhook, nil
}

func (w *WebhookService) PostWebhookSubscription(postBody webhookmodels.WebhookSubscriptionRequest) (webhookmodels.Webhook, error) {
	var responseWebhook webhookmodels.Webhook
	reqBody, err := json.Marshal(postBody)
	if err != nil {
		return responseWebhook, fmt.Errorf("error converting to body: %v", err)
	}
	fmt.Println(string(reqBody))
	resp, err := w.SendRequest("POST", "hub/webhooks", reqBody)
	if err != nil {
		return responseWebhook, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusAccepted)
	if err != nil {
		body, _ := io.ReadAll(resp.Body)
		return responseWebhook, fmt.Errorf("error handling response code: %v, body: %s", err, string(body))
	}
	err = json.Unmarshal(respBody, &responseWebhook)
	if err != nil {
		return responseWebhook, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return responseWebhook, nil
}

func (w *WebhookService) GetActiveWebhooks() (webhookmodels.Webhooks, error) {
	var webhooks webhookmodels.Webhooks
	resp, err := w.SendRequest("GET", "hub/webhooks", nil)
	if err != nil {
		return webhooks, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return webhooks, fmt.Errorf("error handling response code: %v", err)
	}
	err = json.Unmarshal(respBody, &webhooks)
	if err != nil {
		return webhooks, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return webhooks, nil
}
