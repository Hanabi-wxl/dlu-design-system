package Service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type ApplyService interface {
	AddSubject(subject *model.Subject) *errno.Errno
	CountSubject(teacherId int64) (int64, *errno.Errno)
}
