package platDb

// AccountRole 用户角色表
type AccountRole struct {
	ID        uint64 // 默认数据ID
	AccountId uint64 // 账号ID
	RoleId    uint64 // 角色ID

}

// TableName 实现自定义表名
func (t *AccountRole) TableName() string {
	return "account_role"
}

// FindAccountRoleIds 读取账号角色
func (t *AccountRole) FindAccountRoleIds() (res []uint64, err error) {
	r := platDb.Select("role_id").Where(t).Find(&res)
	err = r.Error
	return
}
