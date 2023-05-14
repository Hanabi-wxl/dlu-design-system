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
	LimitRequestErrno    = &Errno{LimitRequestErrCode, LimitRequestErrMsg}
	ResetPasswordErr     = &Errno{ResetPasswordErrCode, ResetPasswordErrMsg}
	UpdateTeacherRoleErr = &Errno{UpdateTeacherRoleErrCode, UpdateTeacherRoleErrMsg}
	DeleteUserErr        = &Errno{DeleteUserErrCode, DeleteUserErrMsg}
	UpdateUserErr        = &Errno{UpdateUserErrCode, UpdateUserErrMsg}
	DeleteRoleErr        = &Errno{DeleteRoleErrCode, DeleteRoleErrMsg}
	UpdateRoleErr        = &Errno{UpdateRoleErrCode, UpdateRoleErrMsg}
	DeleteManagerErr     = &Errno{DeleteManagerErrCode, DeleteManagerErrMsg}
)

const (
	DefaultErrCode      = 20001    // 默认异常状态码
	DefaultErrMsg       = "failed" // 默认异常信息
	LimitRequestErrCode = 20002
	LimitRequestErrMsg  = "请求过于频繁请稍后再试"

	NumberNotFoundErrCode    = 30003
	NumberNotFoundErrMsg     = "学号/工号不存在"
	NumberDuplicateErrCode   = 30004
	NumberDuplicateErrMsg    = "学号/工号重复"
	PasswordIncorrectErrCode = 300005
	PasswordIncorrectErrMsg  = "用户名或密码错误"
	UnAuthorizationErrCode   = 300006
	UnAuthorizationErrMsg    = "登陆已过期"
	NoPermissionErrMsg       = "无权限"
	ResetPasswordErrCode     = 300007
	ResetPasswordErrMsg      = "修改密码失败"
	UpdateTeacherRoleErrCode = 300008
	UpdateTeacherRoleErrMsg  = "修改教师角色失败"
	DeleteUserErrCode        = 300009
	DeleteUserErrMsg         = "删除用户失败"
	UpdateUserErrCode        = 300010
	UpdateUserErrMsg         = "更新用户失败"
	DeleteRoleErrCode        = 30011
	DeleteRoleErrMsg         = "删除角色失败"
	UpdateRoleErrCode        = 30012
	UpdateRoleErrMsg         = "更新角色失败"
	DeleteManagerErrCode     = 30013
	DeleteManagerErrMsg      = "删除管理员失败"
)

const (
	DbErrorCode = 30002 // 数据库异常状态码
)
