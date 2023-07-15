package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
	"time"
)

// Account 用户表
type Account struct {
	ID             uint64     // 账号ID
	Account        string     // 账号
	Encryption     string     // 密文密码
	SaltKey        string     // 盐值秘钥
	PwdExpTime     *time.Time // 密码超期时间（修改后的90天）
	LastLoginTime  *time.Time //最后登陆时间
	TenantId       uint64     // 租户ID
	DeptId         uint64     // 部门ID
	PermissionType string     // 权限类型 0全局数据 1跟随部门 2仅本部门 3本部门及子部门
	Name           string     // 姓名（拓展信息）
	Email          string     // 邮箱（拓展信息）
	Phone          string     // 手机号（拓展信息）
	Mark           string     // 变更标识 0可变更 1禁止变更（除密码）
	Common
}

// AccountTable 账号泛型构造器
var AccountTable actuator.Table[Account]

// DataBase 实现指定数据库
func (t Account) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Account) TableName() string {
	return "account"
}
