package platService

import (
	"siteOl.com/stone/server/src/data/model/platModel"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/utils/log"
)

// ListDict 读取字典下拉列表
func ListDict(traceID string, req *platModel.DictListReq) *resp.ResBody {
	// 如果查询key不为空
	if len(req.GroupKeys) > 0 {
		dictListMap := make(map[string][]*platModel.DictRes, len(req.GroupKeys))
		dictValueMap := make(map[string]map[string]string, len(req.GroupKeys))
		// 遍历查询
		for _, groupKey := range req.GroupKeys {
			dictList, err := (&platDb.Dict{GroupKey: groupKey}).FindSelectList()
			if err != nil {
				log.WarnTF(traceID, "ListDict Fail . GroupKey Query By : %s , Err is : %v", groupKey, err)
				dictListMap[groupKey] = make([]*platModel.DictRes, 0)
				continue
			}
			dictListMap[groupKey], dictValueMap[groupKey] = platModel.DictToListMapRes(dictList, req.Local)
		}
		return resp.SuccessUnPop(platModel.DictListRes{List: dictListMap, Map: dictValueMap})
	}
	return resp.SuccessUnPop(nil)
}
