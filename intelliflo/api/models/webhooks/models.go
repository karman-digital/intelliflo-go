package webhookmodels

import (
	addressmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/address"
	clientmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/client"
	contactdetailmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/contactdetails"
	planmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/plans"
)

type WebhookPayload struct {
	ID           string         `json:"id"`
	Event        string         `json:"event"`
	TimeStamp    string         `json:"timeStamp"`
	InstalledFor map[string]any `json:"-"`
	Payload      map[string]any `json:"payload"`
	Version      string         `json:"version"`
	TenantId     int            `json:"-"`
	UserId       int            `json:"-"`
}

type AppWebhookPayload struct {
	ID        string `json:"id"`
	Event     string `json:"event"`
	TimeStamp string `json:"timeStamp"`
}

type AppInstallWebhookBody struct {
	AppWebhookPayload
	InstalledFor InstalledFor      `json:"installedFor"`
	Payload      AppInstallPayload `json:"payload"`
}

type AppUninstallWebhookBody struct {
	AppWebhookPayload
	UninstalledFor InstalledFor      `json:"uninstalledFor"`
	Payload        AppInstallPayload `json:"payload"`
}

type WebhookBody struct {
	ID        string `json:"id"`
	Event     string `json:"event"`
	TimeStamp string `json:"timeStamp"`
	TenantId  string `json:"tenantId"`
	UserId    string `json:"userId"`
}

type ClientChangedWebhookBody struct {
	WebhookBody
	Payload clientmodels.Client `json:"payload"`
}

type ContactDetailWebhookBody struct {
	WebhookBody
	Payload contactdetailmodels.ContactDetail `json:"payload"`
}

type AddressWebhookBody struct {
	WebhookBody
	Payload addressmodels.Residence `json:"payload"`
}

type PlanStatusChangeWebhookBody struct {
	WebhookBody
	Payload planmodels.PlanStatusChangePayload `json:"payload"`
}

type InstalledFor struct {
	Tenant              InstalledForID `json:"tenant"`
	User                InstalledForID `json:"user"`
	IsUninstalledForAll bool           `json:"isUninstalledForAll"`
}

type InstalledForID struct {
	ID int `json:"id"`
}

type AppInstallPayload struct {
	ID                   string       `json:"id"`
	Name                 string       `json:"name"`
	IsApprovedForInstall bool         `json:"isApprovedForInstall"`
	Summary              string       `json:"summary"`
	LastUpdatedAt        string       `json:"lastUpdatedAt"`
	BillingModel         BillingModel `json:"billingModel"`
	Version              string       `json:"version"`
}

type BillingModel struct {
	NetPrice         NetPrice `json:"netPrice"`
	Frequency        string   `json:"frequency"`
	ApiUsagePlan     string   `json:"apiUsagePlan"`
	InstallationType string   `json:"installationType"`
	ChargeBasis      string   `json:"chargeBasis"`
	CustomChargeDesc string   `json:"customChargeDesc"`
	ExternalRef      string   `json:"externalRef"`
}

type NetPrice struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type WebhookSubscriptionRequest struct {
	Topic        string `json:"topic"`
	Callback     string `json:"callback"`
	LeaseSeconds string `json:"leaseSeconds,omitempty"`
	Secret       string `json:"secret"`
}

type Webhooks struct {
	Href  string    `json:"href"`
	Items []Webhook `json:"items"`
	Count int32     `json:"count"`
}

type Webhook struct {
	ID           string `json:"id"`
	Href         string `json:"href"`
	Topic        string `json:"topic"`
	Callback     string `json:"callback"`
	LeaseSeconds int    `json:"leaseSeconds"`
}
