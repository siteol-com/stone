package redis

import (
	"os"
	"siteOl.com/stone/server/src/data/config"
	"siteOl.com/stone/server/src/utils/log"

	"github.com/go-redis/redis"
)

var cluster *redis.Client

// Init Redis初始化
func Init() {
	// redis
	cluster = redis.NewClient(&redis.Options{
		Addr:     config.JsonConfig.Redis.Addr,
		DB:       config.JsonConfig.Redis.DB,
		Password: config.JsonConfig.Redis.Password,
	})
	err := cluster.Ping().Err()
	if err != nil {
		log.ErrorF("Init Redis Fail . Err : %s", err)
		os.Exit(1)
	} else {
		log.InfoF("Init Redis success")
	}
}
