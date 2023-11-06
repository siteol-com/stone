package subRouter

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/sevices/plat/platHandler"
	"siteOl.com/stone/server/src/sevices/router/middleware"
)

// OpenRouter Open 开放路由 通过数据库开放
func OpenRouter(router *gin.Engine) {
	PlatFormRouter := router.Group("/open", middleware.AuthMiddleWare)
	{
		// 开放账密登陆
		PlatFormRouter.POST("/auth/login", platHandler.AuthLogin)
		// 开放租户信息获取
		PlatFormRouter.POST("/tenant/get", platHandler.GetOpenTenant)
	}
}
