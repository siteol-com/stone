package sevices

import (
	"errors"
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/data/validate"
)

/**
 *
 * 服务类公共方法
 *
 *
 * @author 米虫丨www.mebugs.com
 * @since 2023-03-08
 */

// GetLoginUser 从上下文获取登录用户授权机构体
func GetLoginUser(c *gin.Context) *model.AuthUser {
	obj, ok := c.Get(constant.AuthUser)
	if ok {
		authUser := &model.AuthUser{}
		authUser = obj.(*model.AuthUser)
		return authUser
	}
	return nil
}

// ValidateReqObj 读取并验证请求数据（并处理响应）
func ValidateReqObj(c *gin.Context, req any) (traceID string, reqObj any, err error) {
	// traceID 日志追踪
	traceID = c.GetString(constant.TraceID)
	// 校验并且 解析请求数据
	res, reqObj := validate.Readable(c, req)
	if res != nil {
		err = errors.New(res.Msg)
		// 处理响应
		JsonRes(c, res)
	}
	return
}

// GetRouterConf 从上下文获取登录用户授权机构体
func GetRouterConf(c *gin.Context) *platDb.Router {
	obj, ok := c.Get(constant.RouterConf)
	if ok {
		router := &platDb.Router{}
		router = obj.(*platDb.Router)
		return router
	}
	return nil
}

// JsonRes 执行Json响应
func JsonRes(c *gin.Context, res *resp.ResBody) {
	// traceID 日志追踪
	traceID := c.GetString(constant.TraceID)
	// 获取路由配置
	router := GetRouterConf(c)
	// 对Res进行翻译
	resp.ReturnMsgTrans(res, c, router, traceID)
	c.JSON(res.HttpCode, res)
}
