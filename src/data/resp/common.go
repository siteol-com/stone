package resp

/**
 *
 * 统一数据JSON返回结构
 *
 *
 * @author 米虫丨www.mebugs.com
 * @since 2022-08-16
 */

// 定义一些常量
var (
	// OK Json默认成功返回
	OK      = Success(nil)
	OKUnPop = SuccessUnPop(nil)
	// SysErr Json默认系统异常
	SysErr = Error()
)

type ResBody struct {
	Code  string `json:"code"`  // 响应码
	Msg   string `json:"msg"'`  // 响应消息
	Data  any    `json:"data"`  // 响应数据
	UnPop bool   `json:"unPop"` // 免弹窗提示
}

// Success Json数据返回
func Success(data any) ResBody {
	return jsonResult("200", "", data, false)
}

// SuccessUnPop 成功但提示前端不弹Pop
func SuccessUnPop(data any) ResBody {
	return jsonResult("200", "", data, true)
}

// SuccessWithCode Json数据返回
func SuccessWithCode(code string, data any) ResBody {
	return jsonResult(code, "", data, false)
}

// Validate Json校验返回400（已翻译）
func Validate(err error) ResBody {
	return jsonResult("400", err.Error(), nil, false)
}

// Error Json错误返回500
func Error() ResBody {
	return jsonResult("500", "", nil, false)
}

// Fail 业务失败数据（无响应体）
func Fail(code string) ResBody {
	return jsonResult(code, "", nil, false)
}

// 公共调用
func jsonResult(code string, msg string, data any, unPop bool) ResBody {
	resp := ResBody{
		Code:  code,
		Msg:   msg,
		Data:  data,
		UnPop: unPop,
	}
	return resp
}
