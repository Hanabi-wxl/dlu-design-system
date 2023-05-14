package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type MajorService interface {
	GetMajorList(size, num int) ([]*model.Major, *errno.Errno)
	AddMajor(major *model.Major) *errno.Errno
	DeleteMajor(id int64) *errno.Errno
	UpdateMajor(major *model.Major) *errno.Errno
	GetMajor(id int64) (*model.Major, *errno.Errno)
	GetMajorListByCollegeId(collegeId int64) ([]*model.Major, *errno.Errno)
}
