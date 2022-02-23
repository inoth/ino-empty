package main

import (
	"defaultProject/config"
	"defaultProject/global"
	"defaultProject/httpsvc"
	"defaultProject/logger"
	"defaultProject/src/router"
)

func main() {
	global.Register(
		&config.ViperConfig{},
		&logger.LogrusConfig{},
		// &cache.RedisCache{},
		// &db.MysqlConnect{},
		// &db.MongoConnect{},
		// &queue.NsqQueue{},
	).Init().
		SubServe(httpsvc.NewGinConfig(":10087").SetRouter(&router.ProjectRouter{}, &router.Project2Router{})).
		Run(httpsvc.NewGinConfig(":10086").SetRouter(&router.ProjectRouter{}, &router.Project2Router{}))
}
