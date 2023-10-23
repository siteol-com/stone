package router

import (
	"io/ioutil"
	"siteOl.com/stone/server/src/router/middleware"
	"siteOl.com/stone/server/src/router/subRouter"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	gin.ForceConsoleColor() // 颜色日志
	// 基础路由
	router := gin.Default()
	// 公共的Panic中间件
	router.Use(middleware.Recover)
	// API文档（示例文档）
	subRouter.DocsRouter(router)
	// 开放路由
	subRouter.OpenRouter(router)
	// 平台路由
	subRouter.PlatFormRouter(router)
	return router
}
