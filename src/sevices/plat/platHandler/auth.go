package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/model/platModel"
	"siteOl.com/stone/server/src/sevices"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

// AuthLogin godoc
// @id			 AuthLogin开放账密登陆
// @Summary      开放账密登陆
// @Description  平台最基础的账号密码登陆方式登陆
// @Router       /open/auth/login [post]
// @Tags         开放接口
// @Accept       json
// @Produce      json
// @Param        req body platModel.AuthLoginReq true "请求"
// @Success      200 {object} resp.ResBody{data=platModel.AuthLoginRes} "登陆成功"
func AuthLogin(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.AuthLoginReq{})
	if err == nil {
		// 执行登陆
		sevices.JsonRes(c, platService.AuthLogin(traceID, req.(*platModel.AuthLoginReq)))
	}
}
