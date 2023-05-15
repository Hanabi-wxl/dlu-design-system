package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	"github.com/sirupsen/logrus"
	"sync"
)

type InfoServiceImpl struct {
}

//在实现类中写一个user的struct

type User struct {
	ID          int64   `json:"id"`
	Number      string  `json:"number"`
	Name        string  `json:"name"`
	Phone       *string `json:"phone"`
	Email       *string `json:"email"`
	CollegeID   int64   `json:"college_id"`
	CollegeName string  `json:"college_name"`
	MajorID     int64   `json:"major_id"`
	MajorName   string  `json:"major_name"`
	ClassID     int64   `json:"class_id"`
	RoleId      int     `json:"role_id"`
	IsStu       int8    `json:"is_stu"`
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
			RoleId:    1,
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
			RoleId:    int(teacher.RoleID),
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
			RoleId:    1,
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
			RoleId:    int(teacher.RoleID),
		}
	}
	return &u, nil
}

func (InfoServiceImpl) GetUserByNumberMajor(req types.NumberMajorRequest) (*types.UserResp, *errno.Errno) {
	var user types.UserResp
	if req.IsStu == 1 {
		err := db.Student.Select(db.Student.ALL, db.Major.Name.As("MajorName"), db.College.Name.As("CollegeName"),
			db.Class.Name.As("ClassName")).Omit(db.Student.Password).
			Where(db.Student.Number.Eq(req.Number)).
			LeftJoin(db.Major, db.Major.ID.EqCol(db.Student.MajorID)).
			LeftJoin(db.College, db.College.ID.EqCol(db.Student.CollegeID)).
			LeftJoin(db.Class, db.Class.ID.EqCol(db.Student.ClassID)).
			Scan(&user)
		user.IsStu = 1
		if err != nil {
			return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
		}
		user.IsStu = 2
	} else {
		teacher := db.Teacher

		err := teacher.Select(teacher.ALL, db.Major.Name.As("MajorName"), db.College.Name.As("CollegeName"),
			db.TeacherTitle.Name.As("TitleName"), db.TeacherDegree.Name.As("DegreeName"), db.TeacherOffice.Name.As("OfficeName")).
			Omit(db.Teacher.Password).
			Where(teacher.Number.Eq(req.Number)).
			LeftJoin(db.Major, db.Major.ID.EqCol(teacher.MajorID)).
			LeftJoin(db.College, db.College.ID.EqCol(teacher.CollegeID)).
			LeftJoin(db.TeacherTitle, db.TeacherTitle.ID.EqCol(teacher.TitleID)).
			LeftJoin(db.TeacherDegree, db.TeacherDegree.ID.EqCol(teacher.DegreeID)).
			LeftJoin(db.TeacherOffice, db.TeacherOffice.ID.EqCol(teacher.OfficeID)).
			Scan(&user)
		user.IsStu = 2
		if err != nil {
			return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
		}
	}
	return &user, nil
}

func (InfoServiceImpl) GetUsersByMajor(majorId int64, isStu int8, size, num int) (*[]User, *errno.Errno) {
	var Us []User
	if isStu == 1 {
		stus, err := db.Student.Where(db.Student.MajorID.Eq(majorId)).Limit(size).Offset(size * (num - 1)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(len(stus))
		for _, val := range stus {
			go func(val *model.Student) {
				defer wg.Done()
				major, err1 := db.Major.Where(db.Major.ID.Eq(val.MajorID)).First()
				if err1 != nil {
					logrus.Error(err1)
				}
				u := User{
					ID:        val.ID,
					Number:    val.Number,
					Name:      val.Name,
					Phone:     val.Phone,
					Email:     val.Email,
					CollegeID: val.CollegeID,
					MajorID:   val.MajorID,
					IsStu:     1,
					MajorName: major.Name,
					RoleId:    1,
				}
				Us = append(Us, u)
			}(val)
		}
		wg.Wait()
	} else {
		teachers, err := db.Teacher.Where(db.Teacher.MajorID.Eq(majorId)).Limit(size).Offset(size * (num - 1)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}

		var wg sync.WaitGroup
		wg.Add(len(teachers))
		for _, val := range teachers {
			go func(val *model.Teacher) {
				defer wg.Done()
				major, err1 := db.Major.Where(db.Major.ID.Eq(val.MajorID)).First()
				if err1 != nil {
					logrus.Error(err1)
				}
				u := User{
					ID:        val.ID,
					Number:    val.Number,
					Name:      val.Name,
					Phone:     val.Phone,
					Email:     val.Email,
					CollegeID: val.CollegeID,
					MajorID:   val.MajorID,
					IsStu:     0,
					MajorName: major.Name,
					RoleId:    int(val.RoleID),
				}
				Us = append(Us, u)
			}(val)
		}
		wg.Wait()
	}
	return &Us, nil
}

func (InfoServiceImpl) GetUsersByCollege(collegeId int64, isStu int8, size, num int) (*[]User, *errno.Errno) {
	var Us []User
	if isStu == 1 {
		stus, err := db.Student.Where(db.Student.CollegeID.Eq(collegeId)).Limit(size).Offset(size * (num - 1)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}
		var wg sync.WaitGroup
		wg.Add(len(stus))
		for _, val := range stus {
			go func(val *model.Student) {
				defer wg.Done()
				major, err1 := db.College.Where(db.College.ID.Eq(val.CollegeID)).First()
				if err1 != nil {
					logrus.Error(err1)
				}
				u := User{
					ID:          val.ID,
					Number:      val.Number,
					Name:        val.Name,
					Phone:       val.Phone,
					Email:       val.Email,
					CollegeID:   val.CollegeID,
					MajorID:     val.MajorID,
					IsStu:       1,
					CollegeName: major.Name,
					RoleId:      1,
				}
				Us = append(Us, u)
			}(val)
		}
		wg.Wait()
	} else {
		teachers, err := db.Teacher.Where(db.Teacher.CollegeID.Eq(collegeId)).Limit(size).Offset(size * (num - 1)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, nil
		}
		var wg sync.WaitGroup
		wg.Add(len(teachers))
		for _, val := range teachers {
			go func(val *model.Teacher) {
				defer wg.Done()
				college, err1 := db.College.Where(db.College.ID.Eq(val.CollegeID)).First()
				if err1 != nil {
					logrus.Error(err1)
				}
				u := User{
					ID:          val.ID,
					Number:      val.Number,
					Name:        val.Name,
					Phone:       val.Phone,
					Email:       val.Email,
					CollegeID:   val.CollegeID,
					MajorID:     val.MajorID,
					IsStu:       0,
					CollegeName: college.Name,
					RoleId:      int(val.RoleID),
				}
				Us = append(Us, u)
			}(val)
		}
		wg.Wait()
	}
	return &Us, nil
}
