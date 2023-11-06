package platService

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model/platModel"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/utils/log"
	"time"
)

// GetOpenTenant 获取租户信息（开放），失败不返回具体信息
func GetOpenTenant(traceID string, req *platModel.OpenTenantReq) *resp.ResBody {
	// 获取租户查询结构体
	tenant, err := platDb.TenantTable.FindOneByObject(&platDb.Tenant{Alias: req.TenantAlias})
	if err != nil {
		log.ErrorTF(traceID, "GetOpenTenant Fail . Err is %v", err)
		return resp.Fail(constant.TenantGetNG) // 租户查询失败
	}
	// 检查租户，检查不通过
	check, checkRes := CheckTenant(&tenant)
	if !check {
		return checkRes
	}
	// 响应安全结构体 租户信息获取成功
	return resp.SuccessWithCode(constant.TenantGetOK, platModel.OpenTenantRes{
		Name:  tenant.Name,
		Alias: tenant.Alias,
		Theme: tenant.Theme,
		Logo:  tenant.Logo,
		Icon:  tenant.Icon,
	})
}

// CheckTenant 检查租户可用性
func CheckTenant(tenant *platDb.Tenant) (check bool, res *resp.ResBody) {
	if tenant.Status != constant.StatusOpen {
		res = resp.Fail(constant.TenantStatusNG) // 租户状态不可用
		return
	}
	if tenant.ExpiryTime != nil && time.Now().After(*tenant.ExpiryTime) {
		res = resp.Fail(constant.TenantExpNG) // 租户已过期
		return
	}
	return true, resp.SuccessUnPop(nil)
}
