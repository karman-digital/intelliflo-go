package addresses

import (
	"encoding/json"
	"fmt"
	"net/http"

	addressmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/address"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (c *AddressService) GetAddresses(clientId int, opts ...sharedmodels.GetOptions) (addressmodels.Addresses, error) {
	var addresses addressmodels.Addresses
	resp, err := c.SendRequest("GET", fmt.Sprintf("/clients/%d/addresses", clientId), nil, opts...)
	if err != nil {
		return addresses, fmt.Errorf("error making get request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return addresses, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &addresses)
	if err != nil {
		return addresses, fmt.Errorf("error parsing body: %v", err)
	}
	return addresses, nil
}

func (c *AddressService) PostAddress(clientId int, address addressmodels.Residence) (addressmodels.Residence, error) {
	var newAddress addressmodels.Residence
	reqBody, err := json.Marshal(address)
	if err != nil {
		return newAddress, fmt.Errorf("error marshalling address: %v", err)
	}
	resp, err := c.SendRequest("POST", fmt.Sprintf("/v2/clients/%d/addresses", clientId), reqBody)
	if err != nil {
		return newAddress, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusCreated)
	if err != nil {
		return newAddress, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &newAddress)
	if err != nil {
		return newAddress, fmt.Errorf("error parsing body: %v", err)
	}
	return newAddress, nil
}

func (c *AddressService) PutAddress(clientId int, addressId int, address addressmodels.Residence) (addressmodels.Residence, error) {
	var updatedAddress addressmodels.Residence
	reqBody, err := json.Marshal(address)
	if err != nil {
		return updatedAddress, fmt.Errorf("error marshalling address: %v", err)
	}
	resp, err := c.SendRequest("PUT", fmt.Sprintf("/v2/clients/%d/addresses/%d", clientId, addressId), reqBody)
	if err != nil {
		return updatedAddress, fmt.Errorf("error making put request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return updatedAddress, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &updatedAddress)
	if err != nil {
		return updatedAddress, fmt.Errorf("error parsing body: %v", err)
	}
	return updatedAddress, nil
}
