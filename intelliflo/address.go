package intelliflo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/models"
)

func (c *credentials) GetAddresses(clientId int) (intelliflomodels.Addresses, error) {
	var addresses intelliflomodels.Addresses
	reqUrl := fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/addresses", clientId)
	req, err := retryablehttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return addresses, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return addresses, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return addresses, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return addresses, fmt.Errorf("error returned by endpoint: %v, with body: %v", resp.StatusCode, string(respBody))
	}
	err = json.Unmarshal(respBody, &addresses)
	if err != nil {
		return addresses, fmt.Errorf("error parsing body: %v", err)
	}
	return addresses, nil
}

func (c *credentials) PostAddress(clientId int, address intelliflomodels.Residence) (intelliflomodels.Residence, error) {
	var newAddress intelliflomodels.Residence
	reqUrl := fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/addresses", clientId)
	reqBody, err := json.Marshal(address)
	if err != nil {
		return newAddress, fmt.Errorf("error marshalling address: %v", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return newAddress, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return newAddress, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return newAddress, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return newAddress, fmt.Errorf("error returned by endpoint: %v, with body: %v", resp.StatusCode, string(respBody))
	}
	err = json.Unmarshal(respBody, &newAddress)
	if err != nil {
		return newAddress, fmt.Errorf("error parsing body: %v", err)
	}
	return newAddress, nil
}

func (c *credentials) PutAddress(clientId int, addressId int, address intelliflomodels.Residence) (intelliflomodels.Residence, error) {
	var updatedAddress intelliflomodels.Residence
	reqUrl := fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/addresses/%d", clientId, addressId)
	reqBody, err := json.Marshal(address)
	if err != nil {
		return updatedAddress, fmt.Errorf("error marshalling address: %v", err)
	}
	req, err := retryablehttp.NewRequest("PUT", reqUrl, reqBody)
	if err != nil {
		return updatedAddress, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return updatedAddress, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return updatedAddress, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return updatedAddress, fmt.Errorf("error returned by endpoint: %v, with body: %v", resp.StatusCode, string(respBody))
	}
	err = json.Unmarshal(respBody, &updatedAddress)
	if err != nil {
		return updatedAddress, fmt.Errorf("error parsing body: %v", err)
	}
	return updatedAddress, nil
}
