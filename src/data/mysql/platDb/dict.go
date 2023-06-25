package platDb

// Dict 字典表
type Dict struct {
	ID       uint64 // 数据ID
	GroupKey string // 字典分组Key
	IntVal   uint8  // 字典值（数字型）
	StrVal   string // 字典值（字符型）
	Pid      uint64 // 父级字典ID 默认 1（根数据）
	Sort     uint16 // 字典排序
	Remark   string // 字典描述
	Common
}

// DictGroup 字典分组
type DictGroup struct {
	ID   uint64 // 数据ID
	Name string // 字典名称
	Key  string // 字典分组Key
}

// TableName 实现自定义表名
func (t *Dict) TableName() string {
	return "dict"
}

// TableName 实现自定义表名
func (t *DictGroup) TableName() string {
	return "dict_group"
}

// FindAll 基于对象实施查询
func (t *DictGroup) FindAll() (res []DictGroup, err error) {
	r := platDb.Find(&res)
	if r.Error != nil {
		err = r.Error
	}
	return
}
