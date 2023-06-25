package platDb

// Permission 权限表
type Permission struct {
	ID     uint64 // 默认数据ID
	Name   string // 权限名称，界面展示，建议与界面导航一致
	Alias  string // 权限别名，英文+下划线，规范如下： sys  sys_account sys_account_add
	Level  uint8  // 权限等级 1分组（一级导航）2模块（页面）3功能（按钮） 第四级路由不在本表中体现
	Pid    uint64 // 父级ID，默认为1
	Sort   uint16 // 字典排序
	Static uint8  // 默认启用权限，1 启用 2 不启 启用后，该权限默认被分配，不可去勾
	Common
}

// TableName 实现自定义表名
func (t *Permission) TableName() string {
	return "permission"
}

// FindByIds 根据IDS获取权限别名
func (t *Permission) FindByIds(ids []uint64) (res []*Permission, err error) {
	r := platDb.Where("id IN ?", ids).Find(&res)
	err = r.Error
	return
}
