package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type SectionServiceImpl struct {
}

func (c SectionServiceImpl) GetSections() ([]*model.TeacherOffice, *errno.Errno) {
	teacherSections, err := db.TeacherOffice.Where(db.TeacherOffice.ID.Gt(0)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return teacherSections, nil
}

func (c SectionServiceImpl) GetSection(id int64) (*model.TeacherOffice, *errno.Errno) {
	TeacherSection, err := db.TeacherOffice.Where(db.TeacherOffice.ID.Eq(id)).Take()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return TeacherSection, nil
}
