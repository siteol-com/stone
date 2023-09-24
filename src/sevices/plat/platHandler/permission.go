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

// TreePermission /plat/permission/tree 权限树
func TreePermission(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.TraceID)
	// TODO 根据用户生成权限树查询对象
	req := &platModel.PermissionBashReq{
		IsSupper: true,
		TenantId: 1,
	}
	res := platService.TreePermission(traceID, req)
	c.Set(constant.RespBody, res)
}

// AddPermission /plat/permission/add 权限创建
func AddPermission(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platDb.Permission{})
	if err != nil {
		return
	}
	res := platService.AddPermission(traceID, req.(*platDb.Permission))
	c.Set(constant.RespBody, res)
}

// GetPermission /plat/permission/get 权限查询
func GetPermission(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &model.IdReq{})
	if err != nil {
		return
	}
	res := platService.GetPermission(traceID, req.(*model.IdReq))
	c.Set(constant.RespBody, res)
}

// EditPermission /plat/permission/edit 权限编辑
func EditPermission(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &platDb.Permission{})
	if err != nil {
		return
	}
	res := platService.EditPermission(traceID, req.(*platDb.Permission))
	c.Set(constant.RespBody, res)
}

// DelPermission /plat/permission/del 权限删除
func DelPermission(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &model.IdReq{})
	if err != nil {
		return
	}
	res := platService.DelPermission(traceID, req.(*model.IdReq))
	c.Set(constant.RespBody, res)
}

// BroPermission /plat/permission/bro 获取兄弟权限
func BroPermission(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &model.IdReq{})
	if err != nil {
		return
	}
	res := platService.BroPermission(traceID, req.(*model.IdReq))
	c.Set(constant.RespBody, res)
}

// SortPermission /plat/permission/sort 权限排序
func SortPermission(c *gin.Context) {
	traceID, req, err := plat.ValidateReqObj(c, &[]*model.SortReq{})
	if err != nil {
		return
	}
	res := platService.SortPermission(traceID, req.(*[]*model.SortReq))
	c.Set(constant.RespBody, res)
}
