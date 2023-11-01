package intelliflo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hatch-models/intelliflomodels"
)

func (c *credentials) GetPlans(clientId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Plans, error) {
	var plans intelliflomodels.Plans
	req, err := retryablehttp.NewRequest("GET", fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/plans", clientId), nil)
	if err != nil {
		return plans, fmt.Errorf("error creating request: %v", err)
	}
	if len(options) > 0 {
		q := req.URL.Query()
		if options[0].Skip != 0 {
			q.Add("skip", fmt.Sprintf("%d", options[0].Skip))
		}
		if options[0].Top != 0 {
			q.Add("top", fmt.Sprintf("%d", options[0].Top))
		} else {
			q.Add("top", "100")
		}
		req.URL.RawQuery = q.Encode()
	} else {
		q := req.URL.Query()
		q.Add("top", "100")
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.APIKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.AccessToken)}
	resp, err := c.Client.Do(req)
	if err != nil {
		return plans, fmt.Errorf("error making get request: %v", err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return plans, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return plans, fmt.Errorf("error when getting plans. error code:%d, body: %s", resp.StatusCode, string(respBody))
	}
	defer resp.Body.Close()
	err = json.Unmarshal(respBody, &plans)
	if err != nil {
		return plans, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return plans, nil
}
