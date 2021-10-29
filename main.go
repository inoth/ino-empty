package main

import (
	"<project-name>/cache"
	"<project-name>/src/db"
	"<project-name>/src/router"
)

func main() {
	cache.Init()
	db.Init()
	router.ServerStar()
}
