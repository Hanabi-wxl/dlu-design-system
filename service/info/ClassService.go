package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type ClassService interface {
	GetClassList(size, num int) ([]*model.Class, *errno.Errno)
	AddClass(class *model.Class) *errno.Errno
	DeleteClass(id int64) *errno.Errno
	UpdateClass(class *model.Class) *errno.Errno
	GetClass(id int64) (*model.Class, *errno.Errno)
}
