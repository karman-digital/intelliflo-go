package intelliflo

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/models"
)

func (c *credentials) GenerateAccessTokenScopes(clientSecret string, clientId string, scope []string) (intelliflomodels.TokenResponse, error) {
	var tokenBody = intelliflomodels.TokenResponse{}
	encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))
	data := url.Values{
		"grant_type": []string{"tenant_client_credentials"},
		"tenant_id":  []string{fmt.Sprint(c.tenantId)},
	}
	scopes := ""
	for _, s := range scope {
		scopes += s + " "
	}
	data.Add("scope", scopes)

	req, err := retryablehttp.NewRequest("POST", "https://identity.intelliflo.com/core/connect/token", strings.NewReader(data.Encode()))
	if err != nil {
		return tokenBody, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header["authorization"] = []string{"Basic " + encoded}
	resp, err := c.client.Do(req)
	if err != nil {
		return tokenBody, fmt.Errorf("error making post request: %s", err)
	}
	defer resp.Body.Close()
	tokenRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return tokenBody, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return tokenBody, fmt.Errorf("error returned by endpoint: %s", tokenRawBody)
	}
	err = json.Unmarshal(tokenRawBody, &tokenBody)
	if err != nil {
		return tokenBody, fmt.Errorf("error parsing body: %s", err)
	}
	return tokenBody, nil
}

func GenerateAccessToken(clientSecret string, clientId string, tenantId int, scope []string) (intelliflomodels.TokenResponse, error) {
	var tokenBody = intelliflomodels.TokenResponse{}
	encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))
	data := url.Values{
		"grant_type": []string{"tenant_client_credentials"},
		"tenant_id":  []string{fmt.Sprint(tenantId)},
	}
	scopes := ""
	for _, s := range scope {
		scopes += s + " "
	}
	data.Add("scope", scopes)

	req, err := retryablehttp.NewRequest("POST", "https://identity.intelliflo.com/core/connect/token", strings.NewReader(data.Encode()))
	if err != nil {
		return tokenBody, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header["authorization"] = []string{"Basic " + encoded}
	client := retryablehttp.NewClient()
	resp, err := client.Do(req)
	if err != nil {
		return tokenBody, fmt.Errorf("error making post request: %s", err)
	}
	defer resp.Body.Close()
	tokenRawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return tokenBody, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return tokenBody, fmt.Errorf("error returned by endpoint: %s", tokenRawBody)
	}
	err = json.Unmarshal(tokenRawBody, &tokenBody)
	if err != nil {
		return tokenBody, fmt.Errorf("error parsing body: %s", err)
	}
	return tokenBody, nil
}
