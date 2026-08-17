package fees

import (
	"encoding/json"
	"fmt"
	"net/http"

	feesmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/fees"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	intelliflohelpers "github.com/karman-digital/intelliflo-go/intelliflo/helpers"
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

func (c *FeeService) GetFees(clientID int) (feesmodels.Fees, error) {
	var fees feesmodels.Fees
	options := sharedmodels.GetOptions{Top: 500}

	for {
		resp, err := c.SendRequest(http.MethodGet, fmt.Sprintf("clients/%d/fees", clientID), nil, options)
		if err != nil {
			return feesmodels.Fees{}, fmt.Errorf("error sending request: %w", err)
		}
		respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
		resp.Body.Close()
		if err != nil {
			return feesmodels.Fees{}, fmt.Errorf("error handling response code: %w", err)
		}

		var page feesmodels.Fees
		if err := json.Unmarshal(respBody, &page); err != nil {
			return feesmodels.Fees{}, fmt.Errorf("error unmarshalling response: %w", err)
		}

		if fees.Href == "" {
			fees.Href = page.Href
			fees.FirstHref = page.FirstHref
			fees.LastHref = page.LastHref
			fees.PrevHref = page.PrevHref
		}
		fees.NextHref = page.NextHref
		fees.Count = page.Count
		fees.Items = append(fees.Items, page.Items...)

		if page.NextHref == "" {
			return fees, nil
		}

		skip, err := intelliflohelpers.ExtractSkipValueFromIntellifloResponse(page.NextHref)
		if err != nil {
			return feesmodels.Fees{}, fmt.Errorf("error extracting skip value: %w", err)
		}
		options.Skip = skip
	}
}
