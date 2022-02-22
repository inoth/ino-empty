package main

import (
	"defaultProject/cache"
	"defaultProject/config"
	"defaultProject/db"
	"defaultProject/global"
	"defaultProject/httpsvc"
	"defaultProject/logger"
	"defaultProject/queue"
	"defaultProject/src/router"
)

func main() {
	global.Register(
		&config.ViperConfig{},
		&logger.LogrusConfig{},
		&cache.RedisCache{},
		&db.MysqlConnect{},
		&db.MongoConnect{},
		&queue.NsqQueue{},
	).Init().ServeStart(httpsvc.NewGinConfig().SetRouter(&router.ProjectRouter{}))
}
