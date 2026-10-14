package contactdetails

import (
	"encoding/json"
	"fmt"
	"net/http"

	contactdetailmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/contactdetails"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (c *ContactDetailsService) GetContactDetails(clientId int, opts ...sharedmodels.GetOptions) (contactdetailmodels.ContactDetails, error) {
	var contactDetails contactdetailmodels.ContactDetails
	resp, err := c.SendRequest("GET", fmt.Sprintf("clients/%d/contactdetails", clientId), nil, opts...)
	if err != nil {
		return contactDetails, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := handleContactDetailResponseCode(resp, http.StatusOK)
	if err != nil {
		return contactDetails, err
	}
	err = json.Unmarshal(respBody, &contactDetails)
	if err != nil {
		return contactDetails, fmt.Errorf("error parsing body: %v", err)
	}
	return contactDetails, nil
}

func (c *ContactDetailsService) PostContactDetail(clientId int, contactDetail contactdetailmodels.ContactDetail) (contactdetailmodels.ContactDetail, error) {
	var newContactDetail contactdetailmodels.ContactDetail
	reqBody, err := json.Marshal(contactDetail)
	if err != nil {
		return newContactDetail, fmt.Errorf("error marshalling contactDetail: %v", err)
	}
	resp, err := c.SendRequest("POST", fmt.Sprintf("clients/%d/contactdetails", clientId), reqBody)
	if err != nil {
		return newContactDetail, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := handleContactDetailResponseCode(resp, http.StatusCreated)
	if err != nil {
		return newContactDetail, err
	}
	err = json.Unmarshal(respBody, &newContactDetail)
	if err != nil {
		return newContactDetail, fmt.Errorf("error parsing body: %v", err)
	}
	return newContactDetail, nil
}

func (c *ContactDetailsService) PutContactDetail(clientId int, contactDetailId int, contactDetail contactdetailmodels.ContactDetail) (contactdetailmodels.ContactDetail, error) {
	var updatedContactDetail contactdetailmodels.ContactDetail
	reqBody, err := json.Marshal(contactDetail)
	if err != nil {
		return updatedContactDetail, fmt.Errorf("error marshalling contactDetail: %v", err)
	}
	resp, err := c.SendRequest("PUT", fmt.Sprintf("clients/%d/contactdetails/%d", clientId, contactDetailId), reqBody)
	if err != nil {
		return updatedContactDetail, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := handleContactDetailResponseCode(resp, http.StatusOK)
	if err != nil {
		return updatedContactDetail, err
	}
	err = json.Unmarshal(respBody, &updatedContactDetail)
	if err != nil {
		return updatedContactDetail, fmt.Errorf("error parsing body: %v", err)
	}
	return updatedContactDetail, nil
}

func handleContactDetailResponseCode(resp *http.Response, expectedStatus int) ([]byte, error) {
	respBody, err := shared.HandleCustomResponseCode(resp, expectedStatus)
	if err != nil {
		return respBody, fmt.Errorf("error returned by endpoint, status code: %d, body: %s: %w", resp.StatusCode, respBody, err)
	}
	return respBody, nil
}
