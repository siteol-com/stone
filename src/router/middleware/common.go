package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/utils/log"
)

// 获取语言并设置
func setLang(c *gin.Context) {
	lang := c.GetHeader(constant.HeaderLang)
	if lang == "" {
		lang = "zh-CN"
	}
	c.Set(constant.HeaderLang, lang)
}

// 处理业务响应
func returnJSON(c *gin.Context, middlewareName, traceID string) {
	// 响应读取 读取失败响应系统异常
	respBody, ok := c.Get(constant.RespBody)
	if !ok {
		log.ErrorTF(traceID, "Resp Get Fail")
		respBody = resp.SysErr
	}
	// 处理响应翻译
	returnMsgTrans(respBody, c, middlewareName, traceID)
}

// 执行响应码 => 响应文言 翻译
func returnMsgTrans(respBody interface{}, c *gin.Context, middlewareName, traceID string) {
	// 类型回转
	res, ok := respBody.(resp.ResBody)
	if !ok {
		log.ErrorTF(traceID, "ResBody Type UnKnow")
		res = resp.SysErr
	} else {
		// 非400错误执行翻译
		if res.Code != constant.RespValidateErrCode {
			// 执行翻译
			res.Msg = runMsgTrans(res, c.GetString(constant.HeaderLang), c.GetString(constant.TraceID))
		}
	}
	// JSON序列化
	resByte, _ := json.Marshal(res)
	log.InfoTF(traceID, "%s End. RespBody: %s", middlewareName, resByte)
	// 响应结果
	c.JSON(http.StatusOK, res)
}
