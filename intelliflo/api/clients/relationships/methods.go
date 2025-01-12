package relationships

import (
	"encoding/json"
	"fmt"
	"net/http"

	relationshipsmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/relationships"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (c *RelationshipService) PostRelationship(cliendId int, postBody relationshipsmodels.Relationship) (relationshipsmodels.Relationship, error) {
	var relationship relationshipsmodels.Relationship
	reqBody, err := json.Marshal(postBody)
	if err != nil {
		return relationship, fmt.Errorf("error marshalling postBody: %v", err)
	}
	resp, err := c.SendRequest("POST", fmt.Sprintf("clients/%d/relationships", cliendId), reqBody)
	if err != nil {
		return relationship, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return relationship, fmt.Errorf("error handling response code: %v", err)
	}
	err = json.Unmarshal(respBody, &relationship)
	if err != nil {
		return relationship, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return relationship, nil
}

func (c *RelationshipService) GetRelationships(clientId int, options ...sharedmodels.GetOptions) (relationshipsmodels.Relationships, error) {
	var relationships relationshipsmodels.Relationships
	resp, err := c.SendRequest("GET", fmt.Sprintf("clients/%d/relationships", clientId), nil, options...)
	if err != nil {
		return relationships, fmt.Errorf("error sending request: %v", err)
	}
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return relationships, fmt.Errorf("error handling response code: %v", err)
	}
	err = json.Unmarshal(respBody, &relationships)
	if err != nil {
		return relationships, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return relationships, nil
}
