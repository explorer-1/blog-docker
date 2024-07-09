package statuscode

const (
	OK             = 400
	ERROR          = 200
	INVALID_PARAMS = 201

	ERROR_EXIST_TAG         = 202
	ERROR_NOT_EXIST_TAG     = 203
	ERROR_NOT_EXIST_ARTICLE = 204

	ERROR_AUTH_GENERATE_TOKEN      = 301
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 302
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 303
	ERROR_AUTH                     = 304

	ERROR_NOT_EXIST_USERNAME = 501
	ERROR_PASSWORD           = 502
)

var statusMsg = map[int]string{
	OK:             "请求成功",
	ERROR:          "内部错误",
	INVALID_PARAMS: "请求参数错误",

	ERROR_EXIST_TAG:                "标签已存在",
	ERROR_NOT_EXIST_TAG:            "标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "文章不存在",
	ERROR_AUTH_GENERATE_TOKEN:      "token生成失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token超时",
	ERROR_AUTH:                     "token错误",
	ERROR_NOT_EXIST_USERNAME:       "用户名不存在",
	ERROR_PASSWORD:                 "密码错误",
}

func GetStatusMsg(code int) string {
	msg, ok := statusMsg[code]
	if ok {
		return msg
	}

	return statusMsg[ERROR]
}
