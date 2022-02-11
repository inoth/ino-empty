package middleware

import "github.com/gin-gonic/gin"

//type GinGlobalException struct {}

func GinGlobalException() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover().(error); err != nil {
				c.Abort()
				c.JSON(500, gin.H{"code": 500, "msg": err.Error()})
			}
		}()
		c.Next()
	}
}
