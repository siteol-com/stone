package constant

const (
	ProjectName = "Stone"    // 项目名
	HeaderLang  = "Lang"     // 固定请求头（语言）
	HeaderToken = "Token"    // 固定请求头（登陆Token）
	RespBody    = "respBody" // 响应报文KEY
	TraceID     = "traceID"  // 日志链路跟踪ID
	AuthUser    = "authUser" // 授权用户对象

	RespTypeJSON        = "application/json; charset=utf-8" // 固定响应格式，默认JSON
	RespValidateErrCode = "400"                             // 固定400校验错误码

	StatusOpen  = "0" // 正常 启动
	StatusLock  = "1" // 禁用 锁定 登出
	StatusClose = "2" // 移除 弃用 踢出

	TimeNormal = "2006-01-02 15:04:05" // 常见时间格式
	TimeNumber = "20060102150405"      // 存数字时间格式

	LoginTypeAuth = "1" // 账号授权登陆

	PermissionTypeAll        = "0"  // 全局数据
	PermissionTypeDeptFellow = "1"  // 跟随部门
	PermissionTypeDeptThis   = "2"  // 当前部门
	PermissionTypeDeptGroup  = "3"  // 当前部门与子部门
	PermissionTypeNull       = "99" // 无数据权限

	AuthCacheSecond = 900 // 授权缓存900秒默认15分钟

	DBDuplicateErr = "Error 1062 (23000): Duplicate entry" // 唯一索引错误

	RouterTypeAuth  = "1" // 授权路由
	RouterTypeWhite = "2" // 白名单路由

	CacheKeyRouterWhite = ProjectName + "::WhiteRouterList"
	CacheKeyTransLang   = ProjectName + "::TranLangMap" // 响应码翻译缓存Map的Key
)

// TransLangSupport 支持更多语言请添加
var TransLangSupport = []string{"zh-CN", "en-US"}
