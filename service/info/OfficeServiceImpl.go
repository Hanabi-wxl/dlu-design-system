package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type SectionServiceImpl struct {
}

func (c SectionServiceImpl) GetOffices() ([]*model.TeacherOffice, *errno.Errno) {
	teacherOffices, err := db.TeacherOffice.Where(db.TeacherOffice.ID.Gt(0)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return teacherOffices, nil
}

func (c SectionServiceImpl) GetOffice(id int64) (*model.TeacherOffice, *errno.Errno) {
	teacherOffice, err := db.TeacherOffice.Where(db.TeacherOffice.ID.Eq(id)).Take()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return teacherOffice, nil
}
