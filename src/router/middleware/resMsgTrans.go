package middleware

import (
	"encoding/json"
	"fmt"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/redis"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/utils/log"
	"strings"
)

// TableMsgTrans 执行Msg翻译
func TableMsgTrans(res resp.ResBody, lang, traceID string) string {
	// 获取翻译缓存
	tranStr, err := redis.Get(constant.CacheKeyTransLang)
	if err != nil {
		log.ErrorTF(traceID, "GetTransLangCacheMap Fail . Err Is : %v", err)
		// 出错不翻译
		return res.Msg

	}
	transMap := make(map[string]map[string]string)
	err = json.Unmarshal([]byte(tranStr), &transMap)
	if err != nil {
		log.ErrorTF(traceID, "JsonUnmarshal TransMap Fail . Err Is : %v", err)
		// 出错不翻译
		return res.Msg
	}
	// 读取配置
	codeMap, ok := transMap[res.Code]
	if ok {
		langTemple, lok := codeMap[lang]
		if lok {
			// 检查是否有变量
			if strings.Index(langTemple, "}}") > -1 {
				return TableValReplace(langTemple, res.Data)
			} else {
				return langTemple
			}
		}
	}
	// 无相关翻译
	return res.Msg
}

// 执行变量替换
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
