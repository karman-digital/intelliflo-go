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

func (c *credentials) GetUserById(id int) (intelliflomodels.User, error) {
	var user intelliflomodels.User
	req, err := retryablehttp.NewRequest("GET", fmt.Sprintf("https://api.gb.intelliflo.net/v2/users/%d", id), nil)
	if err != nil {
		return user, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return user, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return user, fmt.Errorf("error returned by endpoint: %v", resp.StatusCode)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return user, fmt.Errorf("error reading body: %v", err)
	}
	err = json.Unmarshal(respBody, &user)
	if err != nil {
		return user, fmt.Errorf("error parsing body: %v", err)
	}
	return user, nil
}

func (c *credentials) GetUsersByEmail(email string) (intelliflomodels.Users, error) {
	var users intelliflomodels.Users
	req, err := retryablehttp.NewRequest("GET", "https://api.gb.intelliflo.net/v2/users", nil)
	if err != nil {
		return users, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	queryParams := url.Values{}
	queryParams.Add("filter", fmt.Sprintf("email eq '%s'", email))
	req.URL.RawQuery = queryParams.Encode()
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return users, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return users, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return users, fmt.Errorf("error returned by endpoint: %v, with body: %v", resp.StatusCode, string(respBody))
	}
	err = json.Unmarshal(respBody, &users)
	if err != nil {
		return users, fmt.Errorf("error parsing body: %v", err)
	}
	return users, nil
}
