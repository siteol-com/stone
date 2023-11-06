package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/redis"
	"siteOl.com/stone/server/src/utils/log"
)

// 获取路由配置
func getRouter(url, traceID string) (router *platDb.Router) {
	// 默认路由配置
	router = &platDb.Router{
		Url:      url,
		Type:     constant.RouterTypeAuth, // 默认授权
		PrintReq: constant.RouterLogPrint, // 默认打印
		PrintRes: constant.RouterLogPrint, // 默认打印
	}
	cache, err := redis.Get(constant.CacheKeyRouterMap)
	if err != nil {
		log.WarnTF(traceID, "Get Router[%s] Cache Fail. Err: %s", url, err)
		return
	}
	routerMap := make(map[string]*platDb.Router, 0)
	err = json.Unmarshal([]byte(cache), &routerMap)
	if err != nil {
		log.WarnTF(traceID, "Json Unmarshal Router[%s] Cache Fail. Err: %s", url, err)
		return
	}
	// 获取缓存设定
	cacheRouter, ok := routerMap[url]
	if ok {
		return cacheRouter
	}
	return
}

// 获取语言并设置
func setLang(c *gin.Context) {
	lang := c.GetHeader(constant.HeaderLang)
	if lang == "" || lang == "null" {
		lang = "zh-CN"
	}
	c.Set(constant.HeaderLang, lang)
}

// 记录请求
func readReq(c *gin.Context, router *platDb.Router, traceID string) error {
	if c.Request.Method == http.MethodGet {
		// 收集日志
		return nil
	}
	printBts := []byte("{}")
	if router.PrintReq == constant.RouterLogPrintNot {
		printBts = []byte("{ Req Set Not Print }")
	} else {
		bodyBts, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.ErrorTF(traceID, "ReqBody Read Fail: %s", err)
			return err
		}
		// 写回body
		bodyGo := ioutil.NopCloser(bytes.NewBuffer(bodyBts))
		c.Request.Body = bodyGo
		printBts = bodyBts
	}

	log.InfoTF(traceID, "ReqBody: %s", printBts)
	return nil
}
