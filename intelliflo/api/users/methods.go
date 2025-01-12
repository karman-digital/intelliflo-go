package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	usersmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/admin/users"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

func (c *UserService) GetUserById(id int) (usersmodels.User, error) {
	var user usersmodels.User
	resp, err := c.SendRequest("GET", fmt.Sprintf("users/%d", id), nil)
	if err != nil {
		return user, fmt.Errorf("error sending request: %v", err)
	}
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return user, fmt.Errorf("error handling response code: %v", err)
	}
	err = json.Unmarshal(respBody, &user)
	if err != nil {
		return user, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return user, nil
}

func (c *UserService) GetUsersByEmail(email string) (usersmodels.Users, error) {
	var users usersmodels.Users
	resp, err := c.SendRequest("GET", "users", nil, sharedmodels.GetOptions{
		Filter: fmt.Sprintf("email eq '%s'", email),
	})
	if err != nil {
		return users, fmt.Errorf("error sending request: %v", err)
	}
	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return users, fmt.Errorf("error handling response code: %v", err)
	}
	err = json.Unmarshal(respBody, &users)
	if err != nil {
		return users, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return users, nil
}
