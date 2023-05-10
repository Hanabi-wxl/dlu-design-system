package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type MajorServiceImpl struct {
}

func (c MajorServiceImpl) GetMajorList(size, num int) ([]*model.Major, *errno.Errno) {
	majors, err := db.Major.Limit(size).Offset(size * (num - 1)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return majors, nil
}

func (c MajorServiceImpl) AddMajor(major *model.Major) *errno.Errno {
	err := db.Major.Create(major)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c MajorServiceImpl) DeleteMajor(id int64) *errno.Errno {
	_, err := db.Major.Where(db.Major.ID.Eq(id)).Delete()
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c MajorServiceImpl) UpdateMajor(major *model.Major) *errno.Errno {
	err := db.Major.Save(major)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c MajorServiceImpl) GetMajor(id int64) (*model.Major, *errno.Errno) {
	major, err := db.Major.Where(db.Major.ID.Eq(id)).Take()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return major, nil
}
