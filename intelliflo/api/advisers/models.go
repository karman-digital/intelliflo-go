package advisers

import (
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	"github.com/karman-digital/intelliflo-go/intelliflo/interfaces"
)

type Adviser struct {
	AdviserEndpoint interfaces.Adviser
}

type AdviserService struct {
	credentials.Credentials
}
