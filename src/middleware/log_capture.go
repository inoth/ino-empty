package middleware

import "github.com/gin-gonic/gin"

func LogCapture() gin.HandlerFunc {
	// do something
	return func(c *gin.Context) {
		c.Next()
	}
}
