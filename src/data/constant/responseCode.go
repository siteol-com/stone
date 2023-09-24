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
	Success       = "200" // 默认成功文言（内置禁止修改）
	SysFail       = "500" // 默认系统错误（内置禁止修改）
	ValidateFail  = "400" // 参数校验错误（内置禁止修改）
	LoginFail     = "401" // 默认登陆错误（内置禁止修改）
	AuthFail      = "403" // 默认授权错误（内置禁止修改）
	AuthResetFail = "405" // 默认授权刷新（内置禁止修改）

	TenantGetOK    = "2001000" // 租户信息获取成功
	TenantGetNG    = "5001000" // 租户信息获取失败
	TenantStatusNG = "5001001" // 该租户暂不可用
	TenantExpNG    = "5001002" // 该租户已过期

	AccountLoginOK  = "2002000" // 账号登陆成功
	AccountLoginNG  = "5002000" // 账号或密码错误
	AccountStatusNG = "5002001" // 账号暂不可用

	RouteAddOK     = "2003000" // 路由创建成功
	RouteEditOK    = "2003001" // 路由更新成功
	RouteDelOK     = "2003002" // 路由删除成功
	RouteGetNG     = "5003000" // 路由查询失败
	RouteUniUrlNG  = "5003001" // 路由地址不可重复
	RouteUniNameNG = "5003002" // 路由名称不可重复

	ResponseAddOK     = "2005000" // 响应码创建成功
	ResponseEditOK    = "2005001" // 响应码编辑成功
	ResponseDelOK     = "2005002" // 响应码删除成功
	ResponseGetNG     = "5005000" // 响应码查询失败
	ResponseUniCodeNG = "5005001" // 响应码不可重复

	PermissionAddOK       = "2006000" // 权限创建成功
	PermissionEditOK      = "2006001" // 权限编辑成功
	PermissionDelOK       = "2006002" // 权限删除成功
	PermissionSortOK      = "2006003" // 权限排序成功
	PermissionGetNG       = "5006000" // 权限查询失败
	PermissionUniNameNG   = "5006001" // 权限名不可重复
	PermissionUniAliasNG  = "5006002" // 权限别名不可重复
	PermissionUniRouterNG = "5006003" // 权限关联路由不可重复
	PermissionDelChildNG  = "5006004" // 权限存在子集不可删除
	PermissionSortNG      = "5006005" // 权限排序失败

)
