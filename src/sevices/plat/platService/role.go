package platService

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/model/platModel"
	"siteOl.com/stone/server/src/data/mysql/actuator"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/utils/log"
	"strings"
	"time"
)

// PageRole 查询角色分页
func PageRole(traceID string, req *platModel.RolePageReq) *resp.ResBody {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query := actuator.InitQuery()
	if req.Name != "" {
		query.Like("name", req.Name)
	}
	// TODO 过滤租户ID
	// 仅查询未被封存的角色
	query.Lt("status", constant.StatusClose)
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	// 查询分页
	total, list, err := platDb.RoleTable.Page(query)
	if err != nil {
		log.ErrorTF(traceID, "PageRole Fail . Err Is : %v", err)
		return resp.ResFail
	}
	return resp.SuccessUnPop(model.SetPageRes(list, total))
}

// AddRole 创建角色
func AddRole(traceID string, req *platModel.RoleAddReq) *resp.ResBody {
	dbReq := platModel.RoleReqToDbReq(req)
	err := platDb.RoleTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddRole Fail . Err Is : %v", err)
		return checkRoleDBErr(err)
	}
	// 插入角色权限集
	if len(dbReq.PermissionIds) > 0 {
		err := refreshRolePermissions(dbReq, false, traceID)
		if err != nil {
			return checkRoleDBErr(err)
		}
	}
	// 角色创建成功
	return resp.SuccessWithCode(constant.RoleAddOK, true)
}

// GetRole 查询角色 tenantId（控制层补充）
func GetRole(traceID string, req *model.IdAnTenantReq) *resp.ResBody {
	role, err := platDb.RoleTable.FindOneByObject(req)
	if err != nil {
		log.ErrorTF(traceID, "GetRole By Id %d Fail . Err Is : %v", req.ID, err)
		// 角色查询失败
		return resp.Fail(constant.RoleGetNG)
	}
	// 角色查询成功
	return resp.SuccessUnPop(role)
}

// EditRole 编辑角色
func EditRole(traceID string, req *platModel.RoleEditReq) *resp.ResBody {
	if req.ID == 0 {
		// 角色不存在 角色查询失败
		return resp.Fail(constant.RoleGetNG)
	}
	res, err := platDb.RoleTable.FindOneByObject(&model.IdAnTenantReq{ID: req.ID, TenantId: req.TenantId})
	if err != nil {
		log.ErrorTF(traceID, "GetRole By Id %d Fail . Err Is : %v", req.ID, err)
		// 角色查询失败
		return resp.Fail(constant.RoleGetNG)
	}
	// 禁止变更
	if res.Mark == constant.StatusLock {
		// 角色禁止变更
		return resp.Fail(constant.RoleEditLockNG)
	}
	now := time.Now()
	// 仅可修改以下项目
	res.UpdateAt = &now
	res.Name = req.Name
	res.Remark = req.Remark
	// 更新数据
	err = platDb.RoleTable.UpdateOne(res)
	if err != nil {
		log.ErrorTF(traceID, "EditRole By Id %d Fail . Err Is : %v", req.ID, err)
		return checkRoleDBErr(err)
	}
	// 更新角色权限集
	res.PermissionIds = req.PermissionIds
	err = refreshRolePermissions(&res, true, traceID)
	if err != nil {
		return checkRoleDBErr(err)
	}
	// 角色变动通知账号
	go notifyChangeByRoleIds([]uint64{req.ID}, false, traceID)
	// 角色编辑成功
	return resp.SuccessWithCode(constant.RoleEditOK, true)
}

// DelRole 删除角色
func DelRole(traceID string, req *model.IdAnTenantReq) *resp.ResBody {
	role, err := platDb.RoleTable.FindOneByObject(req)
	if err != nil {
		log.ErrorTF(traceID, "GetRole By Id %d Fail . Err Is : %v", req.ID, err)
		// 角色查询失败
		return resp.Fail(constant.RoleGetNG)
	}
	// 禁止变更
	if role.Mark == constant.StatusLock {
		// 角色禁止删除
		return resp.Fail(constant.RoleDelLockNG)
	}
	// 角色是硬删除行为
	err = platDb.RoleTable.DeleteOne(role.ID)
	if err != nil {
		log.ErrorTF(traceID, "DelRole By Id %d Fail . Err Is : %v", req.ID, err)
		return resp.ResFail
	}
	// 角色变动通知账号
	go notifyChangeByRoleIds([]uint64{req.ID}, true, traceID)
	// 角色删除成功
	return resp.SuccessWithCode(constant.RoleDelOK, true)
}

// 转换数据库错误
func checkRoleDBErr(err error) *resp.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "name_tenant_id_uni") {
			// 租户下角色名唯一
			return resp.Fail(constant.RoleUniNameNG)
		}
		if strings.Contains(errStr, "role_permission_uni") {
			// 角色权限 不可重复
			return resp.Fail(constant.RoleUniPermissionNG)
		}
	}
	// 默认500
	return resp.ResFail
}
