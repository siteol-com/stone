package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model/platModel"
	"siteOl.com/stone/server/src/sevices"
	"siteOl.com/stone/server/src/sevices/plat/platService"
	"strings"
)

// ListDict godoc
// @id			 ListDict字典下拉列表
// @Summary      字典下拉列表
// @Description  获取字典下拉列表，用于选择框
// @Router       /plat/dict/list [post]
// @Tags         数据字典
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        Lang header string false "语言，不传默认为zh-CN"
// @Param        req body platModel.DictListReq true "请求"
// @Success      200 {object} resp.ResBody{data=platModel.DictListRes} "响应成功"
func ListDict(c *gin.Context) {
	traceID, reqObj, err := sevices.ValidateReqObj(c, &platModel.DictListReq{})
	if err == nil {
		req := reqObj.(*platModel.DictListReq)
		// 语言转换
		req.Local = c.GetString(constant.HeaderLang)
		req.Local = req.Local[:strings.Index(req.Local, "-")]
		// 执行查询
		sevices.JsonRes(c, platService.ListDict(traceID, req))
	}
}
