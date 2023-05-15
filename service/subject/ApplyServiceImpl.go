package Service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type ApplyServiceImpl struct {
}

func (c ApplyServiceImpl) AddSubject(subject *model.Subject) *errno.Errno {
	err := db.Subject.Create(subject)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}
func (c ApplyServiceImpl) CountSubject(teacherId int64) (int64, *errno.Errno) {
	count, err := db.Subject.Where(db.Subject.FirstTeacherID.Eq(teacherId)).Count()
	if err != nil {
		logrus.Error(err)
		return 0, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return count, nil
}
