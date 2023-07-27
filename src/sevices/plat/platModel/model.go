package platModel

import (
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/mysql/platDb"
)

// OpenTenantReq 开放租户查询
type OpenTenantReq struct {
	TenantAlias string `json:"tenantAlias" binding:"required"` // 租戶别名
}

// OpenTenantRes 开放租户响应
type OpenTenantRes struct {
	Name  string `json:"name"`  // 租戶名称
	Alias string `json:"alias"` // 租戶别名
	Theme string `json:"theme"` // 租户模板
	Logo  string `json:"logo"`  // 租户Logo
	Icon  string `json:"icon"`  // 租户Icon
}

// AuthLoginReq 账密登陆结构体
type AuthLoginReq struct {
	Account     string `json:"account" binding:"required"`     // 账号
	Password    string `json:"password" binding:"required"`    // 密码
	TenantAlias string `json:"tenantAlias" binding:"required"` // 租户别名
}

// AuthLoginRes 账密登陆响应
type AuthLoginRes struct {
	Token string `json:"token"` // 登陆Token
}

// AuthUser 授权对象
type AuthUser struct {
	UserId             uint64   `json:"userId"`             // 用户ID
	PwdExpTimeStr      string   `json:"pwdExpTimeStr"`      // 密码超期时间（修改后的90天）
	TenantId           uint64   `json:"tenantId"`           // 租户ID
	DeptId             uint64   `json:"deptId"`             // 部门ID
	PermissionType     string   `json:"permissionType"`     // 权限类型 0全局数据 1跟随部门 2仅本部门 3本部门及子部门 99无权限
	PermissionDeptList []uint64 `json:"permissionDeptList"` // 具有数据权限的部门ID列表
	PermissionList     []string `json:"permissionList"`     // 权限集（前端路由与使用）
	RouterList         []string `json:"routerList"`         // 路由集
}

// DictListReq 字典下拉列表
type DictListReq struct {
	GroupKeys []string `json:"groupKeys"` // 需要查询的字典分组
	Local     string   `json:"-"`         // 字典语言
}

// DictListRes 字典下拉响应
type DictListRes struct {
	Label string `json:"label"` // 字典名
	Value string `json:"value"` // 字典值
}

// DictToListMapRes 内部对象到下拉对象数组
func DictToListMapRes(dictList []*platDb.Dict, local string) ([]*DictListRes, map[string]string) {
	labelList := make([]*DictListRes, len(dictList))
	valueMap := make(map[string]string, len(dictList))
	for i, dict := range dictList {
		labelList[i] = &DictListRes{
			Label: dict.Label,
			Value: dict.Val,
		}
		valueMap[dict.Val] = dict.Label
		// 文言翻译
		switch local {
		case "en":
			labelList[i].Label = dict.LabelEn
			valueMap[dict.Val] = dict.LabelEn
		}
	}
	return labelList, valueMap
}

// RouterPageReq 路由分页查询
type RouterPageReq struct {
	Name        string `json:"name"`                                    // 路由名称
	Url         string `json:"url"`                                     // 路由地址，仅允许提交URI
	ServiceCode string `json:"serviceCode" binding:"omitempty,numeric"` // 业务编码，仅允许提交数字
	Type        string `json:"type" binding:"omitempty,oneof='1' '2'"`  // 路由类型，仅允许提交1/2
	model.PageReq
}

// ResponsePageReq 响应码分页请求
type ResponsePageReq struct {
	Code        string `json:"code"`                                           // 响应码
	ServiceCode string `json:"serviceCode" binding:"omitempty,numeric"`        // 业务编码，仅允许提交数字
	Type        string `json:"responseType" binding:"omitempty,oneof='2' '5'"` // 响应类型，该字段用于筛选，可配置2和5
	model.PageReq
}
