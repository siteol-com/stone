package model

// DemoReq 演示查询对象
type DemoReq struct {
	HttpCode int `json:"httpCode" example:"200"` // 响应HTTPCode，不传默认响应200，支持200/400/401/403/500
}

// IdReq ID查询对象
type IdReq struct {
	ID uint64 `json:"id" binding:"required" example:"1"` // 数据ID
}

// SortReq 排序对象
type SortReq struct {
	ID   uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	Sort uint16 `json:"sort" example:"1"`                  // 序号
}

// TenantModel 租户接口
type TenantModel interface {
	SetTenantId(id uint64)
	// GetTenantId() uint64 对象暂时不需要统一读取
}

// IdAnTenantReq ID查询对象（租户过滤）
type IdAnTenantReq struct {
	ID       uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	TenantId uint64 `json:"-"`                                 // 租户ID
}

// SetTenantId IdAnTenantReq 设置租户ID
func (t *IdAnTenantReq) SetTenantId(id uint64) {
	t.TenantId = id
}

// Tree 树对象
type Tree struct {
	Title    string  `json:"title" example:"根节点"` // 树标题
	Key      string  `json:"key" example:"ROOT"`  // 树键
	Children []*Tree `json:"children"`            // 子树
	Level    string  `json:"level" example:"1"`   // 表示树等级
	Id       uint64  `json:"-"`                   // 表示树数据ID
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
