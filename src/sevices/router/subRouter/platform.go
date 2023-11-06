package subRouter

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/sevices/plat/platHandler"
	"siteOl.com/stone/server/src/sevices/router/middleware"
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
		// 角色相关
		routerRole := platRouter.Group("/role")
		{
			routerRole.POST("/page", platHandler.PageRole)
			routerRole.POST("/add", platHandler.AddRole)
			routerRole.POST("/get", platHandler.GetRole)
			routerRole.POST("/edit", platHandler.EditRole)
			routerRole.POST("/del", platHandler.DelRole)
		}
		// 权限相关
		permissionRouter := platRouter.Group("/permission")
		{
			permissionRouter.POST("/tree", platHandler.TreePermission)
			permissionRouter.POST("/add", platHandler.AddPermission)
			permissionRouter.POST("/get", platHandler.GetPermission)
			permissionRouter.POST("/edit", platHandler.EditPermission)
			permissionRouter.POST("/del", platHandler.DelPermission)
			permissionRouter.POST("/bro", platHandler.BroPermission)
			permissionRouter.POST("/sort", platHandler.SortPermission)
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
		// 响应码相关
		responseRouter := platRouter.Group("/response")
		{
			responseRouter.POST("/page", platHandler.PageResponse)
			responseRouter.POST("/add", platHandler.AddResponse)
			responseRouter.POST("/get", platHandler.GetResponse)
			responseRouter.POST("/edit", platHandler.EditResponse)
			responseRouter.POST("/del", platHandler.DelResponse)
		}
		// 字典相关
		dictRouter := platRouter.Group("/dict")
		{
			dictRouter.POST("/list", platHandler.ListDict)
		}
	}
}
