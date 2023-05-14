package service

import (
	"context"
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	"github.com/sirupsen/logrus"
)

type LogServiceImpl struct {
}

func (LogServiceImpl) AddLog(content string, c context.Context, ip string, state int8, method int8) {
	userInfo := jwt.GetUserInfo(c)
	logContent := &model.Log{
		Content:    content,
		State:      state,
		Method:     method,
		Operator:   userInfo.Number,
		OperatorIP: ip,
	}
	if state == 1 {
		logrus.Info(logContent)
	} else {
		logrus.Error(logContent)
	}
	_ = db.Log.Save(logContent)
}

func (LogServiceImpl) LogList(size int, num int) (*types.PageResp, *errno.Errno) {
	if logs, err := db.Log.Limit(size).Offset((size - 1) * num).Find(); err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	} else {
		count, _ := db.Log.Count()
		pages := count/int64(size) + 1
		return &types.PageResp{ItemTotal: count, PageTotal: pages, Array: &logs}, nil
	}
}

func (LogServiceImpl) LogDateList(req *types.LogDateReq) (*types.PageResp, *errno.Errno) {
	if logs, err := db.Log.Where(db.Log.CreatedAt.Gte(req.Start), db.Log.CreatedAt.Lte(req.End)).
		Limit(req.Num).Offset((req.Size - 1) * req.Num).Find(); err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	} else {
		count, _ := db.Log.Count()
		pages := count/int64(req.Size) + 1
		return &types.PageResp{ItemTotal: count, PageTotal: pages, Array: &logs}, nil
	}
}

func (LogServiceImpl) LogStateList(req *types.LogStateReq) (*types.PageResp, *errno.Errno) {
	if logs, err := db.Log.Where(db.Log.State.Eq(req.StateId)).Limit(req.Size).Offset((req.Size - 1) * req.Num).Find(); err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	} else {
		count, _ := db.Log.Count()
		pages := count/int64(req.Size) + 1
		return &types.PageResp{ItemTotal: count, PageTotal: pages, Array: &logs}, nil
	}
}

func (LogServiceImpl) LogMethodList(req *types.LogMethodReq) (*types.PageResp, *errno.Errno) {
	if logs, err := db.Log.Where(db.Log.Method.Eq(req.MethodId)).Limit(req.Size).Offset((req.Size - 1) * req.Num).Find(); err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	} else {
		count, _ := db.Log.Count()
		pages := count/int64(req.Size) + 1
		return &types.PageResp{ItemTotal: count, PageTotal: pages, Array: &logs}, nil
	}
}
