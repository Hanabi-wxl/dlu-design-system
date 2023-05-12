package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogServiceImpl struct {
}

func (LogServiceImpl) AddLog(content string, c *gin.Context, state int8, method int8) {
	ip := c.ClientIP()
	number, _ := c.Get("number")
	logContent := &model.Log{
		Content:    content,
		State:      state,
		Method:     method,
		Operator:   number.(string),
		OperatorIP: ip,
	}
	if state == 1 {
		logrus.Info(logContent)
	} else {
		logrus.Error(logContent)
	}
	_ = db.Log.Save(logContent)
}
