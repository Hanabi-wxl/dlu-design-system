package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type CollegeServiceImpl struct {
}

func (c CollegeServiceImpl) GetCollegeList(size, num int) ([]*model.College, *errno.Errno) {
	colleges, err := db.College.Limit(size).Offset(size * (num - 1)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return colleges, nil
}

func (c CollegeServiceImpl) AddCollege() ([]*model.College, *errno.Errno) {

	return nil, nil
}

func (c CollegeServiceImpl) DeleteCollege() ([]*model.College, *errno.Errno) {

	return nil, nil
}

func (c CollegeServiceImpl) UpdateCollege() ([]*model.College, *errno.Errno) {

	return nil, nil
}

func (c CollegeServiceImpl) GetCollege() ([]*model.College, *errno.Errno) {

	return nil, nil
}
