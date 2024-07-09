package middleware

import (
	helper "blog/common/helper/jwt"
	"blog/common/statuscode"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			statuscode.ResponseErrAuthErr(c, statuscode.ERROR_AUTH, "", nil)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		claim, err := helper.AnalyseToken(tokenString)
		if err != nil {
			statuscode.ResponseErrAuthErr(c, statuscode.ERROR_AUTH_CHECK_TOKEN_FAIL, "", nil)
			c.Abort()
			return
		} else if time.Now().Unix() > claim.ExpiresAt.Unix() {
			statuscode.ResponseErrAuthErr(c, statuscode.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, "", nil)
			c.Abort()
			return
		}

		//无感刷新token
		//不安全，考虑其他做法
		if claim.ExpiresAt.Unix()-time.Now().Unix() <= claim.RefreshToken {
			claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 20))
		}

		c.Set("username", claim.Username)
		c.Set("userid", claim.UserId)

		c.Next()
	}
}
