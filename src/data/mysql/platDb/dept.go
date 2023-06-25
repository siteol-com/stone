package platDb

import (
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// Dept 部门表
type Dept struct {
	Id             uint64 // 默认数据ID
	Name           string // 部门名称
	Pid            uint64 // 父级部门ID，租户创建时默认创建根部门，父级ID=0
	PermissionType uint8  // 权限类型 0全局数据 2仅本部门 3本部门及子部门
	TenantId       uint64 // 租户ID
	Common
}

// TableName 实现自定义表名
func (t *Dept) TableName() string {
	return "dept"
}

// FindOne 基于对象实施查询
func (t *Dept) FindOne() (res *Dept, err error) {
	res = new(Dept)
	err = actuator.FindByObject(platDb, t, res)
	return
}

// FindList 根据对象获取列表
func (t *Dept) FindList() (res []*Dept, err error) {
	err = actuator.FindByObject(platDb, t, res)
	return
}
