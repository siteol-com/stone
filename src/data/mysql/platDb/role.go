package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// Role 角色表
type Role struct {
	ID       uint64 // 默认数据ID
	Name     string // 角色名称
	Remark   string // 角色备注
	TenantId uint64 // 租户ID
	Mark     uint8  // 变更标识 1可变更2禁止变更
	Common
}

// RoleTable 角色泛型构造器
var RoleTable actuator.Table[Role]

// DataBase 实现指定数据库
func (t Role) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Role) TableName() string {
	return "role"
}

// FindByIds 根据角色获取ID列表
func (t Role) FindByIds(ids []uint64) (res []*Role, err error) {
	r := platDb.Where("id IN ?", ids).Find(&res)
	err = r.Error
	return
}
