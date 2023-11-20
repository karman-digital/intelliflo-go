package intelliflomodels

type Relationship struct {
	ID                             int              `json:"id,omitempty"`
	Href                           string           `json:"href,omitempty"`
	RelationshipType               RelationshipType `json:"relationshipType"`
	Subject                        Subject          `json:"subject"`
	Relation                       Relation         `json:"relation"`
	IsFinancialPartnership         bool             `json:"isFinancialPartnership,omitempty"`
	IsFamilyGroup                  bool             `json:"isFamilyGroup,omitempty"`
	IncludeInRelationsFamilyWealth bool             `json:"includeInRelationsFamilyWealth,omitempty"`
	IsPointOfContact               bool             `json:"isPointOfContact,omitempty"`
	StartedOn                      string           `json:"startedOn,omitempty"`
}

type RelationshipType struct {
	Name string `json:"name"`
}

type Subject struct {
	ID                  int              `json:"id"`
	Type                string           `json:"type"`
	Href                string           `json:"href,omitempty"`
	IsHeadOfFamilyGroup bool             `json:"isHeadOfFamilyGroup,omitempty"`
	PartyType           string           `json:"partyType,omitempty"`
	AccessToRelation    AccessToRelation `json:"accessToRelation,omitempty"`
}

type Relation struct {
	ID                  int             `json:"id"`
	Type                string          `json:"type"`
	Href                string          `json:"href,omitempty"`
	IsHeadOfFamilyGroup bool            `json:"isHeadOfFamilyGroup,omitempty"`
	PartyType           string          `json:"partyType,omitempty"`
	AccessToSubject     AccessToSubject `json:"accessToSubject,omitempty"`
}

type AccessToRelation struct {
	GrantedAt string    `json:"grantedAt,omitempty"`
	GrantedBy GrantedBy `json:"grantedBy,omitempty"`
}

type AccessToSubject struct {
	GrantedAt string    `json:"grantedAt,omitempty"`
	GrantedBy GrantedBy `json:"grantedBy,omitempty"`
}

type GrantedBy struct {
	ID   int    `json:"id,omitempty"`
	Href string `json:"href,omitempty"`
}
