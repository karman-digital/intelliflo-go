package feesmodels

import (
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
)

type Fee struct {
	ID               int                              `json:"id"`
	Href             string                           `json:"href"`
	SequentialRef    string                           `json:"sequentialRef"`
	SellingAdviser   sharedmodels.IOSubObject         `json:"sellingAdviser"`
	AdviceCategory   string                           `json:"adviceCategory"`
	FeeType          FeeType                          `json:"feeType"`
	PaymentType      PaymentType                      `json:"paymentType"`
	FeeChargingType  FeeChargingType                  `json:"feeChargingType"`
	FeePercentage    float64                          `json:"feePercentage"`
	Net              Currency                         `json:"net"`
	Vat              Currency                         `json:"vat"`
	VatRate          VatRate                          `json:"vatRate"`
	Status           string                           `json:"status"`
	StatusDate       sharedmodels.IntellifloDateTime  `json:"statusDate"`
	Recurring        *Recurring                       `json:"recurring,omitempty"`
	InitialPeriod    int                              `json:"initialPeriod"`
	IsConsultancyFee bool                             `json:"isConsultancyFee"`
	Instalment       *Instalment                      `json:"instalment,omitempty"`
	SentToClientOn   *sharedmodels.IntellifloDateTime `json:"sentToClientOn,omitempty"`
	Discount         *Discount                        `json:"discount,omitempty"`
	Banding          *sharedmodels.IOSubObject        `json:"banding,omitempty"`
	ForwardIncomeTo  *ForwardIncomeTo                 `json:"forwardIncomeTo,omitempty"`
	Clients          []sharedmodels.IOSubObject       `json:"clients"`
	PlanHref         string                           `json:"plan_href"`
	Plans            []sharedmodels.IOSubObject       `json:"plans"`
	IsRetainer       bool                             `json:"isRetainer"`
	FeeModel         *FeeModel                        `json:"feeModel,omitempty"`
	Tierings         []Tiering                        `json:"tierings"`
	FeeDetail        FeeDetail                        `json:"feeDetail"`
	FeeCode          string                           `json:"feeCode"`
}

type FeeType struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type PaymentType struct {
	Name   string `json:"name"`
	PaidBy string `json:"paidBy"`
}

type FeeChargingType struct {
	Name          string    `json:"name"`
	MinFee        *Currency `json:"minFee,omitempty"`
	MaxFee        *Currency `json:"maxFee,omitempty"`
	MinPercentage *float64  `json:"minPercentage,omitempty"`
	MaxPercentage *float64  `json:"maxPercentage,omitempty"`
}

type Currency struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type VatRate struct {
	Rate     float64 `json:"rate"`
	IsExempt bool    `json:"isExempt"`
}

type Recurring struct {
	Frequency string                           `json:"frequency"`
	StartsOn  sharedmodels.IntellifloDateTime  `json:"startsOn"`
	EndsOn    *sharedmodels.IntellifloDateTime `json:"endsOn,omitempty"`
}

type Instalment struct {
	Frequency string `json:"frequency"`
	Count     int    `json:"count"`
}

type Discount struct {
	Name       string   `json:"name"`
	Percentage float64  `json:"percentage"`
	Total      Currency `json:"total"`
}

type ForwardIncomeTo struct {
	ID         int    `json:"id"`
	Href       string `json:"href"`
	UseBanding bool   `json:"useBanding"`
}

type FeeModel struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	TemplateID int    `json:"templateId"`
}

type Tiering struct {
	Threshold  Currency `json:"threshold"`
	Percentage float64  `json:"percentage"`
}

type FeeDetail struct {
	FeeChargingType          string     `json:"feeChargingType"`
	Name                     string     `json:"name"`
	MinFee                   *AmountFee `json:"minFee,omitempty"`
	MaxFee                   *AmountFee `json:"maxFee,omitempty"`
	MinPercentage            *float64   `json:"minPercentage,omitempty"`
	MaxPercentage            *float64   `json:"maxPercentage,omitempty"`
	IsPlanFee                bool       `json:"isPlanFee"`
	IsRegularContributionFee bool       `json:"isRegularContributionFee"`
	IsLumpSumContributionFee bool       `json:"isLumpSumContributionFee"`
	IsTransferFee            bool       `json:"isTransferFee"`
}

type AmountFee struct {
	Amount float64 `json:"amount"`
}
