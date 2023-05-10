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

func (c SchoolServiceImpl) AddSchool(school *model.School) *errno.Errno {
	err := db.School.Create(school)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c SchoolServiceImpl) DeleteSchool(id int64) *errno.Errno {
	_, err := db.School.Where(db.School.ID.Eq(id)).Delete()
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c SchoolServiceImpl) UpdateSchool(school *model.School) *errno.Errno {
	err := db.School.Save(school)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c SchoolServiceImpl) GetSchool(id int64) (*model.School, *errno.Errno) {
	school, err := db.School.Where(db.School.ID.Eq(id)).Take()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return school, nil
}
