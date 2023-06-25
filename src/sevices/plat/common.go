package plat

import (
	"github.com/gin-gonic/gin"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
)

/**
 *
 * 平台公共方法
 *
 *
 * @author 米虫@mebugs.com
 * @since 2023-03-08
 */

// GetLoginUser 从上下文获取登录用户授权机构体
func GetLoginUser(c *gin.Context) *platModel.AuthUser {
	obj, ok := c.Get(constant.AuthUser)
	if ok {
		authUser := &platModel.AuthUser{}
		authUser = obj.(*platModel.AuthUser)
		return authUser
	}
	return nil
}
