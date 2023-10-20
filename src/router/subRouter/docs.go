package subRouter

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/sevices/plat/platHandler"
)

// DocsRouter API文档路由
func DocsRouter(router *gin.Engine) {
	// API
	// initSwagger
	docsRouter := router.Group("/docs")
	{
		// Swagger资源文件
		docsRouter.GET("/sc/*any", platHandler.ScFile)
		// ReDoc
		docsRouter.GET("/redoc/*any", platHandler.ReDoc)
		// Swagger范本
		docsRouter.GET("/swagger/*any", platHandler.SwaggerDoc)
	}
}
