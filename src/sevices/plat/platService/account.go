package platService

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/resp"
)

// 检查账号状态
func checkAccount(account *platDb.Account) (check bool, res *resp.ResBody) {
	// 账号状态不正确
	if account.Status != constant.StatusOpen {
		res = resp.Fail(constant.AccountStatusNG) // 账号状态不可用
		return
	}
	// TODO 密码超期提示

	check = true
	return
}
