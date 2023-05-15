package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type TitleService interface {
	// GetTitles
	// @Description: 获取全部职称
	// @return []*model.TeacherTitle
	// @return *errno.Errno
	GetTitles() ([]*model.TeacherTitle, *errno.Errno)

	// GetTitle
	// @Description: 查询单个职称
	// @param id
	// @return *model.TeacherTitle
	// @return *errno.Errno
	GetTitle(id int64) (*model.TeacherTitle, *errno.Errno)
}
