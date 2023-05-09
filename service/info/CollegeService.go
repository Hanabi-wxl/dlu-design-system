package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type CollegeService interface {
	GetCollegeList(size, num int) ([]*model.College, *errno.Errno)
	AddCollege() ([]*model.College, *errno.Errno)
	DeleteCollege() ([]*model.College, *errno.Errno)
	UpdateCollege() ([]*model.College, *errno.Errno)
	GetCollege() ([]*model.College, *errno.Errno)
}
