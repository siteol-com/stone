package platModel

// AuthLogin 账密登陆结构体
type AuthLogin struct {
	Account     string `json:"account" binding:"required"`     // 账号
	Password    string `json:"password" binding:"required"`    // 密码
	TenantAlias string `json:"tenantAlias" binding:"required"` // 租户别名
}
