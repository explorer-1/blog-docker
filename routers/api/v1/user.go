package v1

import (
	helper "blog/common/helper/jwt"
	"blog/common/statuscode"
	"blog/model"

	"github.com/gin-gonic/gin"
)

// Login
// @Tags 公共方法
// @Summary 用户登录
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {string} json "{"code":"200", "msg","","data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	if username == "" {
		statuscode.ResponseOK(c, statuscode.INVALID_PARAMS, ":用户名不能为空", nil)
		return
	}

	count, ok := model.GetUserCountByUsername(username)
	if !ok {
		statuscode.ResponseOK(c, statuscode.ERROR, "", nil)
		return
	} else if count == 0 {
		statuscode.ResponseOK(c, statuscode.ERROR_NOT_EXIST_USERNAME, "", nil)
		return
	}

	password := c.PostForm("password")
	if username == "" {
		statuscode.ResponseOK(c, statuscode.INVALID_PARAMS, ":密码不能为空", nil)
		return
	}
	expassword, ok := model.GetUserPasswordByUsername(username)
	if !ok {
		statuscode.ResponseOK(c, statuscode.ERROR, "", nil)
		return
	} else if expassword != password {
		statuscode.ResponseOK(c, statuscode.INVALID_PARAMS, ":密码错误", nil)
		return
	}

	//登陆成功，发放token
	id, ok := model.GetUserIdByUsername(username)
	if !ok {
		statuscode.ResponseOK(c, statuscode.ERROR, "", nil)
		return
	}

	tokenString, err := helper.GenerateToken(username, id)
	if err != nil {
		statuscode.ResponseOK(c, statuscode.ERROR_AUTH_GENERATE_TOKEN, "", nil)
		return
	}

	statuscode.ResponseOK(c, statuscode.OK, "", tokenString)
}
