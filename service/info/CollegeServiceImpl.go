package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CollegeServiceImpl struct {
}

func (c CollegeServiceImpl) GetCollegeList(size, num int) (*types.PageResp, *errno.Errno) {
	colleges, err := db.College.Limit(size).Offset(size * (num - 1)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	count, _ := db.College.Count()
	return &types.PageResp{ItemTotal: count, PageTotal: count / int64(size), Array: colleges}, nil
}

func (c CollegeServiceImpl) AddCollege(college *types.CollegeReq) *errno.Errno {
	err := db.College.Create(&model.College{ID: college.ID, Name: college.Name})
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c CollegeServiceImpl) DeleteCollege(id int64) *errno.Errno {
	deletes, err := db.College.Where(db.College.ID.Eq(id)).Delete()
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	if deletes.RowsAffected == 0 {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, gorm.ErrRecordNotFound.Error())
	}
	return nil
}

func (c CollegeServiceImpl) UpdateCollege(college *types.CollegeReq) *errno.Errno {
	updates, err := db.College.Updates(college)
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	if updates.RowsAffected == 0 {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, gorm.ErrRecordNotFound.Error())
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
