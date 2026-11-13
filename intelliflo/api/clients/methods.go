package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	clientmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/client"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (c *ClientService) PostClient(clientObj clientmodels.Client) (clientmodels.Client, error) {
	responseClient := clientmodels.Client{}
	clientBody, err := json.Marshal(clientObj)
	if err != nil {
		return responseClient, fmt.Errorf("error converting to body: %v", err)
	}
	resp, err := c.SendRequest("POST", "clients", clientBody)
	if err != nil {
		return responseClient, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusCreated)
	if err != nil {
		return responseClient, fmt.Errorf("error handling response code: %v, with body: %s", err, string(respBody))
	}
	err = json.Unmarshal(respBody, &responseClient)
	if err != nil {
		return responseClient, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return responseClient, nil
}

func (c *ClientService) GetClients(options ...sharedmodels.GetOptions) (clientmodels.Clients, error) {
	var clients clientmodels.Clients
	resp, err := c.SendRequest("GET", "clients", nil, options...)
	if err != nil {
		return clients, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return clients, fmt.Errorf("error handling response code: %v", err)
	}
	err = json.Unmarshal(respBody, &clients)
	if err != nil {
		return clients, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return clients, nil
}

func (c *ClientService) GetCurrentVulnerabilities(tenantID int, options sharedmodels.GetOptions) (clientmodels.Vulnerabilities, error) {
	var vulnerabilities clientmodels.Vulnerabilities
	if tenantID <= 0 {
		return vulnerabilities, fmt.Errorf("vulnerability lookup requires a positive tenant ID")
	}
	if strings.TrimSpace(options.Filter) == "" {
		return vulnerabilities, fmt.Errorf("vulnerability lookup requires a client filter")
	}
	resp, err := c.SendRequest("GET", fmt.Sprintf("clients/vulnerabilities/current?tenantId=%d", tenantID), nil, options)
	if err != nil {
		return vulnerabilities, fmt.Errorf("send current vulnerabilities request: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return vulnerabilities, fmt.Errorf("handle current vulnerabilities response: %w", err)
	}
	if err := json.Unmarshal(respBody, &vulnerabilities); err != nil {
		return vulnerabilities, fmt.Errorf("decode current vulnerabilities response: %w", err)
	}
	if vulnerabilities.Items == nil {
		return vulnerabilities, fmt.Errorf("current vulnerabilities response is missing items")
	}
	return vulnerabilities, nil
}

func (c *ClientService) GetClient(clientId int) (clientmodels.Client, error) {
	var client clientmodels.Client
	resp, err := c.SendRequest("GET", fmt.Sprintf("clients/%d", clientId), nil)
	if err != nil {
		return client, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return client, fmt.Errorf("error handling response code: %w", err)
	}
	err = json.Unmarshal(respBody, &client)
	if err != nil {
		return client, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return client, nil
}

func (c *ClientService) PutClient(clientId int, client clientmodels.Client) (clientmodels.Client, error) {
	var updatedClient clientmodels.Client
	reqBody, err := json.Marshal(client)
	if err != nil {
		return updatedClient, fmt.Errorf("error converting to body: %v", err)
	}
	resp, err := c.SendRequest("PUT", fmt.Sprintf("/clients/%d", clientId), reqBody)
	if err != nil {
		return updatedClient, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return updatedClient, fmt.Errorf("error handling response code: %v, with body: %s", err, string(respBody))
	}
	err = json.Unmarshal(respBody, &updatedClient)
	if err != nil {
		return updatedClient, fmt.Errorf("error unmarshalling response: %v, with body: %s", err, string(respBody))
	}
	return updatedClient, nil
}
