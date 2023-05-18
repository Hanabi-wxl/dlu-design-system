package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	"github.com/sirupsen/logrus"
)

type MajorServiceImpl struct {
}

func (c MajorServiceImpl) GetMajorList(size, num int) (*types.PageResp, *errno.Errno) {
	majors, err := db.Major.Limit(size).Offset(size * (num - 1)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	count, _ := db.Major.Count()
	return &types.PageResp{ItemTotal: count, PageTotal: count/int64(size) + 1, Array: majors}, nil
}

func (c MajorServiceImpl) AddMajor(major *types.MajorReq) *errno.Errno {
	err := db.Major.Create(&model.Major{ID: major.ID, Name: major.Name})
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

func (c MajorServiceImpl) UpdateMajor(major *types.MajorReq) *errno.Errno {
	err := db.Major.Save(&model.Major{Name: major.Name})
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

func (c MajorServiceImpl) GetMajorListByCollegeId(collegeId int64) ([]*model.Major, *errno.Errno) {
	majors, err := db.Major.Where(db.Major.CollegeID.Eq(collegeId)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return majors, nil
}
