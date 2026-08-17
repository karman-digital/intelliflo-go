package servicecases

import (
	"encoding/json"
	"fmt"
	"net/http"

	servicecasemodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/servicecases"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (s *ServiceCasesService) GetServiceCases(clientID int, options ...sharedmodels.GetOptions) (servicecasemodels.ServiceCases, error) {
	var serviceCases servicecasemodels.ServiceCases
	response, err := s.SendRequest("GET", fmt.Sprintf("clients/%d/servicecases", clientID), nil, options...)
	if err != nil {
		return serviceCases, fmt.Errorf("error sending request: %v", err)
	}
	defer response.Body.Close()
	responseBody, err := shared.HandleCustomResponseCode(response, http.StatusOK)
	if err != nil {
		return serviceCases, fmt.Errorf("error handling response code: %v", err)
	}
	if err := json.Unmarshal(responseBody, &serviceCases); err != nil {
		return serviceCases, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return serviceCases, nil
}
