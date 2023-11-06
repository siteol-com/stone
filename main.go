package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"siteOl.com/stone/server/src/data/redis"
	"siteOl.com/stone/server/src/sevices/plat/platService"
	"siteOl.com/stone/server/src/sevices/router"
	"syscall"
	"time"

	"siteOl.com/stone/server/src/data/mysql"

	"siteOl.com/stone/server/src/utils/comm"
	"siteOl.com/stone/server/src/utils/log"

	"siteOl.com/stone/server/src/data/config"
)

// @title			Stone
// @version         1.0
// @description.markdown
// @contact.name 	Stone
// @contact.url		https://stone.siteol.com
// @contact.email	stone@siteol.com
// @host			localhost:8000
// @BasePath  		/
// @accept			json

// @securityDefinitions.apikey 	Token
// @in 							header
// @name 						Token

// @x-logo {"url" :"/docs/sc/logo.png","altText":"Stone"}

// @tag.name 开放接口
// @tag.description 基础开发接口

// @x-tagGroups [{ "name": "基础", "tags": ["开放接口"]},{ "name": "平台", "tags": ["租户管理","集团部门","角色配置","登陆账号","访问权限","路由接口","响应文言","数据字典"]}]

// 主函数
func main() {
	// 初始化数据库
	mysql.Init()
	// 初始化Redis
	redis.Init()
	// 业务初始化
	serviceInit()
	// 初始化路由
	newRouter := router.NewRouter()
	httpServer := &http.Server{Addr: config.JsonConfig.Server.Port, Handler: newRouter}
	// 启用HTTP服务 - 注册自定义路由
	go comm.RecoverWrap(func() {
		log.InfoTF(fmt.Sprintf("%s%s", config.SysNode, "INIT"), "Server Listening on port %s", config.JsonConfig.Server.Port)
		if err := httpServer.ListenAndServe(); err != nil {
			log.ErrorTF(fmt.Sprintf("%s%s", config.SysNode, "DOWN"), "Server Listening on port %s . Err %v", config.JsonConfig.Server.Port, err)
			os.Exit(1)
		}
	})()
	// 优雅关机
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.ErrorTF(fmt.Sprintf("%s%s", config.SysNode, "DOWN"), "Server Get a signal %s, Stop the consume process", sig.String())
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			// gracefully shutdown with timeout
			_ = httpServer.Shutdown(ctx)
			return
		}
	}
}

// serviceInit 业务初始化
func serviceInit() {
	// 主服务进行响应码初始化
	if config.SysNode == "APP01" {
		err := platService.InitResponseCache(fmt.Sprintf("%s%s", config.SysNode, "INIT"))
		if err != nil {
			log.ErrorTF(fmt.Sprintf("%s%s", config.SysNode, "INIT"), "InitResponseCache Fail . Err is : %v", err)
			os.Exit(1)
		}
	}
}
