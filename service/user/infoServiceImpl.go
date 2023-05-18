package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	"sync"
)

type InfoServiceImpl struct {
}

func (InfoServiceImpl) GetUserById(id int64, isStu int8) (*types.UserResp, *errno.Errno) {
	var user types.UserResp
	if isStu == 1 {
		student := db.Student
		err := student.Select(student.ALL, db.Major.Name.As("MajorName"), db.College.Name.As("CollegeName"),
			db.Class.Name.As("ClassName")).Omit(student.Password).
			Where(student.ID.Eq(id)).
			LeftJoin(db.Major, db.Major.ID.EqCol(student.MajorID)).
			LeftJoin(db.College, db.College.ID.EqCol(student.CollegeID)).
			LeftJoin(db.Class, db.Class.ID.EqCol(student.ClassID)).
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
			Where(teacher.ID.Eq(id)).
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

func (InfoServiceImpl) GetUserByNumber(number string, isStu int8) (*types.UserResp, *errno.Errno) {
	var user types.UserResp
	if isStu == 1 {
		student := db.Student
		err := student.Select(student.ALL, db.Major.Name.As("MajorName"), db.College.Name.As("CollegeName"),
			db.Class.Name.As("ClassName")).Omit(student.Password).
			Where(student.Number.Eq(number)).
			LeftJoin(db.Major, db.Major.ID.EqCol(student.MajorID)).
			LeftJoin(db.College, db.College.ID.EqCol(student.CollegeID)).
			LeftJoin(db.Class, db.Class.ID.EqCol(student.ClassID)).
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
			Where(teacher.Number.Eq(number)).
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

func (InfoServiceImpl) GetUserByNumberMajor(req *types.NumberMajorRequest) (*types.UserResp, *errno.Errno) {
	var user types.UserResp
	if req.IsStu == 1 {
		student := db.Student
		err := student.Select(student.ALL, db.Major.Name.As("MajorName"), db.College.Name.As("CollegeName"),
			db.Class.Name.As("ClassName")).Omit(student.Password).
			Where(student.Number.Eq(req.Number), student.MajorID.Eq(req.MajorId)).
			LeftJoin(db.Major, db.Major.ID.EqCol(student.MajorID)).
			LeftJoin(db.College, db.College.ID.EqCol(student.CollegeID)).
			LeftJoin(db.Class, db.Class.ID.EqCol(student.ClassID)).
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

func (InfoServiceImpl) GetUsersByMajor(req *types.UserQueryByMajorReq) (*[]*types.UserResp, *errno.Errno) {
	var users []*types.UserResp
	if req.IsStu == 1 {
		stus, err := db.Student.Where(db.Student.MajorID.Eq(req.MajorId)).Limit(req.Size).Offset(req.Size * (req.Num - 1)).Find()
		if err != nil {
			return nil, nil
		}
		var wg sync.WaitGroup
		wg.Add(len(stus))
		for _, val := range stus {
			go func(val *model.Student) {
				defer wg.Done()
				major, _ := db.Major.Where(db.Major.ID.Eq(val.MajorID)).Take()
				college, _ := db.College.Where(db.College.ID.Eq(val.CollegeID)).Take()
				class, _ := db.College.Where(db.Class.ID.Eq(val.ClassID)).Take()
				u := &types.UserResp{
					ID:          val.ID,
					Number:      val.Number,
					Name:        val.Name,
					Phone:       val.Phone,
					Email:       val.Email,
					CollegeID:   college.ID,
					CollegeName: college.Name,
					MajorID:     major.ID,
					MajorName:   major.Name,
					IsStu:       1,
					Student:     types.Student{ClassID: class.ID, ClassName: class.Name},
				}
				users = append(users, u)
			}(val)
		}
		wg.Wait()
	} else {
		teachers, err := db.Teacher.Where(db.Teacher.MajorID.Eq(req.MajorId)).Limit(req.Size).Offset(req.Size * (req.Num - 1)).Find()
		if err != nil {
			return nil, nil
		}
		var wg sync.WaitGroup
		wg.Add(len(teachers))
		for _, val := range teachers {
			go func(val *model.Teacher) {
				defer wg.Done()
				major, _ := db.Major.Where(db.Major.ID.Eq(val.MajorID)).Take()
				college, _ := db.College.Where(db.College.ID.Eq(val.CollegeID)).Take()
				degree, _ := db.TeacherDegree.Where(db.TeacherDegree.ID.Eq(val.DegreeID)).Take()
				office, _ := db.TeacherOffice.Where(db.TeacherOffice.ID.Eq(val.OfficeID)).Take()
				title, _ := db.TeacherTitle.Where(db.TeacherTitle.ID.Eq(val.TitleID)).Take()
				teacher := &types.UserResp{
					ID:          val.ID,
					Number:      val.Number,
					Name:        val.Name,
					Phone:       val.Phone,
					Email:       val.Email,
					CollegeID:   college.ID,
					CollegeName: college.Name,
					MajorID:     major.ID,
					MajorName:   major.Name,
					IsStu:       1,
					Teacher:     types.Teacher{DegreeID: degree.ID, DegreeName: degree.Name, TitleID: title.ID, TitleName: title.Name, OfficeID: office.ID, OfficeName: office.Name},
				}
				users = append(users, teacher)
			}(val)
		}
		wg.Wait()
	}
	return &users, nil
}

func (InfoServiceImpl) GetUsersByCollege(req *types.UserQueryByCollegeReq) (*[]*types.UserResp, *errno.Errno) {
	var users []*types.UserResp
	if req.IsStu == 1 {
		stus, err := db.Student.Where(db.Student.CollegeID.Eq(req.CollegeId)).Limit(req.Size).Offset(req.Size * (req.Num - 1)).Find()
		if err != nil {
			return nil, nil
		}
		var wg sync.WaitGroup
		wg.Add(len(stus))
		for _, val := range stus {
			go func(val *model.Student) {
				defer wg.Done()
				major, _ := db.Major.Where(db.Major.ID.Eq(val.MajorID)).Take()
				college, _ := db.College.Where(db.College.ID.Eq(val.CollegeID)).Take()
				class, _ := db.College.Where(db.Class.ID.Eq(val.ClassID)).Take()
				u := &types.UserResp{
					ID:          val.ID,
					Number:      val.Number,
					Name:        val.Name,
					Phone:       val.Phone,
					Email:       val.Email,
					CollegeID:   college.ID,
					CollegeName: college.Name,
					MajorID:     major.ID,
					MajorName:   major.Name,
					IsStu:       1,
					Student:     types.Student{ClassID: class.ID, ClassName: class.Name},
				}
				users = append(users, u)
			}(val)
		}
		wg.Wait()
	} else {
		teachers, err := db.Teacher.Where(db.Teacher.CollegeID.Eq(req.CollegeId)).Limit(req.Size).Offset(req.Size * (req.Num - 1)).Find()
		if err != nil {
			return nil, nil
		}
		var wg sync.WaitGroup
		wg.Add(len(teachers))
		for _, val := range teachers {
			go func(val *model.Teacher) {
				defer wg.Done()
				major, _ := db.Major.Where(db.Major.ID.Eq(val.MajorID)).Take()
				college, _ := db.College.Where(db.College.ID.Eq(val.CollegeID)).Take()
				degree, _ := db.TeacherDegree.Where(db.TeacherDegree.ID.Eq(val.DegreeID)).Take()
				office, _ := db.TeacherOffice.Where(db.TeacherOffice.ID.Eq(val.OfficeID)).Take()
				title, _ := db.TeacherTitle.Where(db.TeacherTitle.ID.Eq(val.TitleID)).Take()
				teacher := &types.UserResp{
					ID:          val.ID,
					Number:      val.Number,
					Name:        val.Name,
					Phone:       val.Phone,
					Email:       val.Email,
					CollegeID:   college.ID,
					CollegeName: college.Name,
					MajorID:     major.ID,
					MajorName:   major.Name,
					IsStu:       1,
					Teacher:     types.Teacher{DegreeID: degree.ID, DegreeName: degree.Name, TitleID: title.ID, TitleName: title.Name, OfficeID: office.ID, OfficeName: office.Name},
				}
				users = append(users, teacher)
			}(val)
		}
		wg.Wait()
	}
	return &users, nil
}
