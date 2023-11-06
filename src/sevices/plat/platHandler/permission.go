package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/model/platModel"
	"siteOl.com/stone/server/src/sevices"
	"siteOl.com/stone/server/src/sevices/plat/platService"
)

// TreePermission godoc
// @id			 TreePermission权限树
// @Summary      权限树
// @Description  获取权限树，不同级别用户根据所属租户权限集看到权限树
// @Router       /plat/permission/tree [post]
// @Tags         访问权限
// @Accept       json
// @Produce      json
// @Security	 Token
// @Success      200 {object} resp.ResBody{data=[]model.Tree} "响应成功"
// TreePermission
func TreePermission(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.TraceID)
	// TODO 根据用户生成权限树查询对象
	req := &platModel.PermissionBashReq{
		IsSupper: true,
		TenantId: 1,
	}
	sevices.JsonRes(c, platService.TreePermission(traceID, req))
}

// AddPermission godoc
// @id			 AddPermission权限创建
// @Summary      权限创建
// @Description  创建一个全新的权限，支持勾选接口路由
// @Router       /plat/permission/add [post]
// @Tags         访问权限
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.PermissionAddReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func AddPermission(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.PermissionAddReq{})
	if err == nil {
		sevices.JsonRes(c, platService.AddPermission(traceID, req.(*platModel.PermissionAddReq)))
	}
}

// GetPermission godoc
// @id			 GetPermission权限查询
// @Summary      权限查询
// @Description  根据ID查询权限数据
// @Router       /plat/permission/get [post]
// @Tags         访问权限
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.IdReq true "请求"
// @Success      200 {object} resp.ResBody{data=platDb.Permission} "响应成功"
func GetPermission(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &model.IdReq{})
	if err != nil {
		sevices.JsonRes(c, platService.GetPermission(traceID, req.(*model.IdReq)))
	}
}

// EditPermission godoc
// @id			 EditPermission权限编辑
// @Summary      权限编辑
// @Description  根据ID对权限数据编辑，仅支持部分字段更新
// @Router       /plat/permission/edit [post]
// @Tags         访问权限
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body platModel.PermissionEditReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func EditPermission(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &platModel.PermissionEditReq{})
	if err != nil {
		sevices.JsonRes(c, platService.EditPermission(traceID, req.(*platModel.PermissionEditReq)))
	}
}

// DelPermission godoc
// @id			 DelPermission权限删除
// @Summary      权限删除
// @Description  根据ID对权限数据删除，存在子集时无法删除
// @Router       /plat/permission/del [post]
// @Tags         访问权限
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.IdReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func DelPermission(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &model.IdReq{})
	if err != nil {
		sevices.JsonRes(c, platService.DelPermission(traceID, req.(*model.IdReq)))
	}
}

// BroPermission godoc
// @id			 BroPermission获取兄弟权限
// @Summary      获取兄弟权限
// @Description  根据ID获取当前权限以及兄弟权限列表
// @Router       /plat/permission/bro [post]
// @Tags         访问权限
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.IdReq true "请求"
// @Success      200 {object} resp.ResBody{data=[]platModel.PermissionBroRes} "响应成功"
func BroPermission(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &model.IdReq{})
	if err != nil {
		sevices.JsonRes(c, platService.BroPermission(traceID, req.(*model.IdReq)))
	}
}

// SortPermission godoc
// @id			 SortPermission权限排序
// @Summary      权限排序
// @Description  同级权限排序功能
// @Router       /plat/permission/sort [post]
// @Tags         访问权限
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body []model.SortReq true "请求"
// @Success      200 {object} resp.ResBody{data=bool} "响应成功"
func SortPermission(c *gin.Context) {
	traceID, req, err := sevices.ValidateReqObj(c, &[]*model.SortReq{})
	if err != nil {
		sevices.JsonRes(c, platService.SortPermission(traceID, req.(*[]*model.SortReq)))
	}
}
