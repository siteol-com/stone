package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/validate"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

//  登陆授权

// GetOpenTenant /open/tenant/get 获取租户数据信息（开放）
func GetOpenTenant(c *gin.Context) {
	// TraceID 日志追踪
	traceID := c.GetString(constant.TraceID)
	// 参数读取
	req := &platModel.OpenTenantReq{}
	// 校验并且 解析请求数据
	err, reqObj := validate.Readable(c, req)
	if err != nil {
		return
	}
	req = reqObj.(*platModel.OpenTenantReq)
	// 执行查询
	res := platService.GetOpenTenant(traceID, req)
	c.Set(constant.RespBody, res)
}
