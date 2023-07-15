package subRouter

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/router/middleware"
	"siteOl.com/stone/server/src/sevices/plat/platHandler"
)

// PlatFormRouter 平台业务路由
func PlatFormRouter(router *gin.Engine) {
	platRouter := router.Group("/plat", middleware.AuthMiddleWare)
	{
		// 账号相关
		accountRouter := platRouter.Group("/account")
		{
			accountRouter.POST("/add", platHandler.AddAccount)
		}

		// 路由相关
		routerRouter := platRouter.Group("/router")
		{
			routerRouter.POST("/page", platHandler.PageRouter)
			routerRouter.POST("/add", platHandler.AddRouter)
			routerRouter.POST("/get", platHandler.GetRouter)
			routerRouter.POST("/edit", platHandler.EditRouter)
			routerRouter.POST("/del", platHandler.DelRouter)
		}

		// 字典相关
		dictRouter := platRouter.Group("/dict")
		{
			dictRouter.POST("/list", platHandler.ListDict)
		}
	}
}
