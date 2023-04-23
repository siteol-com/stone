package constant

const (
	HeaderLang  = "Lang"     // 固定请求头（语言）
	HeaderToken = "Token"    // 固定请求头（登陆Token）
	RespBody    = "respBody" // 响应报文KEY
	TraceID     = "traceID"  // 日志链路跟踪ID
	AuthUser    = "authUser" // 授权用户对象

	RespTypeJSON        = "application/json; charset=utf-8" // 固定响应格式，默认JSON
	RespValidateErrCode = "400"                             // 固定400校验错误码

	StatusOpen  = 1 // 正常 启动
	StatusLock  = 2 // 禁用 锁定
	StatusClose = 3 // 移除 弃用
)
