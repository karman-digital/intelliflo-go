package advisers

import "github.com/karman-digital/intelliflo/intelliflo/api/credentials"

func NewAdviserObject(creds credentials.Credentials) *Adviser {
	return &Adviser{
		AdviserEndpoint: NewAdviserService(creds),
	}
}

func NewAdviserService(creds credentials.Credentials) *AdviserService {
	return &AdviserService{
		creds,
	}
}
