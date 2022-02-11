package httpsvc

import (
	"defaultProject/config"
	"defaultProject/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type GinRouter interface {
	Load(g *gin.Engine)
}

type GinConfig struct {
	ginEngine *gin.Engine
}

func NewGinConfig() *GinConfig {
	g := &GinConfig{
		ginEngine: gin.New(),
	}

	g.ginEngine.Use(
		middleware.GinGlobalException(),
		middleware.Cors(),
		middleware.AuthMiddleware())

	g.ginEngine.MaxMultipartMemory = 10 << 20
	// docs.SwaggerInfo.BasePath = ""
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	g.ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return g
}

func (g *GinConfig) SetRouter(router GinRouter) *GinConfig {
	router.Load(g.ginEngine)
	return g
}

func (g *GinConfig) Init() error {
	return g.run()
}

func (g *GinConfig) run() error {
	return g.ginEngine.Run(config.Cfg.GetString("ServerPort"))
}