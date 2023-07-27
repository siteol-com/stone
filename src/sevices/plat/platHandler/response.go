package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/sevices/plat"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

// PageResponse /plat/response/page 响应码分页
func PageResponse(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platModel.ResponsePageReq{})
	if err != nil {
		return
	}
	res := platService.PageResponse(traceID, req.(*platModel.ResponsePageReq))
	c.Set(constant.RespBody, res)
}

// AddResponse /plat/response/add 响应码创建
func AddResponse(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platDb.Response{})
	if err != nil {
		return
	}
	res := platService.AddResponse(traceID, req.(*platDb.Response))
	c.Set(constant.RespBody, res)
}

// GetResponse /plat/response/get 响应码查询
func GetResponse(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &model.IdReq{})
	if err != nil {
		return
	}
	res := platService.GetResponse(traceID, req.(*model.IdReq))
	c.Set(constant.RespBody, res)
}

// EditResponse /plat/response/edit 响应码编辑
func EditResponse(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platDb.Response{})
	if err != nil {
		return
	}
	res := platService.EditResponse(traceID, req.(*platDb.Response))
	c.Set(constant.RespBody, res)
}

// DelResponse /plat/response/del 响应码删除
func DelResponse(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &model.IdReq{})
	if err != nil {
		return
	}
	res := platService.DelResponse(traceID, req.(*model.IdReq))
	c.Set(constant.RespBody, res)
}
