package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type DegreeServiceImpl struct {
}

func (c DegreeServiceImpl) GetDegrees() ([]*model.TeacherDegree, *errno.Errno) {
	teacherDegrees, err := db.TeacherDegree.Where(db.TeacherDegree.ID.Gt(0)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return teacherDegrees, nil
}

func (c DegreeServiceImpl) GetDegree(id int64) (*model.TeacherDegree, *errno.Errno) {
	teacherDegree, err := db.TeacherDegree.Where(db.TeacherDegree.ID.Eq(id)).Take()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return teacherDegree, nil
}
