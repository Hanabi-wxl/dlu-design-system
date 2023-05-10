package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type CollegeService interface {
	GetCollegeList(size, num int) ([]*model.College, *errno.Errno)
	AddCollege(class *model.College) *errno.Errno
	DeleteCollege(id int64) *errno.Errno
	UpdateCollege(college *model.College) *errno.Errno
	GetCollege(id int64) (*model.College, *errno.Errno)
}
