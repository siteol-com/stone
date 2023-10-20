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
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/utils/log"
)

// 获取路由配置
func getRouter(url, middlewareName, traceID string) (router *platDb.Router) {
	// 默认路由配置
	router = &platDb.Router{
		Url:      url,
		Type:     constant.RouterTypeAuth, // 默认授权
		PrintReq: constant.RouterLogPrint, // 默认打印
		PrintRes: constant.RouterLogPrint, // 默认打印
	}
	cache, err := redis.Get(constant.CacheKeyRouterMap)
	if err != nil {
		log.WarnTF(traceID, "%s Get Router[%s] Cache Fail. Err: %s", middlewareName, url, err)
		return
	}
	routerMap := make(map[string]*platDb.Router, 0)
	err = json.Unmarshal([]byte(cache), &routerMap)
	if err != nil {
		log.WarnTF(traceID, "%s Json Unmarshal Router[%s] Cache Fail. Err: %s", middlewareName, url, err)
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
func readReq(c *gin.Context, router *platDb.Router, middlewareName, traceID string) error {
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
			log.ErrorTF(traceID, "%s ReqBody Read Fail: %s", middlewareName, err)
			c.Set(constant.RespBody, resp.SysErr)
			return err
		}
		// 写回body
		bodyGo := ioutil.NopCloser(bytes.NewBuffer(bodyBts))
		c.Request.Body = bodyGo
		printBts = bodyBts
	}

	log.InfoTF(traceID, "%s ReqBody: %s", middlewareName, printBts)
	return nil
}

// 处理业务响应
func returnJSON(c *gin.Context, router *platDb.Router, middlewareName, traceID string) {
	// 响应读取 读取失败响应系统异常
	respBody, ok := c.Get(constant.RespBody)
	if !ok {
		log.ErrorTF(traceID, "%s ResBody Get Fail", middlewareName)
		respBody = resp.SysErr
	}
	// 处理响应翻译
	returnMsgTrans(respBody, c, router, middlewareName, traceID)
}

// 执行响应码 => 响应文言 翻译
func returnMsgTrans(respBody any, c *gin.Context, router *platDb.Router, middlewareName, traceID string) {
	// 类型回转
	res, ok := respBody.(resp.ResBody)
	if !ok {
		log.ErrorTF(traceID, "%s ResBody Type UnKnow", middlewareName)
		res = resp.SysErr
	} else {
		// 非400错误执行翻译
		if res.Code != constant.ValidateFail {
			// 执行翻译
			res.Msg = TableMsgTrans(res, c.GetString(constant.HeaderLang), c.GetString(constant.TraceID))
		}
	}
	resBts, _ := json.Marshal(res)
	printBts := []byte("{}")
	if router.PrintRes == constant.RouterLogPrintNot {
		printBts = []byte("{ Res Set Not Print}")
	} else {
		// JSON序列化
		printBts = resBts
	}
	log.InfoTF(traceID, "%s RespBody: %s", middlewareName, printBts)
	// 响应结果
	c.Writer.Write(resBts)
}
