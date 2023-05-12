package errno

type Errno struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func NewErrno(code int, msg string) *Errno {
	return &Errno{DefaultErrCode, msg}
}

func NewDefaultErrno(msg string) *Errno {
	return &Errno{DefaultErrCode, msg}
}

var (
	LimitRequestErrno = Errno{LimitRequestErrCode, LimitRequestErrMsg}
)

const (
	DefaultErrCode = 20001    // 默认异常状态码
	DefaultErrMsg  = "failed" // 默认异常信息

	LimitRequestErrCode      = 20002
	LimitRequestErrMsg       = "请求过于频繁请稍后再试"
	NumberNotFoundErrCode    = 20003
	NumberNotFoundErrMsg     = "学号/工号不存在"
	NumberDuplicateErrCode   = 20004
	NumberDuplicateErrMsg    = "学号/工号重复"
	PasswordIncorrectErrCode = 200005
	PasswordIncorrectErrMsg  = "用户名或密码错误"
	UnAuthorizationErrCode   = 200006
	UnAuthorizationErrMsg    = "登陆已过期"
)

const (
	DbErrorCode = 30002 // 数据库异常状态码
)
