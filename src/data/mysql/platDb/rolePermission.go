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
	r := platDb.Table(t.TableName()).Distinct("permission_id").Where("role_id IN ?", roleIds).Find(&res)
	err = r.Error
	return
}

// FindPermissionRoleByIds 根据权限清单查询关联的角色（去重）
func (t RolePermission) FindPermissionRoleByIds(permissionIds []uint64) (res []uint64, err error) {
	r := platDb.Table(t.TableName()).Distinct("role_id").Where("permission_id IN ?", permissionIds).Find(&res)
	err = r.Error
	return
}

// DeleteByRoleId 根绝角色ID删除关联的权限
func (t RolePermission) DeleteByRoleId(roleId uint64) (err error) {
	r := platDb.Where("role_id = ?", roleId).Delete(&t)
	err = r.Error
	return
}

// DeleteByPermissionId 根绝权限ID删除关联的权限
func (t RolePermission) DeleteByPermissionId(permissionId uint64) (err error) {
	r := platDb.Where("permission_id = ?", permissionId).Delete(&t)
	err = r.Error
	return
}
