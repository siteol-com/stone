package platModel

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"time"
)

// PermissionBashReq 权限基础查询对象（过滤前提）
type PermissionBashReq struct {
	IsSupper bool            // 超级租户查询（根据登陆账号判定）
	TenantId uint64          // 租户ID（非超级租户需要检索租户权限表）
	IDMap    map[uint64]bool // 租户可使用的权限集合
}

// PermissionBaseReq 权限处理基础对象
type PermissionBaseReq struct {
	Name      string   `json:"name" binding:"required,max=32" example:"账号管理"`                      // 权限名称，界面展示，建议与界面导航一致
	Alias     string   `json:"alias" binding:"required,max=32,letterUnder" example:"plat_account"` // 权限别名，英文+下划线，规范如下： plat  plat_account plat_account_add
	RouterIds []uint64 `json:"routerIds" binding:"unique" example:"1,2,3"`                         // 路由集，提交路由ID数组
}

// PermissionAddReq 权限创建对象
type PermissionAddReq struct {
	PermissionBaseReq
	Level  string `json:"level" binding:"required,oneof='1' '2' '3'" example:"1"` // 权限等级 1分组（一级导航）2模块（页面）3功能（按钮）
	Pid    uint64 `json:"pid" binding:"required,numeric" example:"1"`             // 父级ID，默认为1
	Static string `json:"static" binding:"required,oneof='1' '2'" example:"1"`    // 默认启用权限，1 不启 2启用  启用后，该权限默认被分配，不可去勾
}

// PermissionReqToDbReq 转换请求到数据库对象
func PermissionReqToDbReq(addRed *PermissionAddReq) *platDb.Permission {
	if addRed != nil {
		dbReq := &platDb.Permission{
			Name:      addRed.Name,
			Alias:     addRed.Alias,
			Level:     addRed.Level,
			Pid:       addRed.Pid,
			Static:    addRed.Static,
			RouterIds: addRed.RouterIds,
		}
		now := time.Now()
		dbReq.CreateAt = &now
		dbReq.Status = constant.StatusOpen
		return dbReq
	}
	return nil
}

// PermissionEditReq 权限修改对象
type PermissionEditReq struct {
	ID uint64 `json:"id" binding:"required,numeric" example:"1"` // 默认数据ID
	PermissionBaseReq
}

// PermissionBroRes 兄弟权限获取对象
type PermissionBroRes struct {
	ID   uint64 `json:"id" example:"1"`      // 默认数据ID
	Name string `json:"name" example:"账号管理"` // 权限名称，界面展示，建议与界面导航一致
	Sort uint16 `json:"sort" example:"1"`    // 权限排序
}
