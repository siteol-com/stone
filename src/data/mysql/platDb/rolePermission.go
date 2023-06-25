package platDb

// RolePermission 角色权限表
type RolePermission struct {
	ID           uint64 // 默认数据ID
	RoleId       uint64 // 角色ID
	PermissionId uint64 // 权限ID
}

// TableName 实现自定义表名
func (t *RolePermission) TableName() string {
	return "role_permission"
}

// FindRolePermissionByIds 根据角色清单查权限（去重）
func (t *RolePermission) FindRolePermissionByIds(roleIds []uint64) (res []uint64, err error) {
	r := platDb.Distinct("permission_id").Where("role_id IN ?", roleIds).Find(&res)
	err = r.Error
	return
}
