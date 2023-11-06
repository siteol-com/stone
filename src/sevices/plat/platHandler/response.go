package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/model/platModel"
	"siteOl.com/stone/server/src/sevices"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

// PageResponse godoc
// @id			 PageResponse响应码分页
// @Summary      响应码分页
// @Description  查询响应码分页数据
// @Router       /plat/response/page [post]
// @Tags         响应文言
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.ResponsePageReq true "请求"
// @Success      200 {object} resp.ResBody{data=model.PageRes{list=[]platDb.Response}} "响应成功"
func PageResponse(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.ResponsePageReq{})
	if err == nil {
		sevices.JsonRes(c, platService.PageResponse(traceID, req.(*platModel.ResponsePageReq)))
	}
}

// AddResponse godoc
// @id			 AddResponse响应码创建
// @Summary      响应码创建
// @Description  创建一个全新的响应码
// @Router       /plat/response/add [post]
// @Tags         响应文言
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.ResponseAddReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func AddResponse(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.ResponseAddReq{})
	if err == nil {
		sevices.JsonRes(c, platService.AddResponse(traceID, req.(*platModel.ResponseAddReq)))
	}
}

// GetResponse godoc
// @id			 GetResponse响应码查询
// @Summary      响应码查询
// @Description  根据ID查询响应码
// @Router       /plat/response/get [post]
// @Tags         响应文言
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.IdReq true "请求"
// @Success      200 {object} resp.ResBody{data=platDb.Response} "响应成功"
func GetResponse(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &model.IdReq{})
	if err == nil {
		sevices.JsonRes(c, platService.GetResponse(traceID, req.(*model.IdReq)))
	}
}

// EditResponse godoc
// @id			 EditResponse响应码编辑
// @Summary      响应码编辑
// @Description  根据ID编辑响应码，只有部分字段可以修改
// @Router       /plat/response/edit [post]
// @Tags         响应文言
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.ResponseEditReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func EditResponse(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.ResponseEditReq{})
	if err == nil {
		sevices.JsonRes(c, platService.EditResponse(traceID, req.(*platModel.ResponseEditReq)))
	}
}

// DelResponse godoc
// @id			 DelResponse响应码删除
// @Summary      响应码删除
// @Description  根据ID删除响应码，本数据为软删除
// @Router       /plat/response/del [post]
// @Tags         响应文言
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.IdReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func DelResponse(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &model.IdReq{})
	if err == nil {
		sevices.JsonRes(c, platService.DelResponse(traceID, req.(*model.IdReq)))
	}
}
