package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// INSERT INTO `response_code`(`code`, `service_code`, `response_type`, `zh_cn`, `en_us`, `remark`, `mark`, `status`, `create_at`, `update_at`) VALUES ('2001000', 1, 2, '租户信息获取成功', 'En', '租户相关成功信息', 0, 1, '2023-03-03 03:33:33', '2023-03-03 03:33:33');

// Response 响应码
type Response struct {
	ID          uint64 `json:"id"  binding:"numeric"`                  // 数据ID
	Code        string `json:"code"`                                   // 响应码 2/5+XXX+XXX
	ServiceCode string `json:"serviceCode" binding:"required,numeric"` // 业务ID，来源于字典，指定响应码归属业务
	Type        string `json:"type" binding:"required,oneof='2' '5'"`  // 响应类型，该字段用于筛选，可配置2和5
	ZhCn        string `json:"zhCn"`                                   // 中文响应文言
	EnUs        string `json:"enUs"`                                   // 英文响应文言
	Remark      string `json:"remark"`                                 // 其他备注信息
	Mark        string `json:"mark"`                                   // 变更标识 0可变更 1禁止变更
	Common
}

// ResponseTable 响应码泛型构造器
var ResponseTable actuator.Table[Response]

// DataBase 实现指定数据库
func (t Response) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Response) TableName() string {
	return "response"
}

// FindAll 基于对象实施查询
func (t Response) FindAll() (res []Response, err error) {
	r := platDb.Find(&res)
	if r.Error != nil {
		err = r.Error
	}
	return
}

// CountByGroup 运算分组下的响应码
func (t Response) CountByGroup() (res int64, err error) {
	res = int64(0)
	r := platDb.Raw("SELECT COUNT(id) FROM `response` WHERE `service_code` = ? AND `type` = ?", t.ServiceCode, t.Type).Find(&res)
	if r.Error != nil {
		err = r.Error
	}
	return
}
