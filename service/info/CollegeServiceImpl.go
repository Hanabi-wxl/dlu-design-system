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

func (c CollegeServiceImpl) AddCollege(college *model.College) *errno.Errno {
	err := db.College.Create(college)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c CollegeServiceImpl) DeleteCollege(id int64) *errno.Errno {
	_, err := db.College.Where(db.College.ID.Eq(id)).Delete()
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c CollegeServiceImpl) UpdateCollege(college *model.College) *errno.Errno {
	err := db.College.Save(college)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c CollegeServiceImpl) GetCollege(id int64) (*model.College, *errno.Errno) {
	college, err := db.College.Where(db.College.ID.Eq(id)).Take()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return college, nil
}
