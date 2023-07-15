package platModel

// AuthLoginReq 账密登陆结构体
type AuthLoginReq struct {
	Account     string `json:"account" binding:"required"`     // 账号
	Password    string `json:"password" binding:"required"`    // 密码
	TenantAlias string `json:"tenantAlias" binding:"required"` // 租户别名
}

// AuthLoginRes 账密登陆响应
type AuthLoginRes struct {
	Token string `json:"token"` // 登陆Token
}

// AuthUser 授权对象
type AuthUser struct {
	UserId             uint64   `json:"userId"`             // 用户ID
	PwdExpTimeStr      string   `json:"pwdExpTimeStr"`      // 密码超期时间（修改后的90天）
	TenantId           uint64   `json:"tenantId"`           // 租户ID
	DeptId             uint64   `json:"deptId"`             // 部门ID
	PermissionType     string   `json:"permissionType"`     // 权限类型 0全局数据 1跟随部门 2仅本部门 3本部门及子部门 99无权限
	PermissionDeptList []uint64 `json:"permissionDeptList"` // 具有数据权限的部门ID列表
	PermissionList     []string `json:"permissionList"`     // 权限集（前端路由与使用）
	RouterList         []string `json:"routerList"`         // 路由集
}
