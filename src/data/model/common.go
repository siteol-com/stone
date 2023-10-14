package model

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
