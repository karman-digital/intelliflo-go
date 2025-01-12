package marketingpreferences

import (
	"encoding/json"
	"fmt"
	"net/http"

	marketingpreferencemodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/marketingpreferences"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (c *MarketingPreferencesService) GetMarketingPreference(clientId int) (marketingpreferencemodels.Preferences, error) {
	var prefs marketingpreferencemodels.Preferences
	resp, err := c.SendRequest("GET", fmt.Sprintf("clients/%d/marketing_preferences", clientId), nil)
	if err != nil {
		return prefs, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return prefs, fmt.Errorf("error handling response code: %v", err)
	}
	err = json.Unmarshal(respBody, &prefs)
	if err != nil {
		return prefs, fmt.Errorf("error parsing body: %v", err)
	}
	return prefs, nil
}

func (c *MarketingPreferencesService) PutMarketingPreference(clientId int, body marketingpreferencemodels.Preferences) (marketingpreferencemodels.Preferences, error) {
	var prefs marketingpreferencemodels.Preferences
	reqBody, err := json.Marshal(body)
	if err != nil {
		return prefs, fmt.Errorf("error converting to body: %v", err)
	}
	resp, err := c.SendRequest("PUT", fmt.Sprintf("clients/%d/marketing_preferences", clientId), reqBody)
	if err != nil {
		return prefs, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return prefs, fmt.Errorf("error handling response code: %v", err)
	}
	err = json.Unmarshal(respBody, &prefs)
	if err != nil {
		return prefs, fmt.Errorf("error parsing body: %v", err)
	}
	return prefs, nil
}
