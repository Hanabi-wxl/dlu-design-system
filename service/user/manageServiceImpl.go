package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type ManageServiceImpl struct {
}

func (ManageServiceImpl) SaveUser(s *types.AddUserRequest) *errno.Errno {
	if s.IsStu == 1 {
		err := db.Student.Save(&model.Student{
			Number:    s.Number,
			Name:      s.Name,
			Password:  encodePassword("Dlu@" + s.Number),
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
			Password:  encodePassword("Dlu@" + s.Number),
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

func (ManageServiceImpl) ResetPassword(id int64, isStu int8) *errno.Errno {
	if isStu == 1 {
		if ok, err := db.Student.Where(db.Student.ID.Eq(id)).Update(db.Student.Password, encodePassword("Dlu@123456")); err != nil {
			return errno.NewErrno(errno.DbErrorCode, err.Error())
		} else if ok.RowsAffected == 1 {
			return nil
		}
	} else if isStu == 2 {
		if ok, err := db.Teacher.Where(db.Teacher.ID.Eq(id)).Update(db.Teacher.Password, encodePassword("Dlu@123456")); err != nil {
			return errno.NewErrno(errno.DbErrorCode, err.Error())
		} else if ok.RowsAffected == 1 {
			return nil
		}
	}
	return errno.ResetPasswordErr
}

func (ManageServiceImpl) UpdateTeacherRole(id, roleId int64) *errno.Errno {
	ok, err := db.Teacher.Where(db.Teacher.ID.Eq(id)).Update(db.Teacher.RoleID, roleId)
	if err != nil {
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	if ok.RowsAffected == 1 {
		return nil
	}
	return errno.UpdateTeacherRoleErr
}

func (ManageServiceImpl) UpdateUser(s *types.UpdateUserRequest) *errno.Errno {
	if s.IsStu == 1 {
		ok, err := db.Student.Where(db.Student.ID.Eq(s.ID)).
			Select(db.Student.Phone, db.Student.Email, db.Student.CollegeID, db.Student.MajorID, db.Student.ClassID).
			Updates(&model.Student{
				Phone:     s.Phone,
				Email:     s.Email,
				CollegeID: s.CollegeID,
				MajorID:   s.MajorID,
				ClassID:   s.ClassID,
			})
		if err != nil {
			return errno.NewErrno(errno.DbErrorCode, err.Error())
		}
		if ok.RowsAffected == 1 {
			return nil
		}
	} else {
		ok, err := db.Teacher.Where(db.Teacher.ID.Eq(s.ID)).
			Select(db.Teacher.Phone, db.Teacher.Email, db.Teacher.CollegeID, db.Teacher.MajorID, db.Teacher.TitleID, db.Teacher.OfficeID, db.Teacher.DegreeID).
			Updates(&model.Teacher{
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
		if ok.RowsAffected == 1 {
			return nil
		}
	}
	return errno.UpdateUserErr
}

func (ManageServiceImpl) DeleteUser(id int64, isStu int8) *errno.Errno {
	if isStu == 1 {
		if ok, err := db.Student.Where(db.Student.ID.Eq(id)).Delete(); err != nil {
			return errno.NewErrno(errno.DbErrorCode, err.Error())
		} else if ok.RowsAffected == 1 {
			return nil
		}
	} else if isStu == 2 {
		if ok, err := db.Teacher.Where(db.Teacher.ID.Eq(id)).Delete(); err != nil {
			return errno.NewErrno(errno.DbErrorCode, err.Error())
		} else if ok.RowsAffected == 1 {
			return nil
		}
	}
	return errno.DeleteUserErr
}

// encodePassword 加密
func encodePassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), consts.PassWordCost)
	logrus.Info(err)
	return string(bytes)
}
