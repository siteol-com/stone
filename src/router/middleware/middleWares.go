package middleware

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/utils/comm"
	"siteOl.com/stone/server/src/utils/log"
)

// AuthMiddleWare 鉴权中间件
func AuthMiddleWare(c *gin.Context) {
	// 生成请求唯一标志
	traceID := comm.TraceID()
	log.InfoTF(traceID, "AuthMiddleWare URL = %s", c.Request.URL.Path)
	// 获取路由配置
	router := getRouter(c.Request.URL.Path, "AuthMiddleWare", traceID)
	defer func() {
		// 退出前的追加处理
		returnJSON(c, router, "AuthMiddleWare", traceID)
	}()
	c.Set(constant.TraceID, traceID)
	// 设置语言
	setLang(c)
	// 读取请求
	if readReq(c, router, "AuthMiddleWare", traceID) != nil {
		return
	}
	// 读取鉴权信息
	// TODO
	c.Next()
}
