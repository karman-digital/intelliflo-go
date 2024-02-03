package intelliflo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
)

func (c *credentials) GetClientMarketingPreference(clientId int) (intelliflomodels.Preferences, error) {
	var prefs intelliflomodels.Preferences
	req, err := retryablehttp.NewRequest("GET", fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/marketing_preferences", clientId), nil)
	if err != nil {
		return prefs, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return prefs, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return prefs, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return prefs, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &prefs)
	if err != nil {
		return prefs, fmt.Errorf("error parsing body: %v", err)
	}
	return prefs, nil
}

func (c *credentials) PutClientMarketingPreference(clientId int, body intelliflomodels.Preferences) (intelliflomodels.Preferences, error) {
	var prefs intelliflomodels.Preferences
	reqBody, err := json.Marshal(body)
	if err != nil {
		return prefs, fmt.Errorf("error converting to body: %v", err)
	}
	req, err := retryablehttp.NewRequest("PUT", fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/marketing_preferences", clientId), reqBody)
	if err != nil {
		return prefs, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return prefs, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return prefs, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return prefs, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &prefs)
	if err != nil {
		return prefs, fmt.Errorf("error parsing body: %v", err)
	}
	return prefs, nil
}
