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

// PageRole /plat/role/page 角色分页
func PageRole(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platModel.RolePageReq{})
	if err != nil {
		return
	}
	// TODO 添加租户ID
	res := platService.PageRole(traceID, req.(*platModel.RolePageReq))
	c.Set(constant.RespBody, res)
}

// AddRole /plat/role/add 角色创建
func AddRole(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platDb.Role{})
	if err != nil {
		return
	}
	// TODO 添加租户ID
	res := platService.AddRole(traceID, req.(*platDb.Role))
	c.Set(constant.RespBody, res)
}

// GetRole /plat/role/get 角色查询
func GetRole(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &model.IdAnTenantReq{})
	if err != nil {
		return
	}
	// TODO 添加用户租户ID
	res := platService.GetRole(traceID, req.(*model.IdAnTenantReq))
	c.Set(constant.RespBody, res)
}

// EditRole /plat/role/edit 角色编辑
func EditRole(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platDb.Role{})
	if err != nil {
		return
	}
	// TODO 添加用户租户ID
	res := platService.EditRole(traceID, req.(*platDb.Role))
	c.Set(constant.RespBody, res)
}

// DelRole /plat/role/del 角色删除
func DelRole(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &model.IdAnTenantReq{})
	if err != nil {
		return
	}
	// TODO 添加用户租户ID
	res := platService.DelRole(traceID, req.(*model.IdAnTenantReq))
	c.Set(constant.RespBody, res)
}
