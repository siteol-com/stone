package platService

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/mysql/actuator"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/redis"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/utils/log"
	"strings"
)

// PageRouter 查询路由分页
func PageRouter(traceID string, req *platModel.RouterPageReq) resp.ResBody {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query := actuator.InitQuery()
	if req.Name != "" {
		query.Like("name", req.Name)
	}
	if req.Url != "" {
		query.Like("url", req.Url)
	}
	if req.ServiceCode != "" {
		query.Eq("service_code", req.ServiceCode)
	}
	if req.Type != "" {
		query.Eq("type", req.Type)
	}
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	// 查询分页
	total, list, err := platDb.RouterTable.Page(query)
	if err != nil {
		log.ErrorTF(traceID, "PageRouter Fail . Err Is : %v", err)
		return resp.Error()
	}
	return resp.SuccessUnPop(model.SetPageRes(list, total))
}

// AddRouter 创建路由
func AddRouter(traceID string, req *platDb.Router) resp.ResBody {
	req.Id = 0
	err := platDb.RouterTable.InsertOne(req)
	if err != nil {
		log.ErrorTF(traceID, "PageRouter Fail . Err Is : %v", err)
		return checkRouterDBErr(err)
	}
	// 刷新白名单路由缓存
	go InitWhiteRouterCache(traceID)
	// 路由创建成功
	return resp.SuccessWithCode("2003000", true)
}

// GetRouter 查询路由
func GetRouter(traceID string, req *model.IdReq) resp.ResBody {
	router, err := platDb.RouterTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetRouter By Id %d Fail . Err Is : %v", req.Id, err)
		// 路由查询失败
		return resp.Fail("5003000")
	}
	// 路由创建成功
	return resp.SuccessUnPop(router)
}

// EditRouter 编辑路由
func EditRouter(traceID string, req *platDb.Router) resp.ResBody {
	if req.Id == 0 {
		// 路由不存在 路由查询失败
		return resp.Fail("5003000")
	}
	_, err := platDb.RouterTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetRouter By Id %d Fail . Err Is : %v", req.Id, err)
		// 路由查询失败
		return resp.Fail("5003000")
	}
	// 推送关联
	var needDel bool
	if req.Type == constant.RouterTypeWhite {
		// 如果修改后是白名单数据，需要尝试移除可能绑定的权限
		needDel = true
	}
	// 通知权限刷新，尝试移除可能关联的权限，并可能通知账号权限变动
	err = notifyChangeByRouter(req.Id, needDel, traceID)
	if err != nil {
		log.ErrorTF(traceID, "notifyChangeByRouter By Id %d Fail . Err Is : %v", req.Id, err)
		// 移除路由权限关系失败
		return resp.Error()
	}
	// 更新数据
	err = platDb.RouterTable.UpdateOne(req)
	if err != nil {
		log.ErrorTF(traceID, "EditRouter By Id %d Fail . Err Is : %v", req.Id, err)
		return checkRouterDBErr(err)
	}
	// 刷新白名单缓存
	go InitWhiteRouterCache(traceID)
	// 路由编辑成功
	return resp.SuccessWithCode("2003001", true)
}

// DelRouter 删除路由
func DelRouter(traceID string, req *model.IdReq) resp.ResBody {
	router, err := platDb.RouterTable.FindOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetRouter By Id %d Fail . Err Is : %v", req.Id, err)
		// 路由查询失败
		return resp.Fail("5003000")
	}
	// 如果删除的是授权路由，尝试移除可能关联的权限，并可能通知账号权限变动
	if router.Type == constant.RouterTypeAuth {
		err = notifyChangeByRouter(req.Id, true, traceID)
		if err != nil {
			log.ErrorTF(traceID, "notifyChangeByRouter By Id %d Fail . Err Is : %v", req.Id, err)
			// 移除路由权限关系失败
			return resp.Error()
		}
	}
	err = platDb.RouterTable.DeleteOne(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelRouter By Id %d Fail . Err Is : %v", req.Id, err)
		return resp.Error()
	}
	// 刷新白名单缓存
	go InitWhiteRouterCache(traceID)
	// 路由删除成功
	return resp.SuccessWithCode("2003002", true)
}

// 转换数据库错误
func checkRouterDBErr(err error) resp.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "url_uni") {
			// URL 不可重复
			return resp.Fail("5003001")
		}
		if strings.Contains(errStr, "name_uni") {
			// 路由名称不可重复
			return resp.Fail("5003002")
		}
	}
	// 默认500
	return resp.Error()
}

// InitWhiteRouterCache 初始化免授权路由列表
func InitWhiteRouterCache(traceID string) {
	routers, err := platDb.RouterTable.FindByObject(&platDb.Router{Type: constant.RouterTypeWhite})
	if err != nil {
		log.ErrorTF(traceID, "InitWhiteRouterCache Fail . Err Is : %v", err)
		return
	}
	if len(routers) > 0 {
		urls := make([]string, len(routers))
		for i, item := range routers {
			urls[i] = item.Url
		}
		// 写入缓存（持久生效）
		err = redis.Set(constant.CacheKeyRouterWhite, urls, 0)
		if err != nil {
			log.ErrorTF(traceID, "SetWhiteRouterCache Fail . Err Is : %v", err)
		}
	}
}
