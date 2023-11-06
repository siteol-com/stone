package resp

// DemoOk 200
type DemoOk struct {
	Code  string `json:"code" example:"SX00X/FX00X"` // 响应码
	Msg   string `json:"msg" example:"操作成功/失败"`      // 响应消息
	Data  string `json:"data" example:"响应数据"`        // 响应数据
	UnPop bool   `json:"unPop" example:"true"`       // 免弹窗提示
}

// DemoErr 500
type DemoErr struct {
	Code string `json:"code" example:"E000"` // 响应码
	Msg  string `json:"msg" example:"系统异常"`  // 响应消息
}

// DemoVail 400
type DemoVail struct {
	Code string `json:"code" example:"E001"`     // 响应码
	Msg  string `json:"msg" example:"xx字段应该为必填"` // 响应消息
}

// DemoAuthLg 401
type DemoAuthLg struct {
	Code string `json:"code" example:"E002"`  // 响应码
	Msg  string `json:"msg" example:"当前尚未登陆"` // 响应消息
}

// DemoAuthNg 403
type DemoAuthNg struct {
	Code string `json:"code" example:"E003"`    // 响应码
	Msg  string `json:"msg" example:"当前接口禁止访问"` // 响应消息
}
