package platDb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/config"
	"siteOl.com/stone/server/src/utils/log"
	"time"
)

var platDb *gorm.DB

// InitPlatFromDb 初始化平台数据库
func InitPlatFromDb() {
	// 采用默认配置打开数据可
	db, err := gorm.Open(mysql.Open(config.JsonConfig.MySQL.Plat), &gorm.Config{})
	if err != nil {
		log.FatalF("Open PlatDb Fail . Err Is : %s", err)
		return
	}
	platDb = db
	log.InfoF("Init PlatDb Success . ")
}

// Common 平台通用信息体
type Common struct {
	Status   string     `json:"status"`   // 状态 0正常 1锁定 2封存
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}
