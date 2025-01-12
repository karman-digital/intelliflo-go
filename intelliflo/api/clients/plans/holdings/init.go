package holdings

import "github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"

func NewHoldingService(creds credentials.Credentials) *HoldingService {
	return &HoldingService{
		creds,
	}
}
