package util

//获取页码

import (
	"blog/common/confsetting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	page, _ := com.StrTo(c.Query("page")).Int()
	// if err != nil {

	// }

	var result int
	if page > 0 {
		result = (page - 1) * confsetting.PageSize
	}
	return result
}
