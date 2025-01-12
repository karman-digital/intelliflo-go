package advisers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	advisermodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/adviser"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
)

func (c *AdviserService) GetAdvisersByUserId(userId int) (advisermodels.Advisers, error) {
	var advisers advisermodels.Advisers
	req, err := retryablehttp.NewRequest("GET", "https://api.gb.intelliflo.net/v2/advisers", nil)
	if err != nil {
		return advisers, fmt.Errorf("error creating request: %v", err)
	}
	queryParams := url.Values{}
	queryParams.Add("filter", fmt.Sprintf("userId eq %d", userId))
	req.URL.RawQuery = queryParams.Encode()
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.ApiKey().String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.AccessToken())}
	resp, err := c.Client().Do(req)
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

func (c *AdviserService) GetAdvisers(options ...sharedmodels.GetOptions) (advisermodels.Advisers, error) {
	var advisers advisermodels.Advisers
	req, err := retryablehttp.NewRequest("GET", "https://api.gb.intelliflo.net/v2/advisers", nil)
	if err != nil {
		return advisers, fmt.Errorf("error creating request: %v", err)
	}
	if len(options) > 0 {
		q := req.URL.Query()
		if options[0].Skip != 0 {
			q.Add("skip", fmt.Sprintf("%d", options[0].Skip))
		}
		if options[0].Top != 0 {
			q.Add("top", fmt.Sprintf("%d", options[0].Top))
		} else {
			q.Add("top", "500")
		}
		if options[0].Filter != "" {
			q.Add("filter", options[0].Filter)
		}
		req.URL.RawQuery = q.Encode()
	} else {
		q := req.URL.Query()
		q.Add("top", "500")
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.ApiKey().String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.AccessToken())}
	resp, err := c.Client().Do(req)
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
