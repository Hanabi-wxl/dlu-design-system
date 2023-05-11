package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type InfoService interface {
	UserInfo()
	GetUserById(id, isStu int) (User, *errno.Errno)
	GetUserByNumber(number string, isStu int8) (User, *errno.Errno)
	GetUsersByMajor(majorId, isStu, size, num int) (*[]User, *errno.Errno)
	GetUsersByCollege(collegeId, isStu, size, num int) (*[]User, *errno.Errno)
}
