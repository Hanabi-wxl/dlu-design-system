package Service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
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

func (c ApplyServiceImpl) UpdateSubject(subject *model.Subject) *errno.Errno {
	err := db.Subject.Where(db.Subject.ID.Eq(subject.ID)).Save(subject)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c ApplyServiceImpl) GetSelfSubjectsByYear(year int64, userInfo *jwt.UserInfo) ([]*model.Subject, *errno.Errno) {
	if userInfo.RoleId == 2 {
		subjects, err := db.Subject.Where(db.Subject.Year.Eq(year)).Where(db.Subject.FirstTeacherID.Eq(userInfo.ID), (db.Subject.SecondTeacherID.Eq(userInfo.ID))).Find()
		if err != nil {
			logrus.Error(err)
			return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
		}
		return subjects, nil
	} else if userInfo.RoleId == 1 {
		subjects, err := db.Subject.Where(db.Subject.Year.Eq(year)).Where(db.Subject.StudentID.Eq(userInfo.ID)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
		}
		return subjects, nil
	}
	return nil, nil
}
