package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/model/platModel"
	"siteOl.com/stone/server/src/sevices"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

// GetOpenTenant godoc
// @id			 GetOpenTenant获取租户信息
// @Summary      获取租户信息
// @Description  获取租户基础信息，前置开放接口
// @Router       /open/tenant/get [post]
// @Tags         开放接口
// @Accept       json
// @Produce      json
// @Param        req body platModel.OpenTenantReq true "请求"
// @Success      200 {object} resp.ResBody{data=platModel.OpenTenantRes} "响应成功"
func GetOpenTenant(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.OpenTenantReq{})
	if err == nil {
		// 执行查询
		sevices.JsonRes(c, platService.GetOpenTenant(traceID, req.(*platModel.OpenTenantReq)))
	}
}
