package router

import (
	"ino-empty/config"
	ex "ino-empty/exception"
	mid "ino-empty/middleware"
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
