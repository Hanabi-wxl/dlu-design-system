package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
)

type InfoService interface {
	GetUserById(id int64, isStu int8) (*types.UserResp, *errno.Errno)
	GetUserByNumber(number string, isStu int8) (*types.UserResp, *errno.Errno)
	GetUserByNumberMajor(req *types.NumberMajorRequest) (*types.UserResp, *errno.Errno)
	GetUsersByMajor(req *types.UserQueryByMajorReq) (*[]*types.UserResp, *errno.Errno)
	GetUsersByCollege(req *types.UserQueryByCollegeReq) (*[]*types.UserResp, *errno.Errno)
}
