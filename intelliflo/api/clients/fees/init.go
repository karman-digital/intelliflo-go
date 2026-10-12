package fees

import "github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"

func NewFeeService(creds credentials.Credentials) *FeeService {
	return &FeeService{Credentials: creds}
}
