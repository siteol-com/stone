package platDb

import (
	"siteOl.com/stone/server/src/data/mysql/actuator"
	"time"
)

// Tenant 租戶表
type Tenant struct {
	ID         uint64     // 租戶ID
	Name       string     // 租戶名称
	Alias      string     // 租戶别名
	Type       uint8      // 租户类型
	Theme      string     // 租户模板
	Logo       string     // 租户Logo
	Background string     // 租户背景CSS（图片或颜色）
	Mark       uint8      // 变更标识 0可变更 1禁止变更
	ExpiryTime *time.Time // 过期时间
	Common
}

// TableName 实现自定义表名
func (t *Tenant) TableName() string {
	return "tenant"
}

// FindOne 基于对象实施查询
func (t *Tenant) FindOne() (res *Tenant, err error) {
	res = new(Tenant)
	err = actuator.FindOneByObject(platDb, t, res)
	return
}
