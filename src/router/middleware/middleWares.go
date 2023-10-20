package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/utils/comm"
	"siteOl.com/stone/server/src/utils/log"
)

// AuthMiddleWare 鉴权中间件
func AuthMiddleWare(c *gin.Context) {
	// 生成请求唯一标志
	traceID := comm.TraceID()
	log.InfoTF(traceID, "AuthMiddleWare URL = %s", c.Request.URL.Path)
	c.Set(constant.TraceID, traceID)
	// 获取路由配置
	router := getRouter(c.Request.URL.Path, "AuthMiddleWare", traceID)
	c.Set(constant.RouterConf, router)
	// 设置语言
	setLang(c)
	// 是否是中间件拒绝
	middleRes := true
	defer func() {
		// 中间件拒绝的特殊处理
		if middleRes {
			c.JSON(http.StatusInternalServerError, resp.SysErr)
			// TODO  翻译
		}
		returnJSON(c, router, "AuthMiddleWare", traceID)
	}()
	// 读取请求
	if readReq(c, router, "AuthMiddleWare", traceID) != nil {
		return
	}
	// 读取鉴权信息
	// TODO
	c.Next()
	// 控制层响应
	middleRes = true
}
