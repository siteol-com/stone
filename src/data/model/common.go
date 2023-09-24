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
