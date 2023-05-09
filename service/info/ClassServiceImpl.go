package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type ClassServiceImpl struct {
}

func (c ClassServiceImpl) GetClassList(size, num int) ([]*model.Class, *errno.Errno) {
	classes, err := db.Class.Limit(size).Offset(size * (num - 1)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return classes, nil
}

func (c ClassServiceImpl) AddClass() ([]*model.Class, *errno.Errno) {

	return nil, nil
}

func (c ClassServiceImpl) DeleteClass() ([]*model.Class, *errno.Errno) {

	return nil, nil
}

func (c ClassServiceImpl) UpdateClass() ([]*model.Class, *errno.Errno) {

	return nil, nil
}

func (c ClassServiceImpl) GetClass() ([]*model.Class, *errno.Errno) {

	return nil, nil
}
