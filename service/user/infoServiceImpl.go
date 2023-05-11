package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
)

type InfoServiceImpl struct {
}

//在实现类中写一个user的struct

type User struct {
	ID        int64   `json:"id"`
	Number    string  `json:"number"`
	Name      string  `json:"name"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
	CollegeID int64   `json:"college_id"`
	MajorID   int64   `json:"major_id"`
	//ClassID   int64   `json:"class_id"`
	IsStu int8 `json:"is_stu"`
}

// GetUserById
//
//	@Description:
//	@receiver u
//	@param id
//	@param isStu 0表示老师 1表示学生
//	@return User
//	@return *errno.Errno
func (InfoServiceImpl) GetUserById(id int64, isStu int8) (*User, *errno.Errno) {
	var u User
	if isStu == 1 {
		stu, err := db.Student.Where(db.Student.ID.Eq(id)).Take()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}
		u = User{
			ID:        stu.ID,
			Number:    stu.Number,
			Name:      stu.Name,
			Phone:     stu.Phone,
			Email:     stu.Email,
			CollegeID: stu.CollegeID,
			MajorID:   stu.MajorID,
			IsStu:     1,
		}
	} else {
		teacher, err := db.Teacher.Where(db.Teacher.ID.Eq(id)).Take()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}
		u = User{
			ID:        teacher.ID,
			Number:    teacher.Number,
			Name:      teacher.Name,
			Phone:     teacher.Phone,
			Email:     teacher.Email,
			CollegeID: teacher.CollegeID,
			MajorID:   teacher.MajorID,
			IsStu:     0,
		}
	}
	return &u, nil
}

func (InfoServiceImpl) GetUserByNumber(number string, isStu int8) (*User, *errno.Errno) {
	var u User
	if isStu == 1 {
		stu, err := db.Student.Where(db.Student.Number.Eq(number)).Take()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}
		u = User{
			ID:        stu.ID,
			Number:    stu.Number,
			Name:      stu.Name,
			Phone:     stu.Phone,
			Email:     stu.Email,
			CollegeID: stu.CollegeID,
			MajorID:   stu.MajorID,
			IsStu:     1,
		}
	} else {
		teacher, err := db.Teacher.Where(db.Teacher.Number.Eq(number)).Take()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}
		u = User{
			ID:        teacher.ID,
			Number:    teacher.Number,
			Name:      teacher.Name,
			Phone:     teacher.Phone,
			Email:     teacher.Email,
			CollegeID: teacher.CollegeID,
			MajorID:   teacher.MajorID,
			IsStu:     0,
		}
	}
	return &u, nil
}

func (InfoServiceImpl) GetUsersByMajor(majorId int64, isStu int8, size, num int) (*[]User, *errno.Errno) {
	var Us []User
	if isStu == 1 {
		stus, err := db.Student.Where(db.Student.MajorID.Eq(int64(majorId))).Limit(size).Offset(size * (num - 1)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}
		for _, val := range stus {
			u := User{
				ID:        val.ID,
				Number:    val.Number,
				Name:      val.Name,
				Phone:     val.Phone,
				Email:     val.Email,
				CollegeID: val.CollegeID,
				MajorID:   val.MajorID,
				IsStu:     1,
			}
			Us = append(Us, u)
		}
	} else {
		teachers, err := db.Teacher.Where(db.Teacher.MajorID.Eq(int64(majorId))).Limit(size).Offset(size * (num - 1)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}
		for _, val := range teachers {
			u := User{
				ID:        val.ID,
				Number:    val.Number,
				Name:      val.Name,
				Phone:     val.Phone,
				Email:     val.Email,
				CollegeID: val.CollegeID,
				MajorID:   val.MajorID,
				IsStu:     1,
			}
			Us = append(Us, u)
		}
	}
	return &Us, nil
}

func (InfoServiceImpl) GetUsersByCollege(collegeId int64, isStu int8, size, num int) (*[]User, *errno.Errno) {
	var Us []User
	if isStu == 1 {
		stus, err := db.Student.Where(db.Student.CollegeID.Eq(int64(collegeId))).Limit(size).Offset(size * (num - 1)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}
		for _, val := range stus {
			u := User{
				ID:        val.ID,
				Number:    val.Number,
				Name:      val.Name,
				Phone:     val.Phone,
				Email:     val.Email,
				CollegeID: val.CollegeID,
				MajorID:   val.MajorID,
				IsStu:     1,
			}
			Us = append(Us, u)
		}
	} else {
		teachers, err := db.Teacher.Where(db.Teacher.CollegeID.Eq(int64(collegeId))).Limit(size).Offset(size * (num - 1)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}
		for _, val := range teachers {
			u := User{
				ID:        val.ID,
				Number:    val.Number,
				Name:      val.Name,
				Phone:     val.Phone,
				Email:     val.Email,
				CollegeID: val.CollegeID,
				MajorID:   val.MajorID,
				IsStu:     1,
			}
			Us = append(Us, u)
		}
	}
	return &Us, nil
}
