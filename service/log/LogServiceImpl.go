package service

import (
	"context"
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
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
