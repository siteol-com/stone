package platService

import (
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/mysql/platDb"
	"siteOl.com/stone/server/src/utils/log"
	"time"
)

// InsertLoginRecord 插入一条登录记录
func InsertLoginRecord(accountId, tenantId uint64, loginType string, now time.Time, token, traceID string) {
	insert := &platDb.LoginRecord{
		AccountId: accountId,
		TenantId:  tenantId,
		LoginType: loginType,
		Token:     token,
		LoginTime: &now,
		Common: platDb.Common{
			Status:   constant.StatusOpen,
			CreateAt: &now,
			UpdateAt: &now,
		},
	}
	err := platDb.LoginRecordTable.InsertOne(insert)
	if err != nil {
		// 登陆数据插入异常不影响登陆
		log.ErrorTF(traceID, "InsertLoginRecord Fail . AccountId is %d . Err is %v", accountId, err)
	}
}
