package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type TitleServiceImpl struct {
}

func (c TitleServiceImpl) GetTitles() ([]*model.TeacherTitle, *errno.Errno) {
	teacherTitles, err := db.TeacherTitle.Where(db.TeacherTitle.ID.Gt(0)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return teacherTitles, nil
}

func (c TitleServiceImpl) GetTitle(id int64) (*model.TeacherTitle, *errno.Errno) {
	teacherTitle, err := db.TeacherTitle.Where(db.TeacherTitle.ID.Eq(id)).Take()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return teacherTitle, nil
}
