package holdings

import "github.com/karman-digital/intelliflo/intelliflo/api/credentials"

func NewHoldingService(creds credentials.Credentials) *HoldingService {
	return &HoldingService{
		creds,
	}
}
