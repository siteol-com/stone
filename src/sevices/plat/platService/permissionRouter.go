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
