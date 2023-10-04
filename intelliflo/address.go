package intelliflo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hatch-models/intelliflomodels"
)

func (c *credentials) GetAddresses(clientId string) (intelliflomodels.Addresses, error) {
	var addresses intelliflomodels.Addresses
	reqUrl := fmt.Sprintf("https://api.gb.intelliflo.net/v2/%s/addresses", clientId)
	req, err := retryablehttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return addresses, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.APIKey}
	req.Header["authorization"] = []string{"Bearer " + c.AccessToken}
	resp, err := c.Client.Do(req)
	if err != nil {
		return addresses, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return addresses, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return addresses, fmt.Errorf("error returned by endpoint: %v, with body: %v", resp.StatusCode, string(respBody))
	}
	err = json.Unmarshal(respBody, &addresses)
	if err != nil {
		return addresses, fmt.Errorf("error parsing body: %v", err)
	}
	return addresses, nil
}
