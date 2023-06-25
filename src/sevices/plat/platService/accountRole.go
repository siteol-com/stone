package platService

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/utils/log"
)

// 获取账号对应的角色列表
func getAccountRoleIds(accountId uint64, traceID string) (roleIds []uint64) {
	roleIds, _ = (&platDb.AccountRole{AccountId: accountId}).FindAccountRoleIds()
	// 没有角色
	if len(roleIds) > 0 {
		// 读取角色列表（仅查询可用）.FindByIds(roleIds)
		roles, _ := (&platDb.Role{}).FindByIds(roleIds)
		// 可用角色
		useIds := make([]uint64, 0)
		// 判定角色
		if len(roles) > 0 {
			for _, role := range roles {
				if role.Status == constant.StatusOpen {
					useIds = append(useIds, role.ID)
				}
			}
		}
		// 赋值
		roleIds = useIds
	}
	// 警告
	if len(roleIds) < 1 {
		log.WarnTF(traceID, "GetAccountRoleIds AccountID %d RoleList Is Empty .", accountId)
	}
	return
}
