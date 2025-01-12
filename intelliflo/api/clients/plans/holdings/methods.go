package holdings

import (
	"encoding/json"
	"fmt"
	"net/http"

	fundsmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/plans/funds"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (c *HoldingService) GetHoldings(clientId int, planId int, options ...sharedmodels.GetOptions) (fundsmodels.Holdings, error) {
	var holdings fundsmodels.Holdings
	resp, err := c.SendRequest("GET", fmt.Sprintf("clients/%d/plans/%d/holdings", clientId, planId), nil, options...)
	if err != nil {
		return holdings, fmt.Errorf("error sending request: %v", err)
	}
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return holdings, fmt.Errorf("error handling response code: %v", err)
	}
	err = json.Unmarshal(respBody, &holdings)
	if err != nil {
		return holdings, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return holdings, nil
}
