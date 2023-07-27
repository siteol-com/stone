package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
	"time"
)

// LoginRecord 用户表
type LoginRecord struct {
	ID        uint64     // 默认数据ID
	AccountId uint64     // 账号ID
	TenantId  uint64     // 租户ID
	LoginType string     // 登陆类型 1 账号登录
	LoginTime *time.Time // 登陆时间
	Token     string     // 登陆Token
	Common               // 登陆状态 1正常 2主动登出 3超限下线

}

// LoginRecordTable 登陆记录泛型构造器
var LoginRecordTable actuator.Table[LoginRecord]

// DataBase 实现指定数据库
func (t LoginRecord) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t LoginRecord) TableName() string {
	return "login_record"
}
