package intellifloapp

import intelliflomodels "github.com/karman-digital/intelliflo/intelliflo/api/models"

func (i TenantIntelliflo) GetTenantId() intelliflomodels.TenantId {
	return i.TenantId
}
