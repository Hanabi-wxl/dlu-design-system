package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
)

type MajorService interface {
	// GetMajorList
	// @Description: 分页查询专业
	// @param size
	// @param num
	// @return *types.PageResp
	// @return *errno.Errno
	GetMajorList(size, num int) (*types.PageResp, *errno.Errno)

	// AddMajor
	// @Description: 添加专业
	// @param major
	// @return *errno.Errno
	AddMajor(major *types.MajorReq) *errno.Errno

	// DeleteMajor
	// @Description: 删除专业
	// @param id
	// @return *errno.Errno
	DeleteMajor(id int64) *errno.Errno

	// UpdateMajor
	// @Description: 更新专业
	// @param major
	// @return *errno.Errno
	UpdateMajor(major *types.MajorReq) *errno.Errno

	// GetMajor
	// @Description: 查询单个专业
	// @param id
	// @return *model.Major
	// @return *errno.Errno
	GetMajor(id int64) (*model.Major, *errno.Errno)

	// GetMajorListByCollegeId
	// @Description: 根据学院查询专业
	// @param collegeId 学院id
	// @return []*model.Major
	// @return *errno.Errno
	GetMajorListByCollegeId(collegeId int64) ([]*model.Major, *errno.Errno)
}
