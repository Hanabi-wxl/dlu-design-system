package service

import "github.com/Hanabi-wxl/dlu-design-system/dal/model"

type SchoolService interface {
	GetAllSchool() ([]*model.School, error)
}
