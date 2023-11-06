package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/model/platModel"
	"siteOl.com/stone/server/src/sevices"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

// PageRole godoc
// @id			 PageRole角色分页
// @Summary      角色分页
// @Description  查询角色分页数据
// @Router       /plat/role/page [post]
// @Tags         角色配置
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.RolePageReq true "请求"
// @Success      200 {object} resp.ResBody{data=model.PageRes{list=[]platDb.Role}} "响应成功"
func PageRole(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.RolePageReq{})
	if err == nil {
		// TODO 添加租户ID
		sevices.JsonRes(c, platService.PageRole(traceID, req.(*platModel.RolePageReq)))
	}
}

// AddRole godoc
// @id			 AddRole角色创建
// @Summary      角色创建
// @Description  创建一个全新的角色
// @Router       /plat/role/add [post]
// @Tags         角色配置
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.RoleAddReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func AddRole(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.RoleAddReq{})
	if err == nil {
		// TODO 添加租户ID
		sevices.JsonRes(c, platService.AddRole(traceID, req.(*platModel.RoleAddReq)))
	}
}

// GetRole godoc
// @id			 GetRole角色查询
// @Summary      角色查询
// @Description  根据ID查询角色
// @Router       /plat/role/get [post]
// @Tags         角色配置
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.IdReq true "请求"
// @Success      200 {object} resp.ResBody{data=platDb.Role} "响应成功"
func GetRole(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &model.IdReq{})
	if err == nil {
		// TODO 添加租户ID
		req := req.(*model.IdReq)
		tenReq := &model.IdAnTenantReq{
			ID:       req.ID,
			TenantId: 0,
		}
		sevices.JsonRes(c, platService.GetRole(traceID, tenReq))
	}
}

// EditRole godoc
// @id			 EditRole角色编辑
// @Summary      角色编辑
// @Description  根据ID编辑角色，只有部分字段可以修改
// @Router       /plat/role/edit [post]
// @Tags         角色配置
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.RoleEditReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func EditRole(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.RoleEditReq{})
	if err == nil {
		// TODO 添加租户ID
		req := req.(*platModel.RoleEditReq)
		req.TenantId = 0
		sevices.JsonRes(c, platService.EditRole(traceID, req))
	}
}

// DelRole godoc
// @id			 DelRole角色删除
// @Summary      角色删除
// @Description  根据ID删除角色，本数据为软删除
// @Router       /plat/role/del [post]
// @Tags         角色配置
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.IdReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func DelRole(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &model.IdReq{})
	if err == nil {
		// TODO 添加租户ID
		req := req.(*model.IdReq)
		tenReq := &model.IdAnTenantReq{
			ID:       req.ID,
			TenantId: 0,
		}
		sevices.JsonRes(c, platService.DelRole(traceID, tenReq))
	}
}
