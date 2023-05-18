package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type SectionService interface {
	// GetOffices
	// @Description: 获取全部科室
	// @return []*model.TeacherOffice
	// @return *errno.Errno
	GetOffices() ([]*model.TeacherOffice, *errno.Errno)

	// GetOffice
	// @Description: 查询单个科室
	// @param id
	// @return *model.TeacherOffice
	// @return *errno.Errno
	GetOffice(id int64) (*model.TeacherOffice, *errno.Errno)
}
