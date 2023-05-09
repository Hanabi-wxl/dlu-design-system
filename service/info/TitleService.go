package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type TitleService interface {
	GetTitles() ([]*model.TeacherTitle, *errno.Errno)
	GetTitle() ([]*model.TeacherTitle, *errno.Errno)
}
