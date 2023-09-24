package platService

import (
	"fmt"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/utils/log"
	"sort"
	"strings"
	"time"
)

// TreePermission 查询权限分页
func TreePermission(traceID string, req *platModel.PermissionBashReq) resp.ResBody {
	if !req.IsSupper {
		// TODO 根据租户ID过滤权限
	}
	// 查询根节点
	rootPerm, err := platDb.PermissionTable.FindOneById(1)
	if err != nil {
		log.ErrorTF(traceID, "TreePermission GetRoot Err %s", err)
		return resp.Fail(constant.PermissionGetNG)
	}
	// 创建树节点
	treeNode := &platModel.Tree{
		Title:    rootPerm.Name,
		Key:      fmt.Sprintf("%d", rootPerm.ID),
		Children: nil,
		Level:    rootPerm.Level,
		Id:       rootPerm.ID,
	}
	recursionPermissionTree(traceID, treeNode, req)
	trees := []*platModel.Tree{treeNode}
	return resp.SuccessUnPop(trees)
}

// AddPermission 创建权限
func AddPermission(traceID string, req *platDb.Permission) resp.ResBody {
	req.ID = 0
	now := time.Now()
	req.CreateAt = &now
	req.Status = constant.StatusOpen
	req.Sort = 0
	// 尝试插入权限
	err := platDb.PermissionTable.InsertOne(req)
	if err != nil {
		log.ErrorTF(traceID, "AddPermission Fail . Err Is : %v", err)
		return checkPermissionDBErr(err)
	}
	// 插入路由关系
	if len(req.RouterIds) > 0 {
		err := refreshPermissionRouters(req, false, traceID)
		if err != nil {
			return checkPermissionDBErr(err)
		}
	}
	// 权限创建成功
	return resp.SuccessWithCode(constant.PermissionAddOK, true)
}

// GetPermission 查询权限
func GetPermission(traceID string, req *model.IdReq) resp.ResBody {
	permission, err := platDb.PermissionTable.FindOneById(req.ID)
	if err != nil {
		log.ErrorTF(traceID, "GetPermission By Id %d Fail . Err Is : %v", req.ID, err)
		// 权限查询失败
		return resp.Fail(constant.PermissionGetNG)
	}
	// 查询关联路由
	routers, _ := getPermissionRouters([]uint64{permission.ID}, traceID)
	// 关联routerIds
	if len(routers) > 0 {
		routerIds := make([]uint64, len(routers))
		for i, item := range routers {
			routerIds[i] = item.ID
		}
		permission.RouterIds = routerIds
	}
	res := platModel.PermissionGetRes{Permission: &permission, RouterList: routers}
	return resp.SuccessUnPop(res)
}

// EditPermission 编辑权限
func EditPermission(traceID string, req *platDb.Permission) resp.ResBody {
	permission, err := platDb.PermissionTable.FindOneById(req.ID)
	if err != nil {
		log.ErrorTF(traceID, "GetPermission By Id %d Fail . Err Is : %v", req.ID, err)
		// 权限查询失败
		return resp.Fail(constant.PermissionGetNG)
	}
	now := time.Now()
	// 可编辑字段
	permission.Name = req.Name
	permission.Alias = req.Alias
	permission.UpdateAt = &now
	// 更新数据
	err = platDb.PermissionTable.UpdateOne(permission)
	if err != nil {
		log.ErrorTF(traceID, "EditPermission By Id %d Fail . Err Is : %v", req.ID, err)
		return checkPermissionDBErr(err)
	}
	// 刷新关联的路由
	if len(req.RouterIds) > 0 {
		err := refreshPermissionRouters(req, true, traceID)
		if err != nil {
			return checkPermissionDBErr(err)
		}
	}
	// 通知权限变动，异步通知受影响的Token失效（上层不涉及数据删除等行为）
	go notifyChangeByPermissionIds([]uint64{req.ID}, false, traceID)
	return resp.SuccessWithCode(constant.PermissionEditOK, true)
}

// DelPermission 删除权限
func DelPermission(traceID string, req *model.IdReq) resp.ResBody {
	permission, err := platDb.PermissionTable.FindOneById(req.ID)
	if err != nil {
		log.ErrorTF(traceID, "GetPermission By Id %d Fail . Err Is : %v", req.ID, err)
		// 权限查询失败
		return resp.Fail(constant.PermissionGetNG)
	}
	// 检查子集
	children, err := platDb.PermissionTable.FindByObject(platDb.Permission{Pid: permission.ID})
	if err != nil {
		log.ErrorTF(traceID, "FindPermissionChild Fail . PID %d . Err is : %s", permission.ID, err)
		// 权限查询失败
		return resp.Fail(constant.PermissionGetNG)
	}
	if len(children) > 0 {
		return resp.Fail(constant.PermissionDelChildNG)
	}
	// 先移除绑定关系
	err = refreshPermissionRouters(&platDb.Permission{ID: permission.ID}, true, traceID)
	if err != nil {
		// TODO
		return resp.Fail("")
	}
	err = platDb.PermissionTable.DeleteOne(req.ID)
	if err != nil {
		log.ErrorTF(traceID, "DelPermission By Id %d Fail . Err Is : %v", req.ID, err)
		return resp.SysErr
	}
	// 通知权限变动，异步通知受影响的Token失效（上层不涉及数据删除等行为）
	go notifyChangeByPermissionIds([]uint64{req.ID}, false, traceID)
	return resp.SuccessWithCode(constant.PermissionDelOK, true)
}

// BroPermission 获取兄弟权限
func BroPermission(traceID string, req *model.IdReq) resp.ResBody {
	permission, err := platDb.PermissionTable.FindOneById(req.ID)
	if err != nil {
		log.ErrorTF(traceID, "GetPermission By Id %d Fail . Err Is : %v", req.ID, err)
		// 权限查询失败
		return resp.Fail(constant.PermissionGetNG)
	}
	// 检查子集
	var bros platDb.PermissionArray
	bros, err = platDb.PermissionTable.FindByObject(platDb.Permission{Pid: permission.Pid})
	if err != nil {
		log.ErrorTF(traceID, "FindPermissionChild Fail . PID %d . Err is : %s", permission.ID, err)
		// 权限查询失败
		return resp.Fail(constant.PermissionGetNG)
	}
	// 数据排序
	sort.Sort(bros)
	return resp.SuccessUnPop(bros)
}

// SortPermission 权限排序
func SortPermission(traceID string, req *[]*model.SortReq) resp.ResBody {
	err := platDb.Permission{}.SortPermission(*req)
	if err != nil {
		log.ErrorTF(traceID, "SortPermission Fail .  Err is : %s", err)
		// 权限排序失败
		return resp.Fail(constant.PermissionSortNG)
	}
	return resp.SuccessWithCode(constant.PermissionSortOK, true)
}

// 转换数据库错误
func checkPermissionDBErr(err error) resp.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "alias_uni") {
			// 别名 不可重复
			return resp.Fail(constant.PermissionUniNameNG)
		}
		if strings.Contains(errStr, "name_uni") {
			// 权限名 不可重复
			return resp.Fail(constant.PermissionUniAliasNG)
		}
		if strings.Contains(errStr, "permission_router_uni") {
			// 权限路由 不可重复
			return resp.Fail(constant.PermissionUniRouterNG)
		}
	}
	// 默认500
	return resp.SysErr
}

// recursionPermissionTree 递归权限树
func recursionPermissionTree(traceID string, treeNode *platModel.Tree, req *platModel.PermissionBashReq) (err error) {
	// 没有子级了
	if treeNode.Level == "3" {
		return
	}
	// 查询子集
	var permissionList platDb.PermissionArray
	permissionList, err = platDb.PermissionTable.FindByObject(platDb.Permission{Pid: treeNode.Id})
	if err != nil {
		log.WarnTF(traceID, "RecursionPermissionTree Fail . PID %d . Err is : %s", treeNode.Id, err)
		return
	}
	if len(permissionList) == 0 {
		// 沒有子集推出
		return
	}
	// 数据排序
	sort.Sort(permissionList)
	treeNode.Children = make([]*platModel.Tree, 0)
	// 组装子集
	for _, item := range permissionList {
		// 是否排除权限，排除后不再向后查,跳过当前循环
		if !req.IsSupper {
			if _, ok := req.IDMap[item.ID]; !ok {
				continue
			}
		}
		// 节点对象
		treeChild := &platModel.Tree{
			Title:    item.Name,
			Key:      fmt.Sprintf("%d", item.ID),
			Children: nil,
			Level:    item.Level,
			Id:       item.ID,
		}
		// 递归子集
		recursionPermissionTree(traceID, treeChild, req)
		// 加入子集
		treeNode.Children = append(treeNode.Children, treeChild)
	}
	return
}
