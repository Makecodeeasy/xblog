package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1xxx 表示用户模块的错误

	USERNAME_USED      = 1001
	PASSWORD_WRONG     = 1002
	USER_NOT_EXIST     = 1003
	TOKEN_NOT_EXIST    = 1004
	TOKEN_EXPIRED      = 1005
	TOKEN_WRONG        = 1006
	TOKEN_FORMAT_WRONG = 1007
	USER_NO_PERMISSION = 1008
	// code = 2xxx 表示文章模块的错误
	ARTICAL_NOT_EXIST = 2001

	// code = 3xxx 表示分类模块的错误
	CATEGORY_NOT_EXIST = 3001
	CATEGORY_USED      = 3002
)

var codeMsg = map[int]string{
	SUCCESS:            "ok",
	ERROR:              "FAIL",
	USERNAME_USED:      "用户名已存在",
	USER_NOT_EXIST:     "用户名不存在",
	TOKEN_NOT_EXIST:    "TOKEN不存在",
	TOKEN_EXPIRED:      "TOKEN已过期",
	TOKEN_WRONG:        "TOKEN不正确",
	TOKEN_FORMAT_WRONG: "TOKEN格式不正确",
	ARTICAL_NOT_EXIST:  "文章不存在",
	CATEGORY_USED:      "分类已存在",
	CATEGORY_NOT_EXIST: "分类不存在",
	USER_NO_PERMISSION: "用户没有权限",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
