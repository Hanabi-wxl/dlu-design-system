package Service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type ApplyService interface {
	AddSubject(subject *model.Subject, majorIds []int64) *errno.Errno
	CountSubject(teacherId int64) (int64, *errno.Errno)
	CountSubjectByStudent(studentId int64) (int64, *errno.Errno)
	UpdateSubject(subject *model.Subject, majorIds []int64) *errno.Errno
	GetSelfSubjectsByYear(year int64, userInfo *jwt.UserInfo) ([]*model.Subject, *errno.Errno)
	GetTypes() ([]*model.SubjectType, *errno.Errno)
	GetOrigins() ([]*model.SubjectOrigin, *errno.Errno)
	Delete(subjectId int64) *errno.Errno
	GetSubjectById(subjectId int64) (*model.Subject, *errno.Errno)
	GetSubjectByIds(subjectIds []int64) ([]*model.Subject, *errno.Errno)
}
