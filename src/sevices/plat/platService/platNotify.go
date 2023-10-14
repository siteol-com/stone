package platService

import (
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/utils/log"
)

// 路由变动通知
func notifyChangeByRouter(routerId uint64, remove bool, traceID string) (err error) {
	// DB对象
	permissionRouterDb := platDb.PermissionRouter{}
	// 获取路由绑定权限ID
	permissionIds, _ := permissionRouterDb.FindPermissionsByRouterId(routerId)
	// 如果需要移除关系（通过ID移除）
	if remove {
		err = permissionRouterDb.DeleteByRouterId(routerId)
		log.ErrorTF(traceID, "DeleteRouterPermission By RouterId %d Fail . Err Is : %v", routerId, err)
	}
	// 存在受影响的权限ID，向上同，权限和角色的关系
	if len(permissionIds) > 0 {
		// 异步通知受影响的Token失效（上层不涉及数据删除等行为）
		go notifyChangeByPermissionIds(permissionIds, false, traceID)
	}
	return
}

// 权限变动通知
func notifyChangeByPermissionIds(permissionIds []uint64, remove bool, traceID string) {
	// DB对象
	rolePermissionDb := platDb.RolePermission{}
	// 取的角色和权限关联的数据
	roleIds, _ := rolePermissionDb.FindPermissionRoleByIds(permissionIds)
	// 如果涉及删除行为，说明是某个权限执行了删除行为
	if remove {
		err := rolePermissionDb.DeleteByPermissionId(permissionIds[0])
		if err != nil {
			log.ErrorTF(traceID, "DeleteByPermissionId By %d Fail . Err Is : %v", permissionIds[0], err)
		}
	}
	// 存在受影响的角色，继续向上通知，角色和账号关系
	if len(roleIds) > 0 {
		notifyChangeByRoleIds(roleIds, false, traceID)
	}
}

// 角色变动通知
func notifyChangeByRoleIds(roleIds []uint64, remove bool, traceID string) {
	// TODO
}
