package platDb

// Router 路由表
type Router struct {
	Id          uint64 // 默认数据ID
	Name        string // 路由名称，用于界面展示，与权限关联
	Url         string // 路由地址，后端访问URL 后端不再URL中携带参数，统一Post处理内容
	ServiceCode uint8  // 业务编码（字典），为接口分组
	WhiteFlag   uint8  // 免授权路由 1 授权 2 免授权（系统启动开放免授权）

}

// TableName 实现自定义表名
func (t *Router) TableName() string {
	return "router"
}

// FindByIds 获取路由清单
func (t *Router) FindByIds(ids []uint64) (res []*Router, err error) {
	r := platDb.Where("id IN ?", ids).Find(&res)
	err = r.Error
	return
}
