package contactdetails

import "github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"

func NewContactDetailsService(creds credentials.Credentials) *ContactDetailsService {
	return &ContactDetailsService{
		creds,
	}
}
