package platModel

// OpenTenantReq 开放租户查询
type OpenTenantReq struct {
	TenantAlias string `json:"tenantAlias" binding:"required"` // 租戶别名
}

// OpenTenantRes 开放租户响应
type OpenTenantRes struct {
	Name       string `json:"name"`       // 租戶名称
	Alias      string `json:"alias"`      // 租戶别名
	Theme      string `json:"theme"`      // 租户模板
	Logo       string `json:"logo"`       // 租户Logo
	Background string `json:"background"` // 租户背景CSS（图片或颜色）
}
