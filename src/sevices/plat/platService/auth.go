package platService

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/redis"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/utils/comm"
	"siteOl.com/stone/server/src/utils/log"
	"siteOl.com/stone/server/src/utils/security"
	"time"
)

// AuthLogin 账号登录
func AuthLogin(traceID string, req *platModel.AuthLoginReq) resp.ResBody {
	// 查询租户
	tenant, err := platDb.TenantTable.FindOneByObject(&platDb.Tenant{Alias: req.TenantAlias})
	if err != nil {
		log.ErrorTF(traceID, "AuthLogin GetTenant Fail . Err is %v", err)
		return resp.Fail("5001000") // 租户查询失败
	}
	// 检查租户，检查不通过
	check, checkRes := CheckTenant(&tenant)
	if !check {
		return checkRes
	}
	// 查询账号
	account, err := platDb.AccountTable.FindOneByObject(&platDb.Account{Account: req.Account, TenantId: tenant.ID})
	if err != nil {
		log.ErrorTF(traceID, "AuthLogin GetAccount Fail . Err is %v", err)
		return resp.Fail("5002000") // 混淆错误：账号或密码错误
	}
	// 检查账号，检查不通过
	check, checkRes = checkAccount(&account)
	if !check {
		return checkRes
	}
	// 验证密码
	encryptionReq, _ := security.AESEncrypt(req.Password, account.SaltKey)
	if encryptionReq != account.Encryption {
		log.ErrorTF(traceID, "AuthLogin PasswordEncryption Wrong")
		// TODO 密码错误上限，锁定时间
		return resp.Fail("5002000") // 混淆错误：账号或密码错误
	}
	// 生成Token 时间戳+随机长度 = 32登陆Token
	now := time.Now()
	token := now.Format(constant.TimeNumber) + comm.RandStr(18)
	// 创建登陆记录
	InsertLoginRecord(account.ID, tenant.ID, constant.LoginTypeAuth, now, token, traceID)
	// TODO 踢出超限的登陆数据
	// 初始化登陆权限（塞入Redis）
	setAuthUser(&account, token, traceID)
	// 账号登陆成功
	return resp.SuccessWithCode("2002000", platModel.AuthLoginRes{Token: token})
}

// 设置登陆授权数据 （失败不影响登陆过）
func setAuthUser(account *platDb.Account, token, traceID string) {
	// 初始化授权结构体
	authUser := platModel.AuthUser{
		UserId:         account.ID,
		TenantId:       account.TenantId,
		DeptId:         account.DeptId,
		PermissionType: account.PermissionType,
	}
	// 设置过期时间
	if account.PwdExpTime != nil {
		authUser.PwdExpTimeStr = account.PwdExpTime.Format(constant.TimeNormal)
	}
	// 读取权限集 读取路由集
	aliasList, routers := getUserPermissionRouters(account.ID, traceID)
	authUser.PermissionList = aliasList
	authUser.RouterList = routers
	// 数据权限判定（部门联动）如果不是配置的全局数据，可能需要根据部门联动
	if authUser.PermissionType != constant.PermissionTypeAll {
		permissionType, deptIds := getUserDataPermission(authUser.DeptId, authUser.PermissionType, traceID)
		authUser.PermissionType = permissionType
		authUser.PermissionDeptList = deptIds
	}
	// 放入Redis 默认15分钟 TODO 过期时间取配置
	redis.Set(token, authUser, constant.AuthCacheSecond)
}

// 读取账号权限集
func getUserPermissionRouters(accountId uint64, traceID string) (aliasList []string, routers []string) {
	// 读取账号角色
	roleIds := getAccountRoleIds(accountId, traceID)
	if len(roleIds) < 1 {
		return
	}
	// 读取角色权限ID和别名
	permissionIds, aliasList := getRolePermissions(roleIds, traceID)
	// 没有权限
	if len(permissionIds) < 1 {
		return
	}
	// 根据权限获取路由集
	routers = getPermissionRouters(permissionIds, traceID)
	return
}
