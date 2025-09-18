package opportunities

import (
	"encoding/json"
	"fmt"
	"net/http"

	opportunitiesmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/opportunities"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (s *OpportunityService) CreateClientOpportunity(clientId int, opportunity opportunitiesmodels.OpportunityCreateRequest) (opportunitiesmodels.Opportunity, error) {
	var created opportunitiesmodels.Opportunity
	reqBody, err := json.Marshal(opportunity)
	if err != nil {
		return created, fmt.Errorf("error converting to body: %v", err)
	}
	resp, err := s.SendRequest("POST", fmt.Sprintf("clients/%d/opportunities", clientId), reqBody)
	if err != nil {
		return created, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusCreated)
	if err != nil {
		return created, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &created)
	if err != nil {
		return created, fmt.Errorf("error parsing body: %v", err)
	}
	return created, nil
}

func (s *OpportunityService) GetCampaigns(opts ...sharedmodels.GetOptions) (opportunitiesmodels.OpportunityCampaignOptionsResponse, error) {
	var campaigns opportunitiesmodels.OpportunityCampaignOptionsResponse
	resp, err := s.SendRequest("GET", "opportunities/campaigns", nil, opts...)
	if err != nil {
		return campaigns, fmt.Errorf("error making get request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return campaigns, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &campaigns)
	if err != nil {
		return campaigns, fmt.Errorf("error parsing body: %v", err)
	}
	return campaigns, nil
}

func (s *OpportunityService) GetCampaignTypes(opts ...sharedmodels.GetOptions) (opportunitiesmodels.OpportunityCampaignTypeOptionsResponse, error) {
	var campaignTypes opportunitiesmodels.OpportunityCampaignTypeOptionsResponse
	resp, err := s.SendRequest("GET", "opportunities/campaigntypes", nil, opts...)
	if err != nil {
		return campaignTypes, fmt.Errorf("error making get request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return campaignTypes, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &campaignTypes)
	if err != nil {
		return campaignTypes, fmt.Errorf("error parsing body: %v", err)
	}
	return campaignTypes, nil
}

func (s *OpportunityService) GetStatuses(opts ...sharedmodels.GetOptions) (opportunitiesmodels.OpportunityStatusOptionsResponse, error) {
	var statuses opportunitiesmodels.OpportunityStatusOptionsResponse
	resp, err := s.SendRequest("GET", "opportunities/statuses", nil, opts...)
	if err != nil {
		return statuses, fmt.Errorf("error making get request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return statuses, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &statuses)
	if err != nil {
		return statuses, fmt.Errorf("error parsing body: %v", err)
	}
	return statuses, nil
}

func (s *OpportunityService) GetTypes(opts ...sharedmodels.GetOptions) (opportunitiesmodels.OpportunityTypeOptionsResponse, error) {
	var types opportunitiesmodels.OpportunityTypeOptionsResponse
	resp, err := s.SendRequest("GET", "opportunities/types", nil, opts...)
	if err != nil {
		return types, fmt.Errorf("error making get request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return types, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &types)
	if err != nil {
		return types, fmt.Errorf("error parsing body: %v", err)
	}
	return types, nil
}
