package platModel

// AuthLoginReq 账密登陆结构体
type AuthLoginReq struct {
	Account     string `json:"account" binding:"required" example:"stone"`     // 账号
	Password    string `json:"password" binding:"required" example:"123456"`   // 密码
	TenantAlias string `json:"tenantAlias" binding:"required" example:"stone"` // 租户别名
}

// AuthLoginRes 账密登陆响应
type AuthLoginRes struct {
	Token string `json:"token" example:"20230203090105DpC4K9Tu8IYagEgGM7"` // 32位登陆Token，时间串+18位随机字符
}
