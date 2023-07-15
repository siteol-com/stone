package platModel

import (
	"siteOl.com/stone/server/src/data/mysql/platDb"
)

// DictListReq 字典下拉列表
type DictListReq struct {
	GroupKeys []string `json:"groupKeys"` // 需要查询的字典分组
	Local     string   `json:"-"`         // 字典语言
}

// DictListRes 字典下拉响应
type DictListRes struct {
	Label string `json:"label"` // 字典名
	Value string `json:"value"` // 字典值
}

// DictToListMapRes 内部对象到下拉对象数组
func DictToListMapRes(dictList []*platDb.Dict, local string) ([]*DictListRes, map[string]string) {
	labelList := make([]*DictListRes, len(dictList))
	valueMap := make(map[string]string, len(dictList))
	for i, dict := range dictList {
		labelList[i] = &DictListRes{
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
