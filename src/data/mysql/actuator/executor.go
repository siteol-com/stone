package actuator

import "gorm.io/gorm"

// MYSQL 执行器
// 常见数据库执行API

// FindByObject 实体类作为查询条件
func FindByObject(db *gorm.DB, query interface{}, res interface{}) error {
	r := db.Where(query).Find(res)
	return r.Error
}
