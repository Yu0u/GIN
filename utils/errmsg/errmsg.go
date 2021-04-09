package errmsg

const (
	SUCCSE                = 200
	ERROR                 = 400
	INTERNAL_SERVER_ERROR = 500

	USER_EXIST             = 1001
	USER_NOT_EXIST         = 1002
	ERROR_PASSWORD         = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008

	DATA_VALIDATE_ERROR = 2001
)

var codeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	INTERNAL_SERVER_ERROR:  "服务器内部错误",
	USER_EXIST:             "用户已存在",
	USER_NOT_EXIST:         "用户不存在",
	ERROR_PASSWORD:         "用户名或密码错误",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:    "用户无权限",
	DATA_VALIDATE_ERROR:    "数据校验错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
