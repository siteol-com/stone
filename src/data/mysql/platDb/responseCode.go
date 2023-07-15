package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// INSERT INTO `response_code`(`code`, `service_code`, `response_type`, `zh_cn`, `en_us`, `remark`, `mark`, `status`, `create_at`, `update_at`) VALUES ('2001000', 1, 2, '租户信息获取成功', 'En', '租户相关成功信息', 0, 1, '2023-03-03 03:33:33', '2023-03-03 03:33:33');

// ResponseCode 响应码
type ResponseCode struct {
	ID           uint64 // 数据ID
	Code         string // 响应码 2/5+XXX+XXX
	ServiceCode  string // 业务ID，来源于字典，指定响应码归属业务
	ResponseType string // 响应类型，该字段用于筛选，可配置2和5
	ZhCn         string // 中文响应文言
	EnUs         string // 英文响应文言
	remark       string // 其他备注信息
	Mark         string // 变更标识 0可变更 1禁止变更
	Common
}

// ResponseCodeTable 响应码泛型构造器
var ResponseCodeTable actuator.Table[ResponseCode]

// DataBase 实现指定数据库
func (t ResponseCode) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t ResponseCode) TableName() string {
	return "response_code"
}

// FindAll 基于对象实施查询
func (t ResponseCode) FindAll() (res []ResponseCode, err error) {
	r := platDb.Find(&res)
	if r.Error != nil {
		err = r.Error
	}
	return
}
