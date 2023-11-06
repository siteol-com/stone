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

// RouterBaseReq 路由基础对象
type RouterBaseReq struct {
	Name        string `json:"name" binding:"required,max=32" example:"开放账密登陆"`              // 路由名称，用于界面展示，与权限关联
	Url         string `json:"url" binding:"required,uri,max=64" example:"/open/auth/login"` // 路由地址，后端访问URL 后端不再URL中携带参数，统一Post处理内容
	ServiceCode string `json:"serviceCode" binding:"required,numeric" example:"1"`           // 业务编码（字典），为接口分组
	PrintReq    string `json:"printReq" binding:"required,oneof='1' '2'" example:"1"`        // 请求日志打印 1 不打印 2 打印
	PrintRes    string `json:"printRes" binding:"required,oneof='1' '2'" example:"2"`        // 响应日志打印 1 不打印 2 打印
}

// RouterAddReq 路由创建对象
type RouterAddReq struct {
	RouterBaseReq
	Type string `json:"type" binding:"required,oneof='1' '2'" example:"1"` // 免授权路由 1 授权 2 免授权
}

// RouterReqToDbReq 转换请求到数据库对象
func RouterReqToDbReq(addRed *RouterAddReq) *platDb.Router {
	if addRed != nil {
		dbReq := &platDb.Router{
			Name:        addRed.Name,
			Url:         addRed.Url,
			Type:        addRed.Type,
			ServiceCode: addRed.ServiceCode,
			PrintReq:    addRed.PrintReq,
			PrintRes:    addRed.PrintRes,
		}
		return dbReq
	}
	return nil
}

// RouterEditReq 路由编辑对象
type RouterEditReq struct {
	ID uint64 `json:"id"  binding:"required,numeric" example:"1"` // 默认数据ID
	RouterBaseReq
}

// RouterEditReqToDbReq 转换请求到数据库对象
func RouterEditReqToDbReq(editRed *RouterEditReq) *platDb.Router {
	if editRed != nil {
		dbReq := &platDb.Router{
			ID:          editRed.ID,
			Name:        editRed.Name,
			Url:         editRed.Url,
			ServiceCode: editRed.ServiceCode,
			PrintReq:    editRed.PrintReq,
			PrintRes:    editRed.PrintRes,
		}
		return dbReq
	}
	return nil
}
