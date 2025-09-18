package interfaces

import (
	activitiesmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/activities"
	usersmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/admin/users"
	advisermodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/adviser"
	opportunitiesmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/opportunities"
	addressmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/address"
	clientmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/client"
	contactdetailmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/contactdetails"
	marketingpreferencemodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/marketingpreferences"
	relationshipmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/relationships"
	planmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/plans"
	fundsmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/plans/funds"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	taskmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/tasks"
	webhookmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/webhooks"
)

type Auth interface {
	RefreshToken(tenantId sharedmodels.TenantId, scope []string) error
	ValidateToken() error
}

type Client interface {
	GetClient(clientId int) (clientmodels.Client, error)
	GetClients(options ...sharedmodels.GetOptions) (clientmodels.Clients, error)
	PostClient(clientObj clientmodels.Client) (clientmodels.Client, error)
	PutClient(clientId int, client clientmodels.Client) (clientmodels.Client, error)
}

type Address interface {
	GetAddresses(entityId int, options ...sharedmodels.GetOptions) (addressmodels.Addresses, error)
	PostAddress(entityId int, address addressmodels.Residence) (addressmodels.Residence, error)
	PutAddress(entityId int, addressId int, address addressmodels.Residence) (addressmodels.Residence, error)
}

type ContactDetail interface {
	GetContactDetails(entityId int, options ...sharedmodels.GetOptions) (contactdetailmodels.ContactDetails, error)
	PostContactDetail(entityId int, contactDetail contactdetailmodels.ContactDetail) (contactdetailmodels.ContactDetail, error)
	PutContactDetail(entityId int, contactDetailId int, contactDetail contactdetailmodels.ContactDetail) (contactdetailmodels.ContactDetail, error)
}

type Plan interface {
	GetPlans(clientId int, options ...sharedmodels.GetOptions) (planmodels.Plans, error)
}

type Holding interface {
	GetHoldings(clientId int, planId int, options ...sharedmodels.GetOptions) (fundsmodels.Holdings, error)
}

type Relationship interface {
	PostRelationship(cliendId int, postBody relationshipmodels.Relationship) (relationshipmodels.Relationship, error)
	GetRelationships(clientId int, options ...sharedmodels.GetOptions) (relationshipmodels.Relationships, error)
}

type MarketingPreference interface {
	GetMarketingPreference(clientId int) (marketingpreferencemodels.Preferences, error)
	PutMarketingPreference(clientId int, body marketingpreferencemodels.Preferences) (marketingpreferencemodels.Preferences, error)
}

type User interface {
	GetUserById(id int) (usersmodels.User, error)
	GetUsersByEmail(email string) (usersmodels.Users, error)
}

type Adviser interface {
	GetAdvisersByUserId(userId int) (advisermodels.Advisers, error)
	GetAdvisers(options ...sharedmodels.GetOptions) (advisermodels.Advisers, error)
	GetAdviser(adviserId int) (advisermodels.Adviser, error)
}

type Webhook interface {
	GetActiveWebhooks() (webhookmodels.Webhooks, error)
	GetWebhook(id int) (webhookmodels.Webhook, error)
	PostWebhookSubscription(postBody webhookmodels.WebhookSubscriptionRequest) (webhookmodels.Webhook, error)
}

type Activity interface {
	GetCategories(opts ...sharedmodels.GetOptions) (activitiesmodels.ActivityCategoryResponse, error)
	GetCategory(categoryId int, opts ...sharedmodels.GetOptions) (activitiesmodels.ActivityCategory, error)
	CreateCategory(category activitiesmodels.ActivityCategory) (activitiesmodels.ActivityCategory, error)
	UpdateCategory(categoryId int, category activitiesmodels.ActivityCategory) (activitiesmodels.ActivityCategory, error)
	DeleteCategory(categoryId int) error
	GetTypes(opts ...sharedmodels.GetOptions) (activitiesmodels.ActivityTypeResponse, error)
	GetType(typeId int, opts ...sharedmodels.GetOptions) (activitiesmodels.ActivityType, error)
	CreateType(activityType activitiesmodels.ActivityType) (activitiesmodels.ActivityType, error)
}

type Task interface {
	GetTask(taskId int, opts ...sharedmodels.GetOptions) (taskmodels.Task, error)
	GetTasks(opts ...sharedmodels.GetOptions) (taskmodels.TasksResponse, error)
	CreateTask(task taskmodels.Task) (taskmodels.Task, error)
	UpdateTask(taskId int, task taskmodels.Task) (taskmodels.Task, error)
	DeleteTask(taskId int) error
	GetTaskNotes(taskId int, opts ...sharedmodels.GetOptions) (taskmodels.TaskNotesResponse, error)
	CreateTaskNote(taskId int, note taskmodels.TaskNote) (taskmodels.TaskNote, error)
}

type Opportunity interface {
	CreateClientOpportunity(clientId int, opportunity opportunitiesmodels.OpportunityCreateRequest) (opportunitiesmodels.Opportunity, error)
	GetCampaigns(opts ...sharedmodels.GetOptions) (opportunitiesmodels.OpportunityCampaignOptionsResponse, error)
	GetCampaignTypes(opts ...sharedmodels.GetOptions) (opportunitiesmodels.OpportunityCampaignTypeOptionsResponse, error)
	GetStatuses(opts ...sharedmodels.GetOptions) (opportunitiesmodels.OpportunityStatusOptionsResponse, error)
	GetTypes(opts ...sharedmodels.GetOptions) (opportunitiesmodels.OpportunityTypeOptionsResponse, error)
}
