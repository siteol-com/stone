package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// Router 路由表
type Router struct {
	ID          uint64 `json:"id"  binding:"numeric"`                     // 默认数据ID
	Name        string `json:"name" binding:"required,max=32"`            // 路由名称，用于界面展示，与权限关联
	Url         string `json:"url" binding:"required,uri,max=64"`         // 路由地址，后端访问URL 后端不再URL中携带参数，统一Post处理内容
	ServiceCode string `json:"serviceCode" binding:"required,numeric"`    // 业务编码（字典），为接口分组
	Type        string `json:"type" binding:"required,oneof='1' '2'"`     // 免授权路由 1 授权 2 免授权（系统启动开放免授权）
	PrintReq    string `json:"printReq" binding:"required,oneof='1' '2'"` // 请求日志打印 1 不打印 2 打印
	PrintRes    string `json:"printRes" binding:"required,oneof='1' '2'"` // 响应日志打印 1 不打印 2 打印
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
