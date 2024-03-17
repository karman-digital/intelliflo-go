package contactdetails

import "github.com/karman-digital/intelliflo/intelliflo/api/credentials"

func NewContactDetailsService(creds credentials.Credentials) *ContactDetailsService {
	return &ContactDetailsService{
		creds,
	}
}
