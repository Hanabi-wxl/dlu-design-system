package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/utils"
	"github.com/sirupsen/logrus"
)

type ClassServiceImpl struct {
}

func (c ClassServiceImpl) GetClassList(req *types.ClassQueryReq) (*types.PageResp, *errno.Errno) {
	classes, err := db.Class.Where(db.Class.MajorID.Eq(req.MajorId), db.Class.Grade.Eq(req.Year)).
		Limit(req.Size).Offset(req.Size * (req.Num - 1)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	count, _ := db.Class.Where(db.Class.MajorID.Eq(req.MajorId), db.Class.Grade.Eq(req.Year)).Count()
	return &types.PageResp{ItemTotal: count, PageTotal: count/int64(req.Size) + 1, Array: classes}, nil
}

func (c ClassServiceImpl) AddClass(class *types.ClassReq) *errno.Errno {
	err := db.Class.Create(&model.Class{ID: class.ID, Name: class.Name, MajorID: class.MajorId, Grade: utils.GetGradeByYear(class.Grade)})
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

func (c ClassServiceImpl) UpdateClass(class *types.ClassReq) *errno.Errno {
	err := db.Class.Where(db.Class.ID.Eq(class.ID)).Save(&model.Class{Name: class.Name, MajorID: class.MajorId, Grade: utils.GetGradeByYear(class.Grade)})
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
