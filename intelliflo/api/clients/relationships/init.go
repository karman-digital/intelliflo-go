package relationships

import "github.com/karman-digital/intelliflo/intelliflo/api/credentials"

func NewRelationshipService(creds credentials.Credentials) *RelationshipService {
	return &RelationshipService{
		creds,
	}
}
