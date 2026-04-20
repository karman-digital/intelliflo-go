package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	usersmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/admin/users"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	intelliflohelpers "github.com/karman-digital/intelliflo-go/intelliflo/helpers"
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

func (c *UserService) GetAllUsers() (usersmodels.Users, error) {
	var allUsers usersmodels.Users
	options := sharedmodels.GetOptions{Top: 500}

	for {
		resp, err := c.SendRequest("GET", "users", nil, options)
		if err != nil {
			return usersmodels.Users{}, fmt.Errorf("error sending request: %v", err)
		}
		respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
		if err != nil {
			return usersmodels.Users{}, fmt.Errorf("error handling response code: %v", err)
		}

		var page usersmodels.Users
		err = json.Unmarshal(respBody, &page)
		if err != nil {
			return usersmodels.Users{}, fmt.Errorf("error unmarshalling response: %v", err)
		}

		if allUsers.Href == "" {
			allUsers.Href = page.Href
			allUsers.FirstHref = page.FirstHref
			allUsers.LastHref = page.LastHref
			allUsers.PrevHref = page.PrevHref
		}
		allUsers.NextHref = page.NextHref
		allUsers.Count = page.Count
		allUsers.Items = append(allUsers.Items, page.Items...)

		if page.NextHref == "" {
			break
		}

		skip, err := intelliflohelpers.ExtractSkipValueFromIntellifloResponse(page.NextHref)
		if err != nil {
			return usersmodels.Users{}, fmt.Errorf("error extracting skip value: %v", err)
		}
		options.Skip = skip
	}

	return allUsers, nil
}
