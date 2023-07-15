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

// PageRouter /plat/router/page 路由分页
func PageRouter(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platModel.RouterPageReq{})
	if err != nil {
		return
	}
	res := platService.PageRouter(traceID, req.(*platModel.RouterPageReq))
	c.Set(constant.RespBody, res)
}

// AddRouter /plat/router/add 路由创建
func AddRouter(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platDb.Router{})
	if err != nil {
		return
	}
	res := platService.AddRouter(traceID, req.(*platDb.Router))
	c.Set(constant.RespBody, res)
}

// GetRouter /plat/router/get 路由查询
func GetRouter(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &model.IdReq{})
	if err != nil {
		return
	}
	res := platService.GetRouter(traceID, req.(*model.IdReq))
	c.Set(constant.RespBody, res)
}

// EditRouter /plat/router/edit 路由编辑
func EditRouter(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platDb.Router{})
	if err != nil {
		return
	}
	res := platService.EditRouter(traceID, req.(*platDb.Router))
	c.Set(constant.RespBody, res)
}

// DelRouter /plat/router/del 路由删除
func DelRouter(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &model.IdReq{})
	if err != nil {
		return
	}
	res := platService.DelRouter(traceID, req.(*model.IdReq))
	c.Set(constant.RespBody, res)
}
