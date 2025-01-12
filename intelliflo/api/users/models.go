package users

import (
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	"github.com/karman-digital/intelliflo-go/intelliflo/interfaces"
)

type User struct {
	interfaces.User
}

type UserService struct {
	credentials.Credentials
}
