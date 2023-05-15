package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
)

type CollegeService interface {
	// GetCollegeList
	// @Description: 分页查询学院
	// @param size
	// @param num
	// @return *types.PageResp
	// @return *errno.Errno
	GetCollegeList(size, num int) (*types.PageResp, *errno.Errno)

	// AddCollege
	// @Description: 添加学院
	// @param class
	// @return *errno.Errno
	AddCollege(class *types.CollegeReq) *errno.Errno

	// DeleteCollege
	// @Description: 删除学院
	// @param id
	// @return *errno.Errno
	DeleteCollege(id int64) *errno.Errno

	// UpdateCollege
	// @Description: 更新学院
	// @param college
	// @return *errno.Errno
	UpdateCollege(college *types.CollegeReq) *errno.Errno

	// GetCollege
	// @Description: 根据id查询学院
	// @param id
	// @return *model.College
	// @return *errno.Errno
	GetCollege(id int64) (*model.College, *errno.Errno)
}
