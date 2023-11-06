package constant

/**
 *
 * 响应码常量
 * 具体文言维护在数据库，但编码需要在此处维护，提高代码可读性
 *
 * @author 米虫丨www.mebugs.com
 * @since 2023-07-21
 */

const (
	Success  = "S0000" // 默认成功文言（内置禁止修改）
	SysFail  = "F0000" // 默认业务错误（内置禁止修改）
	SysErr   = "E0000" // 系统未知错误（内置禁止修改）
	ValidErr = "E0001" // 参数校验错误（内置禁止修改）（免翻译）
	LoginErr = "E0002" // 默认登陆错误（内置禁止修改）
	AuthErr  = "E0003" // 默认授权错误（内置禁止修改）
	ResetErr = "E0004" // 默认授权刷新（内置禁止修改）

	TenantGetOK    = "S1000" // 租户信息获取成功
	TenantGetNG    = "F1000" // 租户信息获取失败
	TenantStatusNG = "F1001" // 该租户暂不可用
	TenantExpNG    = "F1002" // 该租户已过期

	AccountLoginOK  = "S2000" // 账号登陆成功
	AccountLoginNG  = "F2000" // 账号或密码错误
	AccountStatusNG = "F2001" // 账号暂不可用

	RouteAddOK     = "S3000" // 路由创建成功
	RouteEditOK    = "S3001" // 路由更新成功
	RouteDelOK     = "S3002" // 路由删除成功
	RouteGetNG     = "F3000" // 路由查询失败
	RouteUniUrlNG  = "F3001" // 路由地址不可重复
	RouteUniNameNG = "F3002" // 路由名称不可重复

	ResponseAddOK     = "S5000" // 响应码创建成功
	ResponseEditOK    = "S5001" // 响应码编辑成功
	ResponseDelOK     = "S5002" // 响应码删除成功
	ResponseGetNG     = "F5000" // 响应码查询失败
	ResponseUniCodeNG = "F5001" // 响应码不可重复

	PermissionAddOK       = "S6000" // 权限创建成功
	PermissionEditOK      = "S6001" // 权限编辑成功
	PermissionDelOK       = "S6002" // 权限删除成功
	PermissionSortOK      = "S6003" // 权限排序成功
	PermissionGetNG       = "F6000" // 权限查询失败
	PermissionUniNameNG   = "F6001" // 权限名不可重复
	PermissionUniAliasNG  = "F6002" // 权限别名不可重复
	PermissionUniRouterNG = "F6003" // 权限关联路由不可重复
	PermissionDelChildNG  = "F6004" // 权限存在子集不可删除
	PermissionSortNG      = "F6005" // 权限排序失败

	RoleAddOK           = "S7000" // 角色创建成功
	RoleEditOK          = "S7001" // 角色编辑成功
	RoleDelOK           = "S7002" // 角色删除成功
	RoleGetNG           = "F7000" // 角色查询失败
	RoleUniNameNG       = "F7001" // 角色名不可重复
	RoleUniPermissionNG = "F7002" // 角色权限不可重复
	RoleEditLockNG      = "F7003" // 该角色不可编辑
	RoleDelLockNG       = "F7004" // 该角色不可删除
)
