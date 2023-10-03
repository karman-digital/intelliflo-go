package intelliflo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/hatch-models/intelliflomodels"
)

func (c *Credentials) CreateNewClient(clientObj intelliflomodels.Client) (intelliflomodels.Client, error) {
	responseClient := intelliflomodels.Client{}
	clientBody, err := json.Marshal(clientObj)
	if err != nil {
		return responseClient, fmt.Errorf("error converting to body: %v", err)
	}
	clientReader := bytes.NewReader(clientBody)
	clientReq, err := retryablehttp.NewRequest("POST", "https://api.gb.intelliflo.net/v2/clients", clientReader)
	if err != nil {
		return responseClient, fmt.Errorf("error creating request: %v", err)
	}
	clientReq.Header.Set("Content-Type", "application/json")
	clientReq.Header["x-api-key"] = []string{c.APIKey}
	clientReq.Header["authorization"] = []string{"Bearer " + c.AccessToken}
	clientResp, err := c.Client.Do(clientReq)
	if err != nil {
		return responseClient, fmt.Errorf("error making post request: %v", err)
	}
	respBody, err := io.ReadAll(clientResp.Body)
	if err != nil {
		return responseClient, fmt.Errorf("error reading body: %v", err)
	}
	if clientResp.StatusCode != http.StatusCreated {
		return responseClient, fmt.Errorf("error when creating client: %s", string(respBody))
	}
	defer clientResp.Body.Close()
	err = json.Unmarshal(respBody, &responseClient)
	if err != nil {
		return responseClient, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return responseClient, nil
}
