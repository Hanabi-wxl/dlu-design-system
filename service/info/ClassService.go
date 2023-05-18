package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
)

type ClassService interface {
	// GetClassList
	// @Description: 分页查询班级
	// @param size 分页大小
	// @param num 分页页数
	// @return *[]*model.Class
	// @return *errno.Errno
	GetClassList(req *types.ClassQueryReq) (*types.PageResp, *errno.Errno)

	// AddClass
	// @Description: 添加班级
	// @param class
	// @return *errno.Errno
	AddClass(class *types.ClassReq) *errno.Errno

	// DeleteClass
	// @Description: 删除班级
	// @param id 班级id
	// @return *errno.Errno
	DeleteClass(id int64) *errno.Errno

	// UpdateClass
	// @Description: 更新班级 不可更新id
	// @param class
	// @return *errno.Errno
	UpdateClass(class *types.ClassReq) *errno.Errno

	// GetClass
	// @Description: 根据id获取班级
	// @param id 班级id
	// @return *model.Class
	// @return *errno.Errno
	GetClass(id int64) (*model.Class, *errno.Errno)
}
