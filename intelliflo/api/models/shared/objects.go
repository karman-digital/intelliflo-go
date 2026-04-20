package sharedmodels

type IOSubObject struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
}

type GetOptions struct {
	Skip    int
	Filter  string
	OrderBy string
	Top     int
}

type Currency struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Prospect struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Href string `json:"href"`
}
