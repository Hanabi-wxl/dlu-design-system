package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
)

type ManageService interface {
	// SaveUser
	// @Description: 添加用户 学生或老师
	// @param userReq 内容
	// @return *errno.Errno 错误
	SaveUser(userReq *types.AddUserRequest) *errno.Errno

	// ResetPassword
	// @Description: 重置用户密码
	// @param id 用户id
	// @param isStu 1学生 2教师
	// @return *errno.Errno
	ResetPassword(id int64, isStu int8) *errno.Errno

	// UpdateTeacherRole
	// @Description: 修改教师角色
	// @param id 教师id
	// @return *errno.Errno
	UpdateTeacherRole(id, roleId int64) *errno.Errno

	// UpdateUser
	// @Description: 更新用户
	// @param id 用户id
	// @param isStu 1学生 2教师
	// @return *errno.Errno
	UpdateUser(userReq *types.UpdateUserRequest) *errno.Errno

	// DeleteUser
	// @Description: 删除用户
	// @param id 用户id
	// @param isStu 1学生 2教师
	// @return *errno.Errno
	DeleteUser(id int64, isStu int8) *errno.Errno

	// ManagerList
	// @Description: 获取管理员名单
	// @param roleId 角色Id
	// @return *[]*model.Teacher 管理员列表
	// @return *errno.Errno
	ManagerList(roleId int64) (*[]*model.Teacher, *errno.Errno)

	// DeleteManager
	// @Description: 删除管理员权限
	// @param id 用户id
	// @return *errno.Errno
	DeleteManager(id int64) *errno.Errno
}
