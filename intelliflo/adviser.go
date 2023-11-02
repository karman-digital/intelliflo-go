package intelliflo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hatch-models/intelliflomodels"
)

func (c *credentials) GetAdvisersByUserId(userId int) (intelliflomodels.Advisers, error) {
	var advisers intelliflomodels.Advisers
	req, err := retryablehttp.NewRequest("GET", "https://api.gb.intelliflo.net/v2/advisers", nil)
	if err != nil {
		return advisers, fmt.Errorf("error creating request: %v", err)
	}
	queryParams := url.Values{}
	queryParams.Add("filter", fmt.Sprintf("userId eq %d", userId))
	req.URL.RawQuery = queryParams.Encode()
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return advisers, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return advisers, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return advisers, fmt.Errorf("error returned by endpoint: %v, with body: %v", resp.StatusCode, string(respBody))
	}
	err = json.Unmarshal(respBody, &advisers)
	if err != nil {
		return advisers, fmt.Errorf("error parsing body: %v", err)
	}
	return advisers, nil
}
