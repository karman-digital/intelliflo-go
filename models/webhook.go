package intelliflomodels

type WebhookPayload struct {
	ID           string         `json:"id"`
	Event        string         `json:"event"`
	TimeStamp    string         `json:"timeStamp"`
	InstalledFor map[string]any `json:"installedFor"`
	Payload      map[string]any `json:"payload"`
	Version      string         `json:"version"`
	TenantId     int            `json:"tenantId"`
	UserId       int            `json:"userId"`
}

type AppInstallWebhookBody struct {
	ID           string            `json:"id"`
	Event        string            `json:"event"`
	TimeStamp    string            `json:"timeStamp"`
	InstalledFor InstalledFor      `json:"installedFor"`
	Payload      AppInstallPayload `json:"payload"`
}

type ClientChangedWebhookBody struct {
	ID        string `json:"id"`
	Event     string `json:"event"`
	TimeStamp string `json:"timeStamp"`
	TenantId  int    `json:"tenantId"`
	UserId    int    `json:"userId"`
	Payload   Client `json:"payload"`
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
