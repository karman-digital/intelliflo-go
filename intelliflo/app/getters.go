package intellifloapp

import (
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
)

func (i TenantIntelliflo) GetTenantId() sharedmodels.TenantId {
	return i.TenantId
}
