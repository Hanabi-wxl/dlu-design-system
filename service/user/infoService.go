package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
)

type InfoService interface {
	GetUserById(id int64, isStu int8) (*User, *errno.Errno)
	GetUserByNumber(number string, isStu int8) (*User, *errno.Errno)
	GetUserByNumberMajor(req types.NumberMajorRequest) (*types.UserResp, *errno.Errno)
	GetUsersByMajor(majorId int64, isStu int8, size, num int) (*[]User, *errno.Errno)
	GetUsersByCollege(collegeId int64, isStu int8, size, num int) (*[]User, *errno.Errno)
}
