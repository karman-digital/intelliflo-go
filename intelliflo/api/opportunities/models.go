package opportunities

import (
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	"github.com/karman-digital/intelliflo-go/intelliflo/interfaces"
)

type Opportunity struct {
	OpportunityEndpoint interfaces.Opportunity
}

type OpportunityService struct {
	credentials.Credentials
}
