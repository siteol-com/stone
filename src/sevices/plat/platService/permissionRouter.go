package platService

import (
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/utils/log"
)

// 获取权限路由地址列表
func getPermissionRoutersUrls(permissionIds []uint64, traceID string) (routerUrls []string) {
	// 可用路由地址
	routerUrls = make([]string, 0)
	routers, _ := getPermissionRouters(permissionIds, traceID)
	// 判定角色
	if len(routers) > 0 {
		for _, router := range routers {
			routerUrls = append(routerUrls, router.Url)
		}
	}
	// 警告
	if len(routerUrls) < 1 {
		log.WarnTF(traceID, "GetPermissionRouters PermissionIds %V RouterList Is Empty .", permissionIds)
	}
	return
}

// 获取权限路由列表
func getPermissionRouters(permissionIds []uint64, traceID string) (routers []*platDb.Router, err error) {
	// 获取权限列表对应的路由列表（去重）
	routerIds, err := platDb.PermissionRouter{}.FindPermissionRouterByIds(permissionIds)
	if err != nil {
		log.ErrorTF(traceID, "FindPermissionRouterByIds By %v Fail . Err Is : %v", permissionIds, err)
		return
	}
	// 检查路由是否启用
	if len(routerIds) > 0 {
		// 读取路由列表.FindByIds(routerIds)
		routers, err = (&platDb.Router{}).FindByIds(routerIds)
		if err != nil {
			log.ErrorTF(traceID, "FindRouterByIds By %v Fail . Err Is : %v", routerIds, err)
		}
	}
	return
}

// refreshPermissionRouters 添加路由
func refreshPermissionRouters(req *platDb.Permission, needDel bool, traceID string) (err error) {
	if needDel {
		err = platDb.PermissionRouter{}.DeleteByPermissionId(req.ID)
		if err != nil {
			log.ErrorTF(traceID, "DeleteByPermissionId By %d Fail . Err Is : %v", req.ID, err)
			return
		}
	}
	// 重新插入路由关系
	routerIds := req.RouterIds
	if len(routerIds) > 0 {
		permissionRouters := make([]platDb.PermissionRouter, len(routerIds))
		for i, item := range routerIds {
			permissionRouters[i] = platDb.PermissionRouter{
				PermissionId: req.ID,
				RouterId:     item,
			}
		}
		err = platDb.PermissionRouterTable.InsertBatch(&permissionRouters)
		if err != nil {
			log.ErrorTF(traceID, "InsertBatchPermissionRouter By PermissionId %d Fail . Err Is : %v", req.ID, err)
		}
	}
	return
}
