package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, msg string, data interface{}) {
	if data == nil {
		c.JSON(http.StatusOK, resultOK(msg))
	} else {
		c.JSON(http.StatusOK, ok(msg, data))
	}
}

func Err(c *gin.Context, msg string) {
	c.JSON(FAILED, err(msg))
}

func NotFound(c *gin.Context, msg string) {
	c.JSON(NOTFOUND, notFound(msg))
}

func ParamErr(c *gin.Context, msg string) {
	c.JSON(PARAMETERERR, paramErr(msg))
}

func Unauthrized(c *gin.Context, msg string) {
	c.JSON(UNAUTHORIZATION, unauthrized(msg))
}
