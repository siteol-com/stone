package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/sevices/plat"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/sevices/plat/platService"
	"strings"
)

// ListDict /plat/dict/list 获取字典下拉列表
func ListDict(c *gin.Context) {
	traceID, reqObj, err := plat.ValidateReqObj(c, &platModel.DictListReq{})
	if err != nil {
		return
	}
	req := reqObj.(*platModel.DictListReq)
	// 语言转换
	req.Local = c.GetString(constant.HeaderLang)
	req.Local = req.Local[:strings.Index(req.Local, "-")]
	// 执行查询
	res := platService.ListDict(traceID, req)
	c.Set(constant.RespBody, res)
}
