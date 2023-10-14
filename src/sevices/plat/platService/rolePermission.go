package platService

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/utils/log"
)

// 获取角色权限列表
func getRolePermissions(roleIds []uint64, traceID string) (permissionIds []uint64, permissionAlias []string) {
	// 获取角色列表对应的权限列表（去重）
	permissionIds, _ = platDb.RolePermission{}.FindRolePermissionByIds(roleIds)
	// 检查权限是否启用
	if len(permissionIds) > 0 {
		// 读取权限列表（仅查询可用）.FindByIds(permissionIds)
		permissions, _ := platDb.Permission{}.FindByIds(permissionIds)
		// 可用权限ID
		useIds := make([]uint64, 0)
		// 可用权限别名
		useAlias := make([]string, 0)
		// 判定角色
		if len(permissions) > 0 {
			for _, permission := range permissions {
				if permission.Status == constant.StatusOpen {
					useIds = append(useIds, permission.ID)
					useAlias = append(useAlias, permission.Alias)
				}
			}
		}
		// 赋值
		permissionIds = useIds
		permissionAlias = useAlias
	}
	// 警告
	if len(permissionIds) < 1 {
		log.WarnTF(traceID, "GetRolePermissions RoleIds %V PermissionList Is Empty .", roleIds)
	}
	return
}

// 刷新角色的权限列表
func refreshRolePermissions(req *platDb.Role, needDel bool, traceID string) (err error) {
	if needDel {
		err = platDb.RolePermission{}.DeleteByRoleId(req.ID)
		if err != nil {
			log.ErrorTF(traceID, "DeleteByRoleId By %d Fail . Err Is : %v", req.ID, err)
			return
		}
	}
	// 重新插入权限关系
	permissionIds := req.PermissionIds
	if len(permissionIds) > 0 {
		rolePermissions := make([]platDb.RolePermission, len(permissionIds))
		for i, item := range permissionIds {
			rolePermissions[i] = platDb.RolePermission{
				RoleId:       req.ID,
				PermissionId: item,
			}
		}
		err = platDb.RolePermissionTable.InsertBatch(&rolePermissions)
		if err != nil {
			log.ErrorTF(traceID, "InsertBatchRolePermission By RoleId %d Fail . Err Is : %v", req.ID, err)
		}
	}
	return
}
