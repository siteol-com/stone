package resp

import (
	"net/http"
	"siteOl.com/stone/server/src/data/constant"
)

/**
 *
 * 统一数据JSON返回结构
 * 200 业务处理完成（可能存在业务异常）
 * 400 请求数据不合法（校验失败）
 * 401 未授权、未登陆
 * 403 禁止访问（无权限）
 * 500 系统未知异常（意料之外的错误）
 *
 * @author 米虫丨www.mebugs.com
 * @since 2022-08-16
 */

// 定义一些常量
var (
	// ResOk 200 默认业务成功
	ResOk = SuccessUnPop(nil)
	// ResFail 200 默认业务失败
	ResFail = Fail(constant.SysFail)
	// SysErr Json 500 默认系统异常
	SysErr = jsonResult(http.StatusInternalServerError, constant.SysErr, "", nil, false)
)

type ResBody struct {
	HttpCode int    `json:"-"`
	Code     string `json:"code" example:"S0000/F0000"` // 响应码
	Msg      string `json:"msg" example:"操作成功/失败"`      // 响应消息
	Data     any    `json:"data"`                       // 响应数据
	UnPop    bool   `json:"unPop" example:"true"`       // 免弹窗提示
}

// SuccessUnPop 200 成功但提示前端不弹Pop（一般用于分页查询、详情查询）
func SuccessUnPop(data any) *ResBody {
	return jsonResult(http.StatusOK, constant.Success, "", data, true)
}

// SuccessWithCode 200 成功携带自定义响应码
func SuccessWithCode(code string, data any) *ResBody {
	return jsonResult(http.StatusOK, code, "", data, false)
}

// Fail 200 业务失败携带自定义响应码
func Fail(code string) *ResBody {
	return jsonResult(http.StatusOK, code, "", nil, false)
}

// Validate 400 Json校验失败
func Validate(err error) *ResBody {
	return jsonResult(http.StatusBadRequest, constant.ValidErr, err.Error(), nil, false)
}

// 公共调用
func jsonResult(httpCode int, code string, msg string, data any, unPop bool) *ResBody {
	resp := &ResBody{
		HttpCode: httpCode,
		Code:     code,
		Msg:      msg,
		Data:     data,
		UnPop:    unPop,
	}
	return resp
}
