package platService

import (
	"fmt"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/mysql/actuator"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/data/redis"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/sevices/plat/platModel"
	"siteOl.com/stone/server/src/utils/log"
	"strconv"
	"strings"
	"time"
)

// PageResponse 查询响应码分页
func PageResponse(traceID string, req *platModel.ResponsePageReq) resp.ResBody {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query := actuator.InitQuery()
	if req.Code != "" {
		query.Like("code", req.Code)
	}
	if req.ServiceCode != "" {
		query.Eq("service_code", req.ServiceCode)
	}
	if req.Type != "" {
		query.Eq("type", req.Type)
	}
	// 仅查询未被封存的响应码
	query.Lt("status", constant.StatusClose)
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	// 查询分页
	total, list, err := platDb.ResponseTable.Page(query)
	if err != nil {
		log.ErrorTF(traceID, "PageResponse Fail . Err Is : %v", err)
		return resp.SysErr
	}
	return resp.SuccessUnPop(model.SetPageRes(list, total))
}

// AddResponse 创建响应码
func AddResponse(traceID string, req *platDb.Response) resp.ResBody {
	req.ID = 0
	now := time.Now()
	req.CreateAt = &now
	req.Status = constant.StatusOpen
	req.Mark = constant.StatusOpen
	// 响应码推算
	err := makeResponseCode(traceID, req)
	if err != nil {
		log.ErrorTF(traceID, "ResponseMakeResponseCode Fail . Err Is : %v", err)
		// 数据库系统异常
		return resp.SysErr
	}
	err = platDb.ResponseTable.InsertOne(req)
	if err != nil {
		log.ErrorTF(traceID, "AddResponse Fail . Err Is : %v", err)
		return checkResponseDBErr(err)
	}
	// 刷新响应码缓存
	go InitResponseCache(traceID)
	// 响应码创建成功
	return resp.SuccessWithCode(constant.ResponseAddOK, true)
}

// 响应码推算
func makeResponseCode(traceID string, req *platDb.Response) (err error) {
	groupCount, err := req.CountByGroup()
	if err != nil {
		return
	}
	serviceCode, _ := strconv.Atoi(req.ServiceCode)
	responseCode := fmt.Sprintf("%s%03d%03d", req.Type, serviceCode, groupCount)
	log.InfoTF(traceID, "ResponseMakeResponseCode Success . Code Is : %s", responseCode)
	req.Code = responseCode
	return
}

// GetResponse 查询响应码
func GetResponse(traceID string, req *model.IdReq) resp.ResBody {
	response, err := platDb.ResponseTable.FindOneById(req.ID)
	if err != nil {
		log.ErrorTF(traceID, "GetResponse By Id %d Fail . Err Is : %v", req.ID, err)
		// 响应码查询失败
		return resp.Fail(constant.ResponseGetNG)
	}
	// 响应码创建成功
	return resp.SuccessUnPop(response)
}

// EditResponse 编辑响应码
func EditResponse(traceID string, req *platDb.Response) resp.ResBody {
	if req.ID == 0 {
		// 响应码不存在 响应码查询失败
		return resp.Fail(constant.ResponseGetNG)
	}
	res, err := platDb.ResponseTable.FindOneById(req.ID)
	if err != nil {
		log.ErrorTF(traceID, "GetResponse By Id %d Fail . Err Is : %v", req.ID, err)
		// 响应码查询失败
		return resp.Fail(constant.ResponseGetNG)
	}
	now := time.Now()
	// 仅可修改以下项目
	res.UpdateAt = &now
	res.ZhCn = req.ZhCn
	res.EnUs = req.EnUs
	res.Remark = req.Remark
	// 更新数据
	err = platDb.ResponseTable.UpdateOne(res)
	if err != nil {
		log.ErrorTF(traceID, "EditResponse By Id %d Fail . Err Is : %v", req.ID, err)
		return checkResponseDBErr(err)
	}
	// 刷新响应码缓存
	go InitResponseCache(traceID)
	// 响应码编辑成功
	return resp.SuccessWithCode(constant.ResponseEditOK, true)
}

// DelResponse 删除响应码
func DelResponse(traceID string, req *model.IdReq) resp.ResBody {
	response, err := platDb.ResponseTable.FindOneById(req.ID)
	if err != nil {
		log.ErrorTF(traceID, "GetResponse By Id %d Fail . Err Is : %v", req.ID, err)
		// 响应码查询失败
		return resp.Fail(constant.ResponseGetNG)
	}
	response.Status = constant.StatusClose
	err = platDb.ResponseTable.UpdateOne(response)
	if err != nil {
		log.ErrorTF(traceID, "DelResponse By Id %d Fail . Err Is : %v", req.ID, err)
		return resp.SysErr
	}
	// 刷新响应码缓存
	go InitResponseCache(traceID)
	// 响应码删除成功
	return resp.SuccessWithCode(constant.ResponseDelOK, true)
}

// 转换数据库错误
func checkResponseDBErr(err error) resp.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "code_uni") {
			// Code 不可重复
			return resp.Fail(constant.ResponseUniCodeNG)
		}
	}
	// 默认500
	return resp.SysErr
}

// InitResponseCache 初始化响应码缓存
func InitResponseCache(traceID string) (err error) {
	allResCodes, err := (&platDb.Response{}).FindAll()
	if err != nil || len(allResCodes) == 0 {
		// 错误或数据库无配置直接返回
		return
	}
	// 组装缓存对象
	resCodeCacheMap := make(map[string]map[string]string, len(allResCodes))
	for _, res := range allResCodes {
		// 遍历支持的语言并写入Map
		langMap := make(map[string]string, len(constant.TransLangSupport))
		for _, lang := range constant.TransLangSupport {
			switch lang {
			case "en-US":
				langMap[lang] = res.EnUs
			case "zh-CN":
				langMap[lang] = res.ZhCn
			}
		}
		resCodeCacheMap[res.Code] = langMap
	}
	// 写入缓存 无超期
	err = redis.Set(constant.CacheKeyTransLang, resCodeCacheMap, 0)
	if err == nil {
		log.InfoTF(traceID, "InitResponseCache Success .")
	}
	return
}
