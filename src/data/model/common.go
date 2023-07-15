package model

// IdReq ID查询对象
type IdReq struct {
	Id uint64 `json:"id" binding:"required"` // ID
}
