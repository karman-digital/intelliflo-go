package users

import "github.com/karman-digital/intelliflo/intelliflo/api/credentials"

func NewUser(creds credentials.Credentials) *User {
	return &User{
		NewUserService(creds),
	}
}

func NewUserService(creds credentials.Credentials) *UserService {
	return &UserService{
		creds,
	}
}
