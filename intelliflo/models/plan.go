package models

import "time"

type ExchangeRate struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

type PlanType struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PortfolioCategory string `json:"portfolioCategory"`
	RegionCode        string `json:"regionCode"`
	AllowsSubPlans    bool   `json:"allowsSubPlans"`
}

type PlanSubObject struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

type IOSubObjectWithName struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
	Name string `json:"name"`
}

type Owner struct {
	Order int    `json:"order"`
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Href  string `json:"href"`
}

type LatestValuation struct {
	ID                     int                  `json:"id"`
	Href                   string               `json:"href"`
	Value                  LatestValuationValue `json:"value"`
	ValuedOn               time.Time            `json:"valuedOn"`
	CreatedAt              time.Time            `json:"createdAt"`
	SurrenderTransferValue LatestValuationValue `json:"surrenderTransferValue"`
	Type                   string               `json:"type"`
	Plan                   LatestValuationPlan  `json:"plan"`
	UpdatedBy              PlanSubObject        `json:"updatedBy"`
}

type LatestValuationPlan struct {
	PolicyNumber string `json:"policyNumber"`
	ID           int    `json:"id"`
	Href         string `json:"href"`
}

type LatestValuationValue struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type OtherReferences struct {
	MigrationReference   string `json:"migrationReference"`
	PortalReference      string `json:"portalReference"`
	PlatformReference    string `json:"platformReference"`
	ApplicationReference string `json:"applicationReference"`
}

type QuoteResult struct {
	ID    int    `json:"id"`
	Href  string `json:"href"`
	Quote struct {
		ID   int    `json:"id"`
		Href string `json:"href"`
	} `json:"quote"`
}

type ForwardIncomeTo struct {
	ID         int    `json:"id"`
	Href       string `json:"href"`
	UseBanding bool   `json:"useBanding"`
}

type AdviceStatus struct {
	Value     string    `json:"value"`
	UpdatedOn time.Time `json:"updatedOn"`
}

type DocumentDelivery struct {
	DeliveryMethod string `json:"deliveryMethod"`
}

type ProviderCodes struct {
	Code1 string `json:"code1"`
	Code2 string `json:"code2"`
	Code3 string `json:"code3"`
}

type Plan struct {
	ID                int                 `json:"id"`
	Href              string              `json:"href"`
	Currency          string              `json:"currency"`
	ExchangeRate      ExchangeRate        `json:"exchangeRate"`
	Discriminator     string              `json:"discriminator"`
	PlanType          PlanType            `json:"planType"`
	PolicyNumber      string              `json:"policyNumber"`
	StartsOn          time.Time           `json:"startsOn"`
	EndsOn            time.Time           `json:"endsOn"`
	ProductName       string              `json:"productName"`
	ProductProvider   IOSubObjectWithName `json:"productProvider"`
	SellingAdviser    PlanSubObject       `json:"sellingAdviser"`
	Owners            []Owner             `json:"owners"`
	IsVisibleToClient bool                `json:"isVisibleToClient"`
	CurrentStatus     string              `json:"currentStatus"`
	SystemStatus      string              `json:"systemStatus"`
	IsPreExisting     bool                `json:"isPreExisting"`
	Reference         string              `json:"reference"`
	LatestValuation   LatestValuation     `json:"latestValuation"`
	PlanTypesHref     string              `json:"planTypes_href"`
	ValuationsHref    string              `json:"valuations_href"`
	ContributionsHref string              `json:"contributions_href"`
	TopupsHref        string              `json:"topups_href"`
	PlanHoldingsHref  string              `json:"planHoldings_href"`
	Lifecycle         IOSubObjectWithName `json:"lifecycle"`
	IsTopup           bool                `json:"isTopup"`
	Parent            struct {
		ID   int    `json:"id"`
		Href string `json:"href"`
	} `json:"parent"`
	IsSubPlan                       bool                `json:"isSubPlan"`
	IsAdviceOffPanel                bool                `json:"isAdviceOffPanel"`
	OtherReferences                 OtherReferences     `json:"otherReferences"`
	ClientCategory                  string              `json:"clientCategory"`
	AvailablePlanPurposesHref       string              `json:"available_plan_purposes_href"`
	PlanPurposesHref                string              `json:"plan_purposes_href"`
	WithdrawalsHref                 string              `json:"withdrawals_href"`
	SubPlansHref                    string              `json:"subPlans_href"`
	IsClientSuitableForTargetMarket bool                `json:"isClientSuitableForTargetMarket"`
	QuoteResult                     QuoteResult         `json:"quoteResult"`
	Banding                         PlanSubObject       `json:"banding"`
	ForwardIncomeTo                 ForwardIncomeTo     `json:"forwardIncomeTo"`
	Administrator                   PlanSubObject       `json:"administrator"`
	Paraplanner                     PlanSubObject       `json:"paraplanner"`
	AdviceStatus                    AdviceStatus        `json:"adviceStatus"`
	Tags                            []string            `json:"tags"`
	Group                           PlanSubObject       `json:"group"`
	CreatedBy                       PlanSubObject       `json:"createdBy"`
	DocumentDelivery                DocumentDelivery    `json:"documentDelivery"`
	IsProviderManaged               bool                `json:"isProviderManaged"`
	PerformanceStartsOn             time.Time           `json:"performanceStartsOn"`
	PerformanceEndsOn               time.Time           `json:"performanceEndsOn"`
	ProviderCodes                   ProviderCodes       `json:"providerCodes"`
	CreatedAt                       time.Time           `json:"createdAt"`
	Program                         IOSubObjectWithName `json:"program"`
	RiskProfile                     PlanSubObject       `json:"riskProfile"`
}

type Plans struct {
	Href      string `json:"href"`
	FirstHref string `json:"first_href"`
	LastHref  string `json:"last_href"`
	NextHref  string `json:"next_href"`
	PrevHref  string `json:"prev_href"`
	Items     []Plan `json:"items"`
	Count     int    `json:"count"`
}
