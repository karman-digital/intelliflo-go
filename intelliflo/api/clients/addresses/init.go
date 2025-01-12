package addresses

import "github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"

func NewAddressService(creds credentials.Credentials) *AddressService {
	return &AddressService{
		creds,
	}
}
