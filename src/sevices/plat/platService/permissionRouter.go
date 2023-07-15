package platService

import (
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/utils/log"
)

// 获取权限路由列表
func getPermissionRouters(permissionIds []uint64, traceID string) (routerUrls []string) {
	// 获取权限列表对应的路由列表（去重）
	routerIds, _ := (&platDb.PermissionRouter{}).FindPermissionRouterByIds(permissionIds)
	// 检查路由是否启用
	if len(routerIds) > 0 {
		// 读取路由列表.FindByIds(routerIds)
		routers, _ := (&platDb.Router{}).FindByIds(routerIds)
		// 可用路由地址
		useUrls := make([]string, 0)
		// 判定角色
		if len(routers) > 0 {
			for _, router := range routers {
				useUrls = append(useUrls, router.Url)
			}
		}
		// 赋值
		routerUrls = useUrls
	}
	// 警告
	if len(permissionIds) < 1 {
		log.WarnTF(traceID, "GetPermissionRouters PermissionIds %V RouterList Is Empty .", permissionIds)
	}
	return
}

// 移除路由绑定的权限关系，涉及权限异步反馈（管理中台暂不考虑事务）
func notifyChangeByRouter(id uint64, remove bool, traceID string) (err error) {
	// 获取路由绑定权限ID
	permissions, err := platDb.PermissionRouterTable.FindByObject(&platDb.PermissionRouter{RouterId: id})
	if err != nil {
		log.ErrorTF(traceID, "FindPermission By RouterId %d Fail . Err Is : %v", id, err)
		return
	}
	// 遍历权限，用于回溯影响的账号
	if len(permissions) > 0 {
		dataIds := make([]uint64, len(permissions))
		permissionIds := make([]uint64, len(permissions))
		for i, item := range permissions {
			permissionIds[i] = item.PermissionId
			dataIds[i] = item.ID
		}
		// 如果需要移除关系（通过ID移除）
		if remove {
			err = platDb.PermissionRouterTable.DeleteByIds(dataIds)
			log.ErrorTF(traceID, "DeleteRouterPermission By RouterId %d Fail . Err Is : %v", id, err)
		}
		// 异步通知受影响的Token失效（上层不涉及数据删除等行为）
		go notifyChangeByPermissionIds(permissionIds, false, traceID)
	}
	return
}
