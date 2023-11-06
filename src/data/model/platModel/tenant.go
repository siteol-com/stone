package platModel

// OpenTenantReq 开放租户查询
type OpenTenantReq struct {
	TenantAlias string `json:"tenantAlias" binding:"required" example:"stone"` // 租戶别名
}

// OpenTenantRes 开放租户响应
type OpenTenantRes struct {
	Name  string `json:"name" example:"基座租户"`                 // 租戶名称，全局唯一
	Alias string `json:"alias" example:"stone"`               // 租戶别名，全局唯一
	Theme string `json:"theme" example:"light"`               // 租户模板，登陆界面的风格模板
	Logo  string `json:"logo" example:"/static/img/logo"`     // 租户Logo，登陆和界面中的Logo
	Icon  string `json:"icon" example:"/static/img/icon.png"` // 租户Icon，浏览器顶部图标
}
