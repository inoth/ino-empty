package router

import "github.com/gin-gonic/gin"

type ProjectRouter struct{}

func (ProjectRouter) Load(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})
}
