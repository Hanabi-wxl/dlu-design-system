package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type MajorService interface {
	GetMajorList(size, num int) ([]*model.Major, *errno.Errno)
	AddMajor() ([]*model.Major, *errno.Errno)
	DeleteMajor() ([]*model.Major, *errno.Errno)
	UpdateMajor() ([]*model.Major, *errno.Errno)
	GetMajor() ([]*model.Major, *errno.Errno)
}
