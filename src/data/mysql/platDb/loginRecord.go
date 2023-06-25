package platDb

import "time"

// LoginRecord 用户表
type LoginRecord struct {
	Id        uint64     // 默认数据ID
	AccountId uint64     // 账号ID
	TenantId  uint64     // 租户ID
	LoginType uint8      // 登陆类型 1 账号登录
	LoginTime *time.Time // 登陆时间
	Token     string     // 登陆Token
	Common               // 登陆状态 1正常 2主动登出 3超限下线

}

// TableName 实现自定义表名
func (t *LoginRecord) TableName() string {
	return "login_record"
}

// InsertByObj 实现自定义表名
func (t *LoginRecord) InsertByObj() (uint64, error) {
	r := platDb.Create(t)
	return t.Id, r.Error
}
