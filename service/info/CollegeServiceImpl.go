package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/sirupsen/logrus"
)

type CollegeServiceImpl struct {
}

func (c CollegeServiceImpl) GetCollegeList(size, num int) ([]*model.College, error) {
	colleges, err := db.College.Limit(size).Offset(size * (num - 1)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return colleges, nil
}
