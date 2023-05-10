package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type SchoolService interface {
	GetSchoolList(size, num int) ([]*model.School, *errno.Errno)
	AddSchool(school *model.School) *errno.Errno
	DeleteSchool(id int64) *errno.Errno
	UpdateSchool(school *model.School) *errno.Errno
	GetSchool(id int64) (*model.School, *errno.Errno)
}
