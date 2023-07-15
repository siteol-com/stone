package platService

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/utils/log"
)

// 获取用户的联动权限
func getUserDataPermission(deptId uint64, permissionType string, traceID string) (string, []uint64) {
	var dept *platDb.Dept
	// 如果数据权限跟随部门，读取部门数据权限
	if permissionType == constant.PermissionTypeDeptFellow {
		dept, err := platDb.DeptTable.FindOneById(deptId)
		if err != nil || dept.Status != constant.StatusOpen {
			log.WarnTF(traceID, "GetUserDataPermission DeptId %d Query Fail Or Not Open .", deptId)
			// 部门查询失败，部门状态不正确不具备任何数据权限
			return constant.PermissionTypeNull, nil
		}
		// 部门数据权限赋予用户
		permissionType = dept.PermissionType
	}
	// 如果数据权限为本级及子集，递归查询相关部门
	deptIds := make([]uint64, 0)
	if permissionType == constant.PermissionTypeDeptGroup {
		deptIds = recursionDept(dept, deptIds)
	} else {
		// 反之仅当前部门
		deptIds = []uint64{deptId}
	}
	return permissionType, deptIds
}

// 递归部门
func recursionDept(dept *platDb.Dept, deptIds []uint64) []uint64 {
	// 部门状态不正确不再递归
	if dept.Status != constant.StatusOpen {
		return deptIds
	}
	// 满足先加入数组
	deptIds = append(deptIds, dept.Id)
	// 查询子部门
	childDeptList, err := platDb.DeptTable.FindByObject(&platDb.Dept{Pid: dept.Id})
	// 查询出错或者不再有子部门，结束递归
	if err != nil || len(childDeptList) < 1 {
		return deptIds
	}
	// 遍历子部门继续递归
	for _, childDept := range childDeptList {
		deptIds = recursionDept(childDept, deptIds)
	}
	return deptIds
}
