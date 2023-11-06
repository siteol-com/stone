package platDb

import (
	"gorm.io/gorm"
	"siteOl.com/stone/server/src/data/mysql/actuator"
)

// Permission 权限表
type Permission struct {
	ID     uint64 `json:"id" example:"1"`               // 默认数据ID
	Name   string `json:"name" example:"账号管理"`          // 权限名称，界面展示，建议与界面导航一致
	Alias  string `json:"alias" example:"plat_account"` // 权限别名，英文+下划线
	Level  string `json:"level" example:"1"`            // 权限等级 1分组（一级导航）2模块（页面）3功能（按钮）
	Pid    uint64 `json:"pid" example:"1"`              // 父级ID，默认为1
	Static string `json:"static" example:"1"`           // 默认启用权限，1 不启 2启用
	Sort   uint16 `json:"sort" example:"1"`             // 权限排序（独立接口）
	Common
	RouterIds  []uint64  `json:"routerIds" gorm:"-" example:"1,2,3"` // 路由集，当前对象会忽略此字段
	RouterList []*Router `json:"routerList"`                         // 路由对象
}

// PermissionTable 权限泛型构造器
var PermissionTable actuator.Table[Permission]

// DataBase 实现指定数据库
func (t Permission) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Permission) TableName() string {
	return "permission"
}

// FindByIds 根据IDS获取权限别名
func (t Permission) FindByIds(ids []uint64) (res []*Permission, err error) {
	r := platDb.Where("id IN ?", ids).Find(&res)
	err = r.Error
	return
}

// SortPermission 事务排序
func (t Permission) SortPermission(req []*Permission) error {
	// 启用事务
	return platDb.Transaction(func(tx *gorm.DB) error {
		for _, item := range req {
			// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
			if err := tx.Model(&t).Where("id = ?", item.ID).Update("sort", item.Sort).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}
		}
		// 返回 nil 提交事务
		return nil
	})
}

// PermissionArray 权限自定义排序
type PermissionArray []*Permission

func (p PermissionArray) Len() int {
	return len(p)
}

func (p PermissionArray) Less(i, j int) bool {
	return p[i].Sort < p[j].Sort
}

func (p PermissionArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
