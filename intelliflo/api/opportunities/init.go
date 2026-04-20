package opportunities

import (
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
)

func NewOpportunityObject(creds credentials.Credentials) *Opportunity {
	return &Opportunity{
		OpportunityEndpoint: NewOpportunityService(creds),
	}
}

func NewOpportunityService(creds credentials.Credentials) *OpportunityService {
	return &OpportunityService{
		creds,
	}
}
