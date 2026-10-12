package servicecases

import "github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"

func NewServiceCasesService(creds credentials.Credentials) *ServiceCasesService {
	return &ServiceCasesService{Credentials: creds}
}
