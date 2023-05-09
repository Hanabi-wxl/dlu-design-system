package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type DegreeService interface {
	GetDegrees() ([]*model.TeacherDegree, *errno.Errno)
	GetDegree() ([]*model.TeacherDegree, *errno.Errno)
}
