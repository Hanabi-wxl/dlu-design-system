package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type SchoolServiceImpl struct {
}

func (c SchoolServiceImpl) GetSchoolList(size, num int) ([]*model.School, *errno.Errno) {
	schools, err := db.School.Limit(size).Offset(size * (num - 1)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return schools, nil
}

func (c SchoolServiceImpl) AddSchool() ([]*model.School, *errno.Errno) {

	return nil, nil
}

func (c SchoolServiceImpl) DeleteSchool() ([]*model.School, *errno.Errno) {

	return nil, nil
}

func (c SchoolServiceImpl) UpdateSchool() ([]*model.School, *errno.Errno) {

	return nil, nil
}

func (c SchoolServiceImpl) GetSchool() ([]*model.School, *errno.Errno) {

	return nil, nil
}
