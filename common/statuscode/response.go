package statuscode

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOK(c *gin.Context, statusCode int, s string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": statusCode,
		"msg":  GetStatusMsg(statusCode) + s,
		"data": data,
	})
}

func ResponseErrAuthErr(c *gin.Context, statusCode int, s string, data interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": statusCode,
		"msg":  GetStatusMsg(statusCode) + s,
		"data": data,
	})
}
