package activities

import (
	"encoding/json"
	"fmt"
	"net/http"

	activitiesmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/activities"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (s *ActivityService) GetCategories(opts ...sharedmodels.GetOptions) (activitiesmodels.ActivityCategoryResponse, error) {
	var categories activitiesmodels.ActivityCategoryResponse
	resp, err := s.SendRequest("GET", "/v2/activities/categories", nil, opts...)
	if err != nil {
		return categories, fmt.Errorf("error making get request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return categories, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &categories)
	if err != nil {
		return categories, fmt.Errorf("error parsing body: %v", err)
	}
	return categories, nil
}

func (s *ActivityService) GetCategory(categoryId int, opts ...sharedmodels.GetOptions) (activitiesmodels.ActivityCategory, error) {
	var category activitiesmodels.ActivityCategory
	resp, err := s.SendRequest("GET", fmt.Sprintf("/v2/activities/categories/%d", categoryId), nil, opts...)
	if err != nil {
		return category, fmt.Errorf("error making get request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return category, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &category)
	if err != nil {
		return category, fmt.Errorf("error parsing body: %v", err)
	}
	return category, nil
}

func (s *ActivityService) CreateCategory(category activitiesmodels.ActivityCategory) (activitiesmodels.ActivityCategory, error) {
	var newCategory activitiesmodels.ActivityCategory
	reqBody, err := json.Marshal(category)
	if err != nil {
		return newCategory, fmt.Errorf("error converting to body: %v", err)
	}
	resp, err := s.SendRequest("POST", "/v2/activities/categories", reqBody)
	if err != nil {
		return newCategory, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusCreated)
	if err != nil {
		return newCategory, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &newCategory)
	if err != nil {
		return newCategory, fmt.Errorf("error parsing body: %v", err)
	}
	return newCategory, nil
}

func (s *ActivityService) UpdateCategory(categoryId int, category activitiesmodels.ActivityCategory) (activitiesmodels.ActivityCategory, error) {
	var updatedCategory activitiesmodels.ActivityCategory
	reqBody, err := json.Marshal(category)
	if err != nil {
		return updatedCategory, fmt.Errorf("error converting to body: %v", err)
	}
	resp, err := s.SendRequest("PUT", fmt.Sprintf("/v2/activities/categories/%d", categoryId), reqBody)
	if err != nil {
		return updatedCategory, fmt.Errorf("error making put request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return updatedCategory, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &updatedCategory)
	if err != nil {
		return updatedCategory, fmt.Errorf("error parsing body: %v", err)
	}
	return updatedCategory, nil
}

func (s *ActivityService) DeleteCategory(categoryId int) error {
	resp, err := s.SendRequest("DELETE", fmt.Sprintf("/v2/activities/categories/%d", categoryId), nil)
	if err != nil {
		return fmt.Errorf("error making delete request: %v", err)
	}
	defer resp.Body.Close()
	_, err = shared.HandleCustomResponseCode(resp, http.StatusNoContent)
	if err != nil {
		return fmt.Errorf("error returned by endpoint, status code: %d", resp.StatusCode)
	}
	return nil
}

func (s *ActivityService) GetTypes(opts ...sharedmodels.GetOptions) (activitiesmodels.ActivityTypeResponse, error) {
	var types activitiesmodels.ActivityTypeResponse
	resp, err := s.SendRequest("GET", "/v2/activities/types", nil, opts...)
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

func (s *ActivityService) GetType(typeId int, opts ...sharedmodels.GetOptions) (activitiesmodels.ActivityType, error) {
	var activityType activitiesmodels.ActivityType
	resp, err := s.SendRequest("GET", fmt.Sprintf("/v2/activities/types/%d", typeId), nil, opts...)
	if err != nil {
		return activityType, fmt.Errorf("error making get request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return activityType, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &activityType)
	if err != nil {
		return activityType, fmt.Errorf("error parsing body: %v", err)
	}
	return activityType, nil
}

func (s *ActivityService) CreateType(activityType activitiesmodels.ActivityType) (activitiesmodels.ActivityType, error) {
	var newActivityType activitiesmodels.ActivityType
	reqBody, err := json.Marshal(activityType)
	if err != nil {
		return newActivityType, fmt.Errorf("error converting to body: %v", err)
	}
	resp, err := s.SendRequest("POST", "/v2/activities/types", reqBody)
	if err != nil {
		return newActivityType, fmt.Errorf("error making post request: %v", err)
	}
	defer resp.Body.Close()
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusCreated)
	if err != nil {
		return newActivityType, fmt.Errorf("error returned by endpoint, status code: %d, body: %s", resp.StatusCode, respBody)
	}
	err = json.Unmarshal(respBody, &newActivityType)
	if err != nil {
		return newActivityType, fmt.Errorf("error parsing body: %v", err)
	}
	return newActivityType, nil
}
