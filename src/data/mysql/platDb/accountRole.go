package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// AccountRole 账号角色表
type AccountRole struct {
	ID        uint64 // 默认数据ID
	AccountId uint64 // 账号ID
	RoleId    uint64 // 角色ID

}

// AccountRoleTable 账号角色泛型构造器
var AccountRoleTable actuator.Table[AccountRole]

// DataBase 实现指定数据库
func (t AccountRole) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t AccountRole) TableName() string {
	return "account_role"
}

// FindAccountRoleIds 读取账号角色
func (t AccountRole) FindAccountRoleIds() (res []uint64, err error) {
	r := platDb.Select("role_id").Where(t).Find(&res)
	err = r.Error
	return
}
