package platModel

import (
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/mysql/platDb"
)

// RouterPageReq 路由分页查询
type RouterPageReq struct {
	Name        string `json:"name"`                                    // 路由名称
	Url         string `json:"url"`                                     // 路由地址，仅允许提交URI
	ServiceCode string `json:"serviceCode" binding:"omitempty,numeric"` // 业务编码，仅允许提交数字
	Type        string `json:"type" binding:"omitempty,oneof='1' '2'"`  // 路由类型，仅允许提交1/2
	model.PageReq
}

// RouterPageRes struct
type RouterPageRes struct {
	*platDb.Router
	ServiceCodeLabel string `json:"serviceCodeLabel"` // 业务编码（字典）翻译
	TypeLabel        string `json:"typeLabel"`        // 路由类型翻译
}
