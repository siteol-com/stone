package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// RolePermission 角色权限表
type RolePermission struct {
	ID           uint64 // 默认数据ID
	RoleId       uint64 // 角色ID
	PermissionId uint64 // 权限ID
}

// RolePermissionTable 角色权限泛型构造器
var RolePermissionTable actuator.Table[RolePermission]

// DataBase 实现指定数据库
func (t RolePermission) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t RolePermission) TableName() string {
	return "role_permission"
}

// FindRolePermissionByIds 根据角色清单查权限（去重）
func (t RolePermission) FindRolePermissionByIds(roleIds []uint64) (res []uint64, err error) {
	r := platDb.Distinct("permission_id").Where("role_id IN ?", roleIds).Find(&res)
	err = r.Error
	return
}
