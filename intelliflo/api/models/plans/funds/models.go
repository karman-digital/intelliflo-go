package fundsmodels

import "time"

type Price struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Codes struct {
	Citi     string `json:"citi"`
	Epic     string `json:"epic"`
	Sedol    string `json:"sedol"`
	Isin     string `json:"isin"`
	Mex      string `json:"mex"`
	Provider string `json:"provider"`
	Apir     string `json:"apir"`
	FeedFund string `json:"feedFund"`
}

type Benchmark struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

type UnitPrice struct {
	BidPrice      Price `json:"bidPrice"`
	MidPrice      Price `json:"midPrice"`
	OfferPrice    Price `json:"offerPrice"`
	DailyChange   Price `json:"dailyChange"`
	YearHighPrice Price `json:"yearHighPrice"`
	YearLowPrice  Price `json:"yearLowPrice"`
}

type Manager struct {
	Name      string `json:"name"`
	Biography string `json:"biography"`
	IsAlpha   bool   `json:"isAlpha"`
	StartedOn string `json:"startedOn"`
}

type Provider struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Href string `json:"href"`
}

type Sector struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Href string `json:"href"`
}

type AllocationItem struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Allocation float64 `json:"allocation"`
}

type Breakdown struct {
	BreakdownOn time.Time        `json:"breakdownOn"`
	Items       []AllocationItem `json:"items"`
}

type Fund struct {
	ID                string    `json:"id"`
	Href              string    `json:"href"`
	Source            string    `json:"source"`
	Name              string    `json:"name"`
	ShortName         string    `json:"shortName"`
	Price             Price     `json:"price"`
	PricedOn          time.Time `json:"pricedOn"`
	FundPricesHref    string    `json:"fundPrices_href"`
	Type              string    `json:"type"`
	Codes             Codes     `json:"codes"`
	Currency          string    `json:"currency"`
	Yield             Price     `json:"yield"`
	IncAcc            string    `json:"incAcc"`
	InitialCharge     int       `json:"initialCharge"`
	ExitCharge        int       `json:"exitCharge"`
	AnnualMgmtCharge  int       `json:"annualMgmtCharge"`
	IsClosed          bool      `json:"isClosed"`
	Objective         string    `json:"objective"`
	LaunchedOn        string    `json:"launchedOn"`
	Size              Price     `json:"size"`
	SizedOn           string    `json:"sizedOn"`
	Benchmark         Benchmark `json:"benchmark"`
	ExpenseRatio      int       `json:"expenseRatio"`
	CrownRating       int       `json:"crownRating"`
	UnitPrice         UnitPrice `json:"unitPrice"`
	Manager           Manager   `json:"manager"`
	Provider          Provider  `json:"provider"`
	Sector            Sector    `json:"sector"`
	SectorBreakdown   Breakdown `json:"sector_breakdown"`
	RegionalBreakdown Breakdown `json:"regional_breakdown"`
	AssetBreakdown    Breakdown `json:"asset_breakdown"`
	HoldingBreakdown  Breakdown `json:"holding_breakdown"`
}

type FundPlan struct {
	PolicyNumber string `json:"policyNumber"`
	ID           int    `json:"id"`
	Href         string `json:"href"`
}

type Equity struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Codes struct {
		Epic string `json:"epic"`
		Isin string `json:"isin"`
	} `json:"codes"`
}

type Units struct {
	Number          float64   `json:"number"`
	NumberUpdatedOn time.Time `json:"numberUpdatedOn"`
	PriceUpdatedOn  time.Time `json:"priceUpdatedOn"`
	Price           Price     `json:"price"`
}

type Holding struct {
	Fund             Fund     `json:"fund"`
	ID               int      `json:"id"`
	Href             string   `json:"href"`
	Plan             FundPlan `json:"plan"`
	ParentHref       string   `json:"parent_href"`
	Equity           Equity   `json:"equity"`
	Units            Units    `json:"units"`
	TransactionsHref string   `json:"transactions_href"`
	TimeseriesHref   string   `json:"timeseries_href"`
}

type Holdings struct {
	Href      string    `json:"href"`
	FirstHref string    `json:"first_href"`
	LastHref  string    `json:"last_href"`
	NextHref  string    `json:"next_href"`
	PrevHref  string    `json:"prev_href"`
	Items     []Holding `json:"items"`
	Count     int       `json:"count"`
}
