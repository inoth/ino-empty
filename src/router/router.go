package router

import (
	"<project-name>/config"
	ex "<project-name>/exception"
	mid "<project-name>/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerStar() {
	r := gin.New()
	r.Use(ex.ExceptionHandle)

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.Group("api", mid.AuthMiddleware)

	r.Run(config.Instance().ServerPort)
}
