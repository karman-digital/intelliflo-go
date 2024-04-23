package webhooks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
)

func (w *WebhookService) GetWebhook(id int) (intelliflomodels.Webhook, error) {
	var webhook intelliflomodels.Webhook
	req, err := retryablehttp.NewRequest("GET", fmt.Sprintf("https://api.gb.intelliflo.net/v2/hub/webhooks/%d", id), nil)
	if err != nil {
		return webhook, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{w.ApiKey().String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", w.AccessToken())}
	resp, err := w.Client().Do(req)
	if err != nil {
		return webhook, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return webhook, fmt.Errorf("error returned by endpoint: %v", resp.StatusCode)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return webhook, fmt.Errorf("error reading body: %v", err)
	}
	err = json.Unmarshal(respBody, &webhook)
	if err != nil {
		return webhook, fmt.Errorf("error parsing body: %v", err)
	}
	return webhook, nil
}

func (w *WebhookService) PostWebhookSubscription(postBody intelliflomodels.WebhookSubscriptionRequest) (intelliflomodels.Webhook, error) {
	var responseWebhook intelliflomodels.Webhook
	reqBody, err := json.Marshal(postBody)
	if err != nil {
		return responseWebhook, fmt.Errorf("error converting to body: %v", err)
	}
	req, err := retryablehttp.NewRequest("POST", "https://api.gb.intelliflo.net/v2/hub/webhooks", reqBody)
	if err != nil {
		return responseWebhook, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{w.ApiKey().String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", w.AccessToken())}
	resp, err := w.Client().Do(req)
	if err != nil {
		return responseWebhook, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusAccepted {
		return responseWebhook, fmt.Errorf("error returned by endpoint, status code: %d", resp.StatusCode)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseWebhook, fmt.Errorf("error reading body: %v", err)
	}
	err = json.Unmarshal(respBody, &responseWebhook)
	if err != nil {
		return responseWebhook, fmt.Errorf("error parsing body: %v", err)
	}
	return responseWebhook, nil
}

func (w *WebhookService) GetActiveWebhooks() (intelliflomodels.Webhooks, error) {
	var webhooks intelliflomodels.Webhooks
	req, err := retryablehttp.NewRequest("GET", "https://api.gb.intelliflo.net/v2/hub/webhooks", nil)
	if err != nil {
		return webhooks, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{w.ApiKey().String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", w.AccessToken())}
	resp, err := w.Client().Do(req)
	if err != nil {
		return webhooks, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return webhooks, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return webhooks, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &webhooks)
	if err != nil {
		return webhooks, fmt.Errorf("error parsing body: %v", err)
	}
	return webhooks, nil
}
