package router

import "github.com/gin-gonic/gin"

type ProjectRouter struct{}

func (ProjectRouter) Load(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.String(200, "hello world project")
	})
}

type Project2Router struct{}

func (Project2Router) Load(g *gin.Engine) {
	g.GET("/v1", func(c *gin.Context) {
		c.String(200, "hello world project2")
	})
}
