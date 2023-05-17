package Service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type ApproveServiceImpl struct {
}

func (c ApproveServiceImpl) MajorApproveSubject(subjectIds []int64) *errno.Errno {
	for _, v := range subjectIds {
		_, err := db.Subject.Where(db.Subject.MajorID.Eq(v)).Update(db.Subject.ProgressID, "2")
		if err != nil {
			logrus.Error(err)
			return errno.NewErrno(errno.DbErrorCode, err.Error())
		}
	}
	return nil
}

func (c ApproveServiceImpl) CollegeApproveSubject(subjectIds []int64) *errno.Errno {
	for _, v := range subjectIds {
		_, err := db.Subject.Where(db.Subject.CollegeID.Eq(v)).Update(db.Subject.ProgressID, "3")
		if err != nil {
			logrus.Error(err)
			return errno.NewErrno(errno.DbErrorCode, err.Error())
		}
	}
	return nil
}
