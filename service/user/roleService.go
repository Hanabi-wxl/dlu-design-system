package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type RoleService interface {
	// GetRoleList
	// @Description: 获取角色列表
	// @return *[]*model.Role 角色列表
	// @return *errno.Errno
	GetRoleList() (*[]*model.Role, *errno.Errno)

	// UpdateRole
	// @Description: 更新角色信息
	// @param role 角色
	// @return *errno.Errno
	UpdateRole(role *model.Role) *errno.Errno

	// DeleteRole
	// @Description: 删除角色
	// @param id 角色id
	// @return *errno.Errno
	DeleteRole(id int64) *errno.Errno

	// AddRole
	// @Description: 添加角色
	// @param role 角色
	// @return *errno.Errno
	AddRole(role *model.Role) *errno.Errno
}
