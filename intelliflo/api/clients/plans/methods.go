package plans

import (
	"encoding/json"
	"fmt"
	"net/http"

	planmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/plans"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (c *PlansService) GetPlans(clientId int, options ...sharedmodels.GetOptions) (planmodels.Plans, error) {
	var plans planmodels.Plans
	resp, err := c.SendRequest("GET", fmt.Sprintf("clients/%d/plans", clientId), nil, options...)
	if err != nil {
		return plans, fmt.Errorf("error sending request: %v", err)
	}
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return plans, fmt.Errorf("error handling response code: %v", err)
	}
	defer resp.Body.Close()
	err = json.Unmarshal(respBody, &plans)
	if err != nil {
		return plans, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return plans, nil
}
