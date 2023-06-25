package resp

/**
 *
 * 统一数据JSON返回结构
 *
 *
 * @author 米虫@mebugs.com
 * @since 2022-08-16
 */

// 定义一些常量
var (
	// OK Json默认成功返回
	OK = Success(nil)
	// SysErr Json默认系统异常
	SysErr = Error()
)

type ResBody struct {
	Code string      `json:"code"` // 响应码
	Msg  string      `json:"msg"'` // 响应消息
	Data interface{} `json:"data"` // 响应数据
}

// Success Json数据返回
func Success(data interface{}) ResBody {
	return jsonResult("200", "", data)
}

// SuccessWithCode Json数据返回
func SuccessWithCode(code string, data interface{}) ResBody {
	return jsonResult(code, "", data)
}

// Validate Json校验返回400（已翻译）
func Validate(err error) ResBody {
	return jsonResult("400", err.Error(), nil)
}

// Error Json错误返回500
func Error() ResBody {
	return jsonResult("500", "", nil)
}

// Fail 业务失败数据（无响应体）
func Fail(code string) ResBody {
	return jsonResult(code, "", nil)
}

// 公共调用
func jsonResult(code string, msg string, data interface{}) ResBody {
	resp := ResBody{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	return resp
}
