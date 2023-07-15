package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/sevices/plat"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

//  登陆授权

// AuthLogin /open/auth/login 开放账密登陆
func AuthLogin(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platModel.AuthLoginReq{})
	if err != nil {
		return
	}
	// 执行登陆
	res := platService.AuthLogin(traceID, req.(*platModel.AuthLoginReq))
	c.Set(constant.RespBody, res)
}
