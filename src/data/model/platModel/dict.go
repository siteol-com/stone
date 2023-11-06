package platModel

import "siteOl.com/stone/server/src/data/mysql/platDb"

// DictListReq 字典下拉列表
type DictListReq struct {
	GroupKeys []string `json:"groupKeys" example:"'serviceCode','responseType'"` // 需要查询的字典分组
	Local     string   `json:"-"`                                                // 字典语言
}

// DictListRes 字典下拉响应
type DictListRes struct {
	List map[string][]*DictRes        `json:"list"` // 字典下拉列表 {'serviceCode':"[{'label':'基础','value':'1'}]"}
	Map  map[string]map[string]string `json:"map"`  // 字典翻译Map {'serviceCode':{'1':'基础'}}
}

// DictRes 字典下拉响应
type DictRes struct {
	Label string `json:"label" example:"基础"` // 字典名
	Value string `json:"value" example:"1"`  // 字典值
}

// DictToListMapRes 内部对象到下拉对象数组
func DictToListMapRes(dictList []*platDb.Dict, local string) ([]*DictRes, map[string]string) {
	labelList := make([]*DictRes, len(dictList))
	valueMap := make(map[string]string, len(dictList))
	for i, dict := range dictList {
		labelList[i] = &DictRes{
			Label: dict.Label,
			Value: dict.Val,
		}
		valueMap[dict.Val] = dict.Label
		// 文言翻译
		switch local {
		case "en":
			labelList[i].Label = dict.LabelEn
			valueMap[dict.Val] = dict.LabelEn
		}
	}
	return labelList, valueMap
}
