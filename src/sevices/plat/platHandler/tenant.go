package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/sevices/plat"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

//  登陆授权

// GetOpenTenant /open/tenant/get 获取租户数据信息（开放）
func GetOpenTenant(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platModel.OpenTenantReq{})
	if err != nil {
		return
	}
	// 执行查询
	res := platService.GetOpenTenant(traceID, req.(*platModel.OpenTenantReq))
	c.Set(constant.RespBody, res)
}
