package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/sevices"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

// AuthLogin godoc
// @id			 AuthLogin开放账密登陆
// @Summary      开放账密登陆
// @Description  平台最基础的账号密码登陆方式登陆
// @Tags         开放接口
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.AuthLoginReq true "请求"
// @Success      200 {object} resp.ResBody{data=platModel.AuthLoginRes} "登陆成功"
// @Failure      400 {object} resp.ResBody "数据校验失败"
// @Failure      500 {object} resp.ResBody "登陆失败相关"
// @Router       /open/auth/login [post]
func AuthLogin(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.AuthLoginReq{})
	if err != nil {
		return
	}
	// 执行登陆
	sevices.JsonRes(c, platService.AuthLogin(traceID, req.(*platModel.AuthLoginReq)))
}
