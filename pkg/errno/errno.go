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
	DefaultErrCode      = 20001         // 默认异常状态码
	DefaultErrMsg       = "failed"      // 默认异常信息
	LimitRequestErrCode = 20002         // 限流异常
	LimitRequestErrMsg  = "请求过于频繁请稍后再试" // 限流信息
)

const (
	DbErrorCode = 30002 // 数据库异常状态码
)
