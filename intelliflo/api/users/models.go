package users

import (
	"github.com/karman-digital/intelliflo/intelliflo/api/credentials"
	"github.com/karman-digital/intelliflo/intelliflo/interfaces"
)

type User struct {
	interfaces.User
}

type UserService struct {
	credentials.Credentials
}
