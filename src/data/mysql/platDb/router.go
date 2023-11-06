package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// Router 路由表
type Router struct {
	ID          uint64 `json:"id" example:"1"`                 // 默认数据ID
	Name        string `json:"name" example:"开放账密登陆"`          // 路由名称，用于界面展示，与权限关联
	Url         string `json:"url" example:"/open/auth/login"` // 路由地址，后端访问URL 后端不再URL中携带参数，统一Post处理内容
	Type        string `json:"type" example:"1"`               // 免授权路由 1 授权 2 免授权
	ServiceCode string `json:"serviceCode" example:"1"`        // 业务编码（字典），为接口分组
	PrintReq    string `json:"printReq" example:"1"`           // 请求日志打印 1 不打印 2 打印
	PrintRes    string `json:"printRes" example:"2"`           // 响应日志打印 1 不打印 2 打印
}

// RouterTable 路由泛型造器
var RouterTable actuator.Table[Router]

// DataBase 实现指定数据库
func (t Router) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Router) TableName() string {
	return "router"
}

// FindByIds 获取路由清单
func (t Router) FindByIds(ids []uint64) (res []*Router, err error) {
	r := platDb.Where("id IN ?", ids).Find(&res)
	err = r.Error
	return
}
