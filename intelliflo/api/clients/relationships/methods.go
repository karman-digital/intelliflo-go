package relationships

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
)

func (c *RelationshipService) PostRelationship(cliendId int, postBody intelliflomodels.Relationship) (intelliflomodels.Relationship, error) {
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
	req.Header["x-api-key"] = []string{c.ApiKey().String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.AccessToken())}
	resp, err := c.Client().Do(req)
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

func (c *RelationshipService) GetRelationships(clientId int, options ...intelliflomodels.GetOptions) (intelliflomodels.Relationships, error) {
	var relationships intelliflomodels.Relationships
	req, err := retryablehttp.NewRequest("GET", fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/relationships", clientId), nil)
	if err != nil {
		return relationships, fmt.Errorf("error creating request: %v", err)
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
	}
	req.Header["x-api-key"] = []string{c.ApiKey().String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.AccessToken())}
	resp, err := c.Client().Do(req)
	if err != nil {
		return relationships, fmt.Errorf("error making post request: %v", err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return relationships, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return relationships, fmt.Errorf("error when creating client: %s", string(respBody))
	}
	defer resp.Body.Close()
	err = json.Unmarshal(respBody, &relationships)
	if err != nil {
		return relationships, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return relationships, nil
}
