package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type SectionService interface {
	GetSections() ([]*model.TeacherOffice, *errno.Errno)
	GetSection() ([]*model.TeacherOffice, *errno.Errno)
}
