package v1

import (
	"blog/common/statuscode"
	"blog/model"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func AddTag(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		statuscode.ResponseOK(c, statuscode.INVALID_PARAMS, ":标签名不能为空", nil)
		return
	}
	count, ok := model.GetTagCountByName(name)
	if !ok {
		statuscode.ResponseOK(c, statuscode.ERROR, "", nil)
		return
	} else if count != 0 {
		statuscode.ResponseOK(c, statuscode.ERROR_EXIST_TAG, "", nil)
		return
	}

	createdBy, ok := c.Get("username")
	if !ok {
		statuscode.ResponseOK(c, statuscode.ERROR, "", nil)
		return
	}

	status := c.DefaultPostForm("status", "0")
	if status == "" {
		statuscode.ResponseOK(c, statuscode.INVALID_PARAMS, ":状态不能为空", nil)
		return
	}
	statusInt := com.StrTo(status).MustInt()

	if ok = model.AddTag(name, statusInt, createdBy.(string)); !ok {
		statuscode.ResponseOK(c, statuscode.ERROR, "", nil)
		return
	}

	statuscode.ResponseOK(c, statuscode.OK, "", nil)
}

func DeleteTag(c *gin.Context) {
	tagName := c.PostForm("name")
	if tagName == "" {
		statuscode.ResponseOK(c, statuscode.INVALID_PARAMS, ":标签名不能为空", nil)
		return
	}
	count, ok := model.GetTagCountByName(tagName)
	if !ok {
		statuscode.ResponseOK(c, statuscode.ERROR, "", nil)
		return
	} else if count == 0 {
		statuscode.ResponseOK(c, statuscode.ERROR_NOT_EXIST_TAG, "", nil)
		return
	}

	ok = model.DeleteTagByTagName(tagName)
	if !ok {
		statuscode.ResponseOK(c, statuscode.ERROR, "", nil)
		return
	}

	statuscode.ResponseOK(c, statuscode.OK, "：已成功删除指定数据", nil)
}

func UpdateTag(c *gin.Context) {

}
