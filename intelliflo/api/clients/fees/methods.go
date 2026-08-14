package fees

import (
	"encoding/json"
	"fmt"
	"net/http"

	feesmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/fees"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (c *FeeService) GetFee(clientID, feeID int) (feesmodels.Fee, error) {
	var fee feesmodels.Fee
	resp, err := c.SendRequest(http.MethodGet, fmt.Sprintf("clients/%d/fees/%d", clientID, feeID), nil)
	if err != nil {
		return fee, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return fee, fmt.Errorf("error handling response code: %w", err)
	}
	if err := json.Unmarshal(respBody, &fee); err != nil {
		return fee, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return fee, nil
}
