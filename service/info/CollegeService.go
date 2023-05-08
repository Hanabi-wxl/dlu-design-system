package service

import "github.com/Hanabi-wxl/dlu-design-system/dal/model"

type CollegeService interface {
	GetCollegeList(size, num int) ([]*model.College, error)
}
