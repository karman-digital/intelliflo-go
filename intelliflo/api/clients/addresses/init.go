package addresses

import "github.com/karman-digital/intelliflo/intelliflo/api/credentials"

func NewAddressService(creds credentials.Credentials) *AddressService {
	return &AddressService{
		creds,
	}
}
