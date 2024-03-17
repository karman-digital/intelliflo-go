package intellifloapp

func (i TenantIntelliflo) GetTenantId() int {
	return i.TenantId.Int()
}
