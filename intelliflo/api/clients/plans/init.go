package plans

import "github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"

func NewPlansService(creds credentials.Credentials) *PlansService {
	return &PlansService{
		creds,
	}
}
