package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"siteOl.com/stone/server/src/data/redis"
	"siteOl.com/stone/server/src/sevices/plat/platService"
	"syscall"
	"time"

	"siteOl.com/stone/server/src/data/mysql"

	"siteOl.com/stone/server/src/utils/comm"
	"siteOl.com/stone/server/src/utils/log"

	"siteOl.com/stone/server/src/data/config"
	"siteOl.com/stone/server/src/router"
)

// @title			Stone
// @version         1.0

// @description   	物联网基座Stone，提供一个多层级SaaS化的基础开箱即用中台管理服务。
// @description   	# 概述
// @description   	该接口文档提供Swagger[支持调试]和ReDoc[阅读增强]两个版本。
// @description
// @description   	[Swagger[支持调试]](/docs/swagger/index.html) 丨 [ReDoc[阅读增强]](/docs/redoc/index.html)
// @description
// @description   	# API说明
// @description   	本系统的全部接口采用【POST】【application/json】方式传输数据。
// @description
// @description   	除开放接口以外的其他接口均需要通过【ApiKeyAuth:请求头[Token]】完成鉴权。

// @contact.name 	Stone
// @contact.url		https://stone.siteol.com
// @contact.email	stone@siteol.com
// @host			127.0.0.1:8000
// @BasePath  		/
// @accept			json

// @securityDefinitions.apikey 	Token
// @in 							header
// @name 						Token

// @x-logo {"url" :"/docs/sc/logo.png","altText":"Stone"}

// @tag.name 开放接口
// @tag.description 基础开发接口
// @tag.name 平台
// @tag.description 基础开发接口

// @x-tagGroups [{ "name": "基础", "tags": ["开放接口"]}]

// 主函数
func main() {
	// 初始化数据库
	mysql.Init()
	// 初始化Redis
	redis.Init()
	// 业务初始化
	serviceInit()
	// 初始化路由
	router := router.NewRouter()
	httpServer := &http.Server{Addr: config.JsonConfig.Server.Port, Handler: router}
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
