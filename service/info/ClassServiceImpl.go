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

func (c ClassServiceImpl) AddClass(class *model.Class) *errno.Errno {
	err := db.Class.Create(class)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c ClassServiceImpl) DeleteClass(id int64) *errno.Errno {
	_, err := db.Class.Where(db.Class.ID.Eq(id)).Delete()
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c ClassServiceImpl) UpdateClass(class *model.Class) *errno.Errno {
	err := db.Class.Save(class)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c ClassServiceImpl) GetClass(id int64) (*model.Class, *errno.Errno) {
	class, err := db.Class.Where(db.Class.ID.Eq(id)).Take()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return class, nil
}
