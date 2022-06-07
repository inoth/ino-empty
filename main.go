package main

import (
	"defaultProject/components/cache"
	"defaultProject/components/config"
	"defaultProject/components/db"
	"defaultProject/components/httpsvc"
	"defaultProject/components/logger"
	"defaultProject/register"
	"defaultProject/src/router"
	"fmt"
	"os"
)

func main() {
	// 注册组件
	err := register.Register(
		&cache.CacheComponents{}, // 本地缓存
		config.Instance(),        // 配置文件
		// &logger.ZapLogger{},       // zap 日志
		&logger.LogrusConfig{},    // logrus日志
		&db.RedisConnectCluster{}, // redis 数据库
		&db.MysqlConnect{},        // mysql 数据库
	).Init().Run(
		httpsvc.NewGinConfig(config.Cfg.GetString("ServerPort")).
			SetRouter(
				&router.ProjectRouter{},
				&router.Project2Router{},
			))
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
