package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/sirupsen/logrus"
)

type SchoolServiceImpl struct {
}

func (s SchoolServiceImpl) GetAllSchool() ([]*model.School, error) {
	schools, err := db.School.Find()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return schools, nil
}
