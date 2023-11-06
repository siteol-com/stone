package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/model/platModel"
	"siteOl.com/stone/server/src/sevices"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

// PageRouter godoc
// @id			 PageRouter路由分页
// @Summary      路由分页
// @Description  查询路由分页数据
// @Router       /plat/router/page [post]
// @Tags         路由接口
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.RouterPageReq true "请求"
// @Success      200 {object} resp.ResBody{data=model.PageRes{list=[]platDb.Router}} "响应成功"
func PageRouter(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.RouterPageReq{})
	if err == nil {
		sevices.JsonRes(c, platService.PageRouter(traceID, req.(*platModel.RouterPageReq)))
	}
}

// AddRouter godoc
// @id			 AddRouter路由创建
// @Summary      路由创建
// @Description  创建一个全新的路由
// @Router       /plat/router/add [post]
// @Tags         路由接口
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.RouterAddReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func AddRouter(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.RouterAddReq{})
	if err == nil {
		sevices.JsonRes(c, platService.AddRouter(traceID, req.(*platModel.RouterAddReq)))
	}
}

// GetRouter godoc
// @id			 GetRouter路由查询
// @Summary      路由查询
// @Description  根据ID查询路由
// @Router       /plat/router/get [post]
// @Tags         路由接口
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.IdReq true "请求"
// @Success      200 {object} resp.ResBody{data=platDb.Router} "响应成功"
func GetRouter(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &model.IdReq{})
	if err == nil {
		sevices.JsonRes(c, platService.GetRouter(traceID, req.(*model.IdReq)))
	}
}

// EditRouter godoc
// @id			 EditRouter路由编辑
// @Summary      路由编辑
// @Description  根据ID编辑路由，只有部分字段可以修改
// @Router       /plat/router/edit [post]
// @Tags         路由接口
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.RouterEditReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func EditRouter(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.RouterEditReq{})
	if err == nil {
		sevices.JsonRes(c, platService.EditRouter(traceID, req.(*platModel.RouterEditReq)))
	}
}

// DelRouter godoc
// @id			 DelRouter路由删除
// @Summary      路由删除
// @Description  根据ID删除路由，本数据为软删除
// @Router       /plat/router/del [post]
// @Tags         路由接口
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.IdReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"

func DelRouter(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &model.IdReq{})
	if err == nil {
		sevices.JsonRes(c, platService.DelRouter(traceID, req.(*model.IdReq)))
	}
}
