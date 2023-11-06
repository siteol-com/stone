package platModel

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"time"
)

// ResponsePageReq 响应码分页请求
type ResponsePageReq struct {
	Code        string `json:"code" example:"F7000"`                                       // 响应码
	ServiceCode string `json:"serviceCode" binding:"omitempty,numeric" example:"7"`        // 业务编码，仅允许提交数字
	Type        string `json:"responseType" binding:"omitempty,oneof='S' 'F'" example:"S"` // 响应类型，该字段用于筛选，可配置S和F
	model.PageReq
}

// ResponseBaseReq 响应码基础请求
type ResponseBaseReq struct {
	ZhCn   string `json:"zhCn" example:"角色查询失败"`            // 中文响应文言
	EnUs   string `json:"enUs" example:"Role query failed"` // 英文响应文言
	Remark string `json:"remark" example:"角色查询失败"`          // 其他备注信息
}

// ResponseAddReq 响应码添加请求
type ResponseAddReq struct {
	ServiceCode string `json:"serviceCode" binding:"required,numeric" example:"7"` // 业务ID，来源于字典，指定响应码归属业务
	Type        string `json:"type" binding:"required,oneof='S' 'F'" example:"F"`  // 响应类型，该字段用于筛选，可配置S和F
	ResponseBaseReq
}

// ResponseReqToDbReq 转换请求到数据库对象
func ResponseReqToDbReq(addRed *ResponseAddReq) *platDb.Response {
	if addRed != nil {
		dbReq := &platDb.Response{
			ServiceCode: addRed.ServiceCode,
			Type:        addRed.Type,
			ZhCn:        addRed.ZhCn,
			EnUs:        addRed.EnUs,
			Remark:      addRed.Remark,
			Mark:        constant.StatusOpen,
		}
		now := time.Now()
		dbReq.CreateAt = &now
		dbReq.Status = constant.StatusOpen
		return dbReq
	}
	return nil
}

// ResponseEditReq 响应码编辑请求
type ResponseEditReq struct {
	ID uint64 `json:"id"  binding:"required,numeric" example:"1"` // 数据ID
	ResponseBaseReq
}
