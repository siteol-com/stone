package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
	"time"
)

// Tenant 租戶表
type Tenant struct {
	ID          uint64     // 默认数据ID
	Name        string     // 租户名称
	Alias       string     // 租户别名，纯英文，将作为前端登陆URL的一部分
	ServiceType string     // 业务类型（字典）1：运营商(专) 2：图文资讯 3：工业物联
	Type        string     // 租户类型（字典）1：运营商  2：加盟商 3：业务商
	Theme       string     // 登陆模板（字典）
	Logo        string     // 租户Logo，建议上传两版（300*80）
	Icon        string     // 租户Icon，图标，浏览器标签图标
	ExpiryTime  *time.Time // 过期时间，过期后暂停服务
	Mark        string     // 变更标识 1可变更2禁止变更
	Common
}

// TenantTable 租户泛型构造器
var TenantTable actuator.Table[Tenant]

// DataBase 实现指定数据库
func (t Tenant) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Tenant) TableName() string {
	return "tenant"
}
