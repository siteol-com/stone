package main

import (
	"context"
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
		log.InfoF("Server Listening on port %s", config.JsonConfig.Server.Port)
		if err := httpServer.ListenAndServe(); err != nil {
			log.ErrorF("Server Listening on port %s . Err %v", config.JsonConfig.Server.Port, err)
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
			log.ErrorF("Server Get a signal %s, Stop the consume process", sig.String())
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
		err := platService.InitResponseCodeCache()
		if err != nil {
			log.ErrorF("InitResponseCodeCache Fail . Err is : %v", err)
			os.Exit(1)
		}
	}
}
