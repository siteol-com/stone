package platService

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/redis"
	"siteOl.com/stone/server/src/utils/log"
)

// InitResponseCodeCache 初始化响应码缓存
func InitResponseCodeCache() (err error) {
	allResCodes, err := (&platDb.ResponseCode{}).FindAll()
	if err != nil || len(allResCodes) == 0 {
		// 错误或数据库无配置直接返回
		return
	}
	// 组装缓存对象
	resCodeCacheMap := make(map[string]map[string]string, len(allResCodes))
	for _, res := range allResCodes {
		// 遍历支持的语言并写入Map
		langMap := make(map[string]string, len(constant.TransLangSupport))
		for _, lang := range constant.TransLangSupport {
			switch lang {
			case "en-US":
				langMap[lang] = res.EnUs
			case "zh-CN":
				langMap[lang] = res.ZhCn
			}
		}
		resCodeCacheMap[res.Code] = langMap
	}
	// 写入缓存 无超期
	err = redis.Set(constant.TransLangCacheKey, resCodeCacheMap, 0)
	if err == nil {
		log.InfoF("InitResponseCodeCache Success .")
	}
	return
}
