package intelliflo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/models"
)

func (c *credentials) PostRelationship(cliendId int, postBody intelliflomodels.Relationship) (intelliflomodels.Relationship, error) {
	var relationship intelliflomodels.Relationship
	body, err := json.Marshal(postBody)
	if err != nil {
		return relationship, fmt.Errorf("error marshalling postBody: %v", err)
	}
	bodyReader := bytes.NewReader(body)
	req, err := retryablehttp.NewRequest("POST", fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/relationships", cliendId), bodyReader)
	if err != nil {
		return relationship, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return relationship, fmt.Errorf("error making post request: %v", err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return relationship, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return relationship, fmt.Errorf("error when creating client: %s", string(respBody))
	}
	defer resp.Body.Close()
	err = json.Unmarshal(respBody, &relationship)
	if err != nil {
		return relationship, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return relationship, nil
}
