package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
)

type RoleServiceImpl struct {
}

func (RoleServiceImpl) GetRoleList() (*[]*model.Role, *errno.Errno) {
	if roles, err := db.Role.Find(); err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	} else {
		return &roles, nil
	}
}

func (RoleServiceImpl) UpdateRole(role *types.RoleReq) *errno.Errno {
	if ok, err := db.Role.Where(db.Role.ID.Eq(role.ID)).Update(db.Role.Name, role.Name); err != nil {
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	} else if ok.RowsAffected != 1 {
		return errno.UpdateRoleErr
	} else {
		return nil
	}
}

func (RoleServiceImpl) DeleteRole(id int64) *errno.Errno {
	if ok, err := db.Role.Where(db.Role.ID.Eq(id)).Delete(); err != nil {
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	} else if ok.RowsAffected != 1 {
		return errno.DeleteRoleErr
	} else {
		return nil
	}
}

func (RoleServiceImpl) AddRole(role *types.RoleReq) *errno.Errno {
	if err := db.Role.Select(db.Role.Name).Save(&model.Role{ID: role.ID, Name: role.Name}); err != nil {
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	} else {
		return nil
	}
}

func (RoleServiceImpl) GetRole(id int64) (*model.Role, *errno.Errno) {
	if role, err := db.Role.Where(db.Role.ID.Eq(id)).Take(); err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	} else {
		return role, nil
	}
}
