package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type SchoolService interface {
	GetSchoolList(size, num int) ([]*model.School, *errno.Errno)
	AddSchool() ([]*model.School, *errno.Errno)
	DeleteSchool() ([]*model.School, *errno.Errno)
	UpdateSchool() ([]*model.School, *errno.Errno)
	GetSchool() ([]*model.School, *errno.Errno)
}
