package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// PermissionRouter 权限路由表
type PermissionRouter struct {
	ID           uint64 // 默认数据ID
	PermissionId uint64 // 权限ID
	RouterId     uint64 // 路由ID
}

// PermissionRouterTable 权限路由泛型构造器
var PermissionRouterTable actuator.Table[PermissionRouter]

// DataBase 实现指定数据库
func (t PermissionRouter) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t PermissionRouter) TableName() string {
	return "permission_router"
}

// FindPermissionRouterByIds 获取权限对应的路由ID
func (t PermissionRouter) FindPermissionRouterByIds(permissionIds []uint64) (res []uint64, err error) {
	r := platDb.Distinct("router_id").Where("permission_id IN ?", permissionIds).Find(&res)
	err = r.Error
	return
}
