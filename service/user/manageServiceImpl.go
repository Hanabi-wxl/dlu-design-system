package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type ManageServiceImpl struct {
}

func (u ManageServiceImpl) SaveUser(s UserRequest) *errno.Errno {
	if s.IsStu == 1 {
		err := db.Student.Save(&model.Student{
			Number:    s.Number,
			Name:      s.Name,
			Password:  encodePassword(s.Number),
			Grade:     utils.GetGrade(),
			Phone:     s.Phone,
			Email:     s.Email,
			Gender:    s.Gender,
			CollegeID: s.CollegeID,
			MajorID:   s.MajorID,
			ClassID:   s.ClassID,
		})
		if err != nil {
			return errno.NewErrno(errno.DbErrorCode, err.Error())
		}
	} else {
		err := db.Teacher.Omit(db.Teacher.RoleID).Save(&model.Teacher{
			Number:    s.Number,
			Name:      s.Name,
			Password:  encodePassword(s.Number),
			Phone:     s.Phone,
			Email:     s.Email,
			CollegeID: s.CollegeID,
			MajorID:   s.MajorID,
			TitleID:   s.TitleID,
			DegreeID:  s.DegreeID,
			OfficeID:  s.OfficeID,
		})
		if err != nil {
			return errno.NewErrno(errno.DbErrorCode, err.Error())
		}
	}
	return nil
}

// encodePassword 加密
func encodePassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), consts.PassWordCost)
	logrus.Info(err)
	return string(bytes)
}
