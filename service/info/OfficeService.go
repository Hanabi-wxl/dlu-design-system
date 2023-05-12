package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type SectionService interface {
	GetOffices() ([]*model.TeacherOffice, *errno.Errno)
	GetOffice(id int64) (*model.TeacherOffice, *errno.Errno)
}
