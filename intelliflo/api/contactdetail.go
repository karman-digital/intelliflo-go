package intelliflo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"
)

func (c *credentials) GetContactDetails(clientId int) (intelliflomodels.ContactDetails, error) {
	var contactDetails intelliflomodels.ContactDetails
	reqUrl := fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/contactdetails", clientId)
	req, err := retryablehttp.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return contactDetails, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, httpErr := c.client.Do(req)
	respBody, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return contactDetails, fmt.Errorf("error reading body: %v", err)
	}
	if httpErr != nil {
		return contactDetails, fmt.Errorf("error making get request: %v, with body: %v", httpErr, string(respBody))
	}
	if resp.StatusCode != http.StatusOK {
		return contactDetails, fmt.Errorf("error returned by endpoint: %v, with body: %v", resp.StatusCode, string(respBody))
	}
	err = json.Unmarshal(respBody, &contactDetails)
	if err != nil {
		return contactDetails, fmt.Errorf("error parsing body: %v", err)
	}
	return contactDetails, nil
}

func (c *credentials) PostContactDetail(clientId int, contactDetail intelliflomodels.ContactDetail) (intelliflomodels.ContactDetail, error) {
	var newContactDetail intelliflomodels.ContactDetail
	reqUrl := fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/contactdetails", clientId)
	reqBody, err := json.Marshal(contactDetail)
	if err != nil {
		return newContactDetail, fmt.Errorf("error marshalling contactDetail: %v", err)
	}
	req, err := retryablehttp.NewRequest("POST", reqUrl, reqBody)
	if err != nil {
		return newContactDetail, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return newContactDetail, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return newContactDetail, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return newContactDetail, fmt.Errorf("error returned by endpoint: %v, with body: %v", resp.StatusCode, string(respBody))
	}
	err = json.Unmarshal(respBody, &newContactDetail)
	if err != nil {
		return newContactDetail, fmt.Errorf("error parsing body: %v", err)
	}
	return newContactDetail, nil
}

func (c *credentials) PutContactDetail(clientId int, contactDetailId int, contactDetail intelliflomodels.ContactDetail) (intelliflomodels.ContactDetail, error) {
	var updatedContactDetail intelliflomodels.ContactDetail
	reqUrl := fmt.Sprintf("https://api.gb.intelliflo.net/v2/clients/%d/contactdetails/%d", clientId, contactDetailId)
	reqBody, err := json.Marshal(contactDetail)
	if err != nil {
		return updatedContactDetail, fmt.Errorf("error marshalling contactDetail: %v", err)
	}
	req, err := retryablehttp.NewRequest("PUT", reqUrl, reqBody)
	if err != nil {
		return updatedContactDetail, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.apiKey.String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.accessToken)}
	resp, err := c.client.Do(req)
	if err != nil {
		return updatedContactDetail, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return updatedContactDetail, fmt.Errorf("error reading body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return updatedContactDetail, fmt.Errorf("error returned by endpoint: %v, with body: %v", resp.StatusCode, string(respBody))
	}
	err = json.Unmarshal(respBody, &updatedContactDetail)
	if err != nil {
		return updatedContactDetail, fmt.Errorf("error parsing body: %v", err)
	}
	return updatedContactDetail, nil
}
