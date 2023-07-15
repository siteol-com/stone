package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// Dept 部门表
type Dept struct {
	Id             uint64 // 默认数据ID
	Name           string // 部门名称
	Pid            uint64 // 父级部门ID，租户创建时默认创建根部门，父级ID=0
	PermissionType string // 权限类型 0全局数据 2仅本部门 3本部门及子部门
	TenantId       uint64 // 租户ID
	Common
}

// DeptTable 部门泛型构造器
var DeptTable actuator.Table[Dept]

// DataBase 实现指定数据库
func (t Dept) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Dept) TableName() string {
	return "dept"
}
