package advisers

import (
	"github.com/karman-digital/intelliflo/intelliflo/api/credentials"
	"github.com/karman-digital/intelliflo/intelliflo/interfaces"
)

type Adviser struct {
	AdviserEndpoint interfaces.Adviser
}

type AdviserService struct {
	credentials.Credentials
}
