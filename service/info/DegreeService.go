package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type DegreeService interface {
	// GetDegrees
	// @Description: 获取全部学位
	// @return []*model.TeacherDegree
	// @return *errno.Errno
	GetDegrees() ([]*model.TeacherDegree, *errno.Errno)

	// GetDegree
	// @Description: 查询单个学位
	// @param id
	// @return *model.TeacherDegree
	// @return *errno.Errno
	GetDegree(id int64) (*model.TeacherDegree, *errno.Errno)
}
