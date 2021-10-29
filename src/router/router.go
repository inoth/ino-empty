package router

import (
	"<project-name>/config"
	ex "<project-name>/exception"
	mid "<project-name>/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"<project-name>/docs"
)

func ServerStar() {
	r := gin.New()
	r.Use(ex.ExceptionHandle)

	r.MaxMultipartMemory = 10 << 20
	docs.SwaggerInfo.BasePath = ""
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.Group("api", mid.AuthMiddleware)

	r.Run(config.Instance().ServerPort)
}
