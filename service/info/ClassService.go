package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type ClassService interface {
	GetClassList(size, num int) ([]*model.Class, *errno.Errno)
	AddClass() ([]*model.Class, *errno.Errno)
	DeleteClass() ([]*model.Class, *errno.Errno)
	UpdateClass() ([]*model.Class, *errno.Errno)
	GetClass() ([]*model.Class, *errno.Errno)
}
