package resp

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/redis"
	"siteOl.com/stone/server/src/utils/log"
	"strings"
)

// ReturnMsgTrans 执行响应码 => 响应文言 翻译
func ReturnMsgTrans(res *ResBody, c *gin.Context, router *platDb.Router, traceID string) {
	// 语言类型
	lang := c.GetString(constant.HeaderLang)
	// 非400校验错误执行翻译
	if res.HttpCode != http.StatusBadRequest {
		// 执行翻译
		tableMsgTrans(res, lang, traceID)
	}
	// 响应日志
	printBts := []byte("{}")
	if router.PrintRes == constant.RouterLogPrintNot {
		printBts = []byte("{ Res Set Not Print}")
	} else {
		// 如需打印日志
		resBts, _ := json.Marshal(res)
		// JSON序列化
		printBts = resBts
	}
	log.InfoTF(traceID, "RespBody: %s", printBts)
}

// tableMsgTrans 执行Msg翻译
func tableMsgTrans(res *ResBody, lang, traceID string) {
	// 获取翻译缓存
	tranStr, err := redis.Get(constant.CacheKeyTransLang)
	if err != nil {
		log.ErrorTF(traceID, "GetTransLangCacheMap Fail . Err Is : %v", err)
		// 出错不翻译
		return

	}
	transMap := make(map[string]map[string]string)
	err = json.Unmarshal([]byte(tranStr), &transMap)
	if err != nil {
		log.ErrorTF(traceID, "JsonUnmarshal TransMap Fail . Err Is : %v", err)
		// 出错不翻译
		return
	}
	// 读取配置
	codeMap, ok := transMap[res.Code]
	if ok {
		langTemple, lok := codeMap[lang]
		if lok {
			// 检查是否有变量
			// 有进行变量替换
			if strings.Index(langTemple, "}}") > -1 {
				res.Msg = TableValReplace(langTemple, res.Data)
			} else {
				res.Msg = langTemple
			}
		}
	}
	// 无相关翻译
	return
}

// TableValReplace 执行变量替换
// 实际生效，{{name}}修改成功 => 米虫修改成功
func TableValReplace(temple string, data any) string {
	if data == nil {
		return temple
	}
	dataStr, err := json.Marshal(data)
	if err != nil {
		return temple
	}
	dataMap := make(map[string]any)
	err = json.Unmarshal([]byte(dataStr), &dataMap)
	if err != nil {
		return temple
	}
	// 提取模板中的变量数据
	valArray := getTempleVal(temple)
	// 提取参数并替换
	for _, val := range valArray {
		valObj, ok := dataMap[val]
		if ok && valObj != nil {
			// 变量存在值则替换
			temple = strings.ReplaceAll(temple, "{{"+val+"}}", fmt.Sprintf("%v", valObj))
		}
	}
	return temple
}

// 提取模板中的变量数据
func getTempleVal(temple string) []string {
	valArray := make([]string, 0)
	strS := strings.Split(temple, "{{")
	for _, i := range strS {
		if strings.Index(i, "}}") > -1 {
			valStr := i[:strings.Index(i, "}}")]
			valArray = append(valArray, valStr)
		}
	}
	return valArray
}
