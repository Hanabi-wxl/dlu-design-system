package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type SchoolService interface {
	// GetSchoolList
	// @Description: 查询全部学校
	// @param size
	// @param num
	// @return []*model.School
	// @return *errno.Errno
	GetSchoolList(size, num int) ([]*model.School, *errno.Errno)

	// AddSchool
	// @Description: 添加学校
	// @param school
	// @return *errno.Errno
	AddSchool(school *model.School) *errno.Errno

	// DeleteSchool
	// @Description: 删除学校
	// @param id
	// @return *errno.Errno
	DeleteSchool(id int64) *errno.Errno

	// UpdateSchool
	// @Description: 更新学校
	// @param school
	// @return *errno.Errno
	UpdateSchool(school *model.School) *errno.Errno

	// GetSchool
	// @Description: 查询学校
	// @param id
	// @return *model.School
	// @return *errno.Errno
	GetSchool(id int64) (*model.School, *errno.Errno)
}
