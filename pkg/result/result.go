package result

import "github.com/Hanabi-wxl/dlu-design-system/pkg/errno"

type Result struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Data       interface{} `json:"data"`
}

func Ok() Result {
	return Result{
		DefaultOkCode, DefaultOkMsg, nil,
	}
}

func Failed() Result {
	return Result{
		errno.DefaultErrCode, errno.DefaultErrMsg, nil,
	}
}

func NewResult(code int, msg string, data interface{}) Result {
	return Result{
		code, msg, data,
	}
}

func NewOkResult(data interface{}) Result {
	return Result{
		DefaultOkCode, DefaultOkMsg, data,
	}
}

func NewFailedResult(msg string) Result {
	return Result{
		errno.DefaultErrCode, errno.DefaultErrMsg, nil,
	}
}

const (
	DefaultOkCode = 10000     // 默认成功状态码
	DefaultOkMsg  = "success" // 默认成功信息
)
