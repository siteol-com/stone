package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// Dict 字典表
type Dict struct {
	ID       uint64 // 数据ID
	GroupKey string // 字典分组Key
	Label    string // 字典名称
	LabelEn  string // 英文字典名称
	Val      string // 字典值
	Pid      uint64 // 父级字典ID 默认 0（根数据）
	Sort     uint16 // 字典排序
	Remark   string // 字典描述
	Common
}

// DictGroup 字典分组
type DictGroup struct {
	ID     uint64 // 数据ID
	Key    string // 字典分组Key
	Name   string // 字典名称
	NameEn string // 字典名称
}

// DictTable 字典泛型构造器
var DictTable actuator.Table[Dict]

// DictGroupTable 字典分组泛型构造器
var DictGroupTable actuator.Table[DictGroup]

// DataBase 实现指定数据库
func (t Dict) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Dict) TableName() string {
	return "dict"
}

// DataBase 实现指定数据库
func (t DictGroup) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t DictGroup) TableName() string {
	return "dict_group"
}

// FindSelectList 查询下拉选择列表
func (t Dict) FindSelectList() (res []*Dict, err error) {
	t.Status = constant.StatusOpen
	r := platDb.Where(t).Order("sort").Find(&res)
	err = r.Error
	return
}
