package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type ManageService interface {
	// SaveUser
	// @Description: 添加用户 学生或老师
	// @param userReq 内容
	// @return *errno.Errno 错误
	SaveUser(userReq UserRequest) *errno.Errno
}
