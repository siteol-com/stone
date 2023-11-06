package platModel

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"time"
)

// RolePageReq 角色分页查询
type RolePageReq struct {
	Name     string `json:"name" example:"admin"` // 角色名称
	TenantId uint64 `json:"-"`                    // 租户ID（根据登陆用户读取）
	model.PageReq
}

// SetTenantId RolePageReq 设置租户ID
func (t *RolePageReq) SetTenantId(id uint64) {
	t.TenantId = id
}

// RoleBaseReq 角色基础请求
type RoleBaseReq struct {
	Name          string   `json:"name" binding:"required,max=16" example:"admin"` // 角色名称
	Remark        string   `json:"remark" binding:"max=64" example:"管理员"`          // 角色备注
	PermissionIds []uint64 `json:"permissionIds" binding:"unique" example:"1,2,3"` // 权限集
	TenantId      uint64   `json:"-"`                                              // 租户ID，后台赋值
}

// SetTenantId RoleBaseReq 设置租户ID
func (t *RoleBaseReq) SetTenantId(id uint64) {
	t.TenantId = id
}

// RoleAddReq 角色添加请求
type RoleAddReq struct {
	RoleBaseReq
}

// RoleReqToDbReq 转换请求到数据库对象
func RoleReqToDbReq(addRed *RoleAddReq) *platDb.Role {
	if addRed != nil {
		dbReq := &platDb.Role{
			Name:          addRed.Name,
			Remark:        addRed.Remark,
			TenantId:      addRed.TenantId,     // 待添加
			Mark:          constant.StatusOpen, // 平台创建角色可变更
			PermissionIds: addRed.PermissionIds,
		}
		now := time.Now()
		dbReq.CreateAt = &now
		dbReq.Status = constant.StatusOpen
		return dbReq
	}
	return nil
}

// RoleEditReq 角色编辑请求
type RoleEditReq struct {
	ID uint64 `json:"id"  binding:"required,numeric" example:"1"` // 数据ID
	RoleBaseReq
}
