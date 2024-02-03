package intelliflo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
)

func (c *credentials) GetHoldings(clientId int, planId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Holdings, error) {
	var holdings intelliflomodels.Holdings
	req, err := retryablehttp.NewRequest("GET", fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/plans/%d/holdings", clientId, planId), nil)
	if err != nil {
		return holdings, fmt.Errorf("error creating request: %v", err)
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
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return holdings, fmt.Errorf("error making get request: %v", err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return holdings, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return holdings, fmt.Errorf("error when getting plans. error code:%d, body: %s", resp.StatusCode, string(respBody))
	}
	defer resp.Body.Close()
	err = json.Unmarshal(respBody, &holdings)
	if err != nil {
		return holdings, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return holdings, nil
}
