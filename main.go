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
	).Init().ServeStart(httpsvc.NewGinConfig().SetRouter(&router.ProjectRouter{}, &router.Project2Router{}))
}
