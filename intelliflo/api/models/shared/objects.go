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
