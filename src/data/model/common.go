package model

// DemoReq 演示查询对象
type DemoReq struct {
	HttpCode int `json:"httpCode" example:"200"` // 响应HTTPCode，不传默认响应200，支持200/400/401/403/500
}

// IdReq ID查询对象
type IdReq struct {
	ID uint64 `json:"id" binding:"required"` // ID
}

// SortReq 排序对象
type SortReq struct {
	ID   uint64 `json:"id" binding:"required"` // ID
	Sort uint16 `json:"sort"`
}

// IdAnTenantReq ID查询对象（租户过滤）
type IdAnTenantReq struct {
	ID       uint64 `json:"id" binding:"required"` // ID
	TenantId uint64 `json:"-"`                     // 租户ID
}
