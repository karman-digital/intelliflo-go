package intelliflomodels

type Relationship struct {
	ID                             int              `json:"id,omitempty"`
	Href                           string           `json:"href,omitempty"`
	RelationshipType               RelationshipType `json:"relationshipType"`
	Subject                        RelationSubject  `json:"subject"`
	Relation                       RelationSubject  `json:"relation"`
	IsFinancialPartnership         bool             `json:"isFinancialPartnership,omitempty"`
	IsFamilyGroup                  bool             `json:"isFamilyGroup,omitempty"`
	IncludeInRelationsFamilyWealth bool             `json:"includeInRelationsFamilyWealth,omitempty"`
	IsPointOfContact               bool             `json:"isPointOfContact,omitempty"`
	StartedOn                      string           `json:"startedOn,omitempty"`
}

type RelationshipType struct {
	Name string `json:"name"`
}

type RelationSubject struct {
	ID                  int    `json:"id"`
	Type                string `json:"type"`
	Href                string `json:"href,omitempty"`
	IsHeadOfFamilyGroup bool   `json:"isHeadOfFamilyGroup,omitempty"`
	PartyType           string `json:"partyType,omitempty"`
	AccessToRelation    Access `json:"accessToRelation,omitempty"`
}

type Access struct {
	GrantedAt string    `json:"grantedAt,omitempty"`
	GrantedBy GrantedBy `json:"grantedBy,omitempty"`
}

type GrantedBy struct {
	ID   int    `json:"id,omitempty"`
	Href string `json:"href,omitempty"`
}
