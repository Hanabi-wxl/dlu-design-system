package Service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal"
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/sirupsen/logrus"
	"gorm.io/gen/examples/dal/query"
)

type ApplyServiceImpl struct {
}

func (c ApplyServiceImpl) AddSubject(subject *model.Subject, majorIds []int64) *errno.Errno {
	// 开启事务
	err := query.Use(dal.DB).Transaction(func(tx *query.Query) error {
		if err := db.Subject.Create(subject); err != nil {
			logrus.Error(err)
			return err
		}
		var subjectMajors []*model.SubjectMajor
		for _, v := range majorIds {
			var subjectMajor model.SubjectMajor
			subjectMajor.SubjectID = subject.ID
			subjectMajor.MajorID = v
			subjectMajor.IsDelete = 0
			subjectMajors = append(subjectMajors, &subjectMajor)
		}
		if err := db.SubjectMajor.Create(subjectMajors...); err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c ApplyServiceImpl) CountSubject(teacherId int64) (int64, *errno.Errno) {
	count, err := db.Subject.Where(db.Subject.FirstTeacherID.Eq(teacherId)).Count()
	if err != nil {
		logrus.Error(err)
		return 0, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return count, nil
}
func (c ApplyServiceImpl) CountSubjectByStudent(studentId int64) (int64, *errno.Errno) {
	count, err := db.Subject.Where(db.Subject.StudentID.Eq(studentId)).Count()
	if err != nil {
		logrus.Error(err)
		return 0, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return count, nil
}
func (c ApplyServiceImpl) UpdateSubject(subject *model.Subject, majorIds []int64) *errno.Errno {
	// 修改subject_major表 同时更新subject
	err := query.Use(dal.DB).Transaction(func(tx *query.Query) error {
		if _, err := db.Subject.Where(db.Subject.ID.Eq(subject.ID)).Updates(subject); err != nil {
			logrus.Error(err)
			return err
		}
		if _, err := db.SubjectMajor.Where(db.SubjectMajor.SubjectID.Eq(subject.ID)).Delete(); err != nil {
			logrus.Error(err)
			return err
		}
		// 添加 subject_major表
		var subjectMajors []*model.SubjectMajor
		for _, v := range majorIds {
			var subjectMajor model.SubjectMajor
			subjectMajor.SubjectID = subject.ID
			subjectMajor.MajorID = v
			subjectMajor.IsDelete = 0
			subjectMajors = append(subjectMajors, &subjectMajor)
		}
		if err := db.SubjectMajor.Create(subjectMajors...); err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c ApplyServiceImpl) GetSelfSubjectsByYear(year int64, userInfo *jwt.UserInfo) ([]*model.Subject, *errno.Errno) {
	if userInfo.RoleId == 1 {
		// 获取学生本年所选报题
		subjects, err := db.Subject.Where(db.Subject.Year.Eq(year), db.Subject.StudentID.Eq(userInfo.ID)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
		}
		return subjects, nil
	} else {
		// 获取老师本年所选报题
		subjects, err := db.Subject.Where(db.Subject.Year.Eq(year)).Or(db.Subject.FirstTeacherID.Eq(userInfo.ID), db.Subject.SecondTeacherID.Eq(userInfo.ID)).Find()
		if err != nil {
			logrus.Error(err)
			return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
		}
		return subjects, nil
	}
}
func (c ApplyServiceImpl) GetTypes() ([]*model.SubjectType, *errno.Errno) {
	types, err := db.SubjectType.Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return types, nil
}
func (c ApplyServiceImpl) GetOrigins() ([]*model.SubjectOrigin, *errno.Errno) {
	origins, err := db.SubjectOrigin.Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return origins, nil
}
func (c ApplyServiceImpl) Delete(subjectId int64) *errno.Errno {
	// 删除报题同时删除 subjectMajor表
	err := query.Use(dal.DB).Transaction(func(tx *query.Query) error {
		if _, err := db.Subject.Where(db.Subject.ID.Eq(subjectId)).Delete(); err != nil {
			logrus.Error(err)
			return err
		}
		if _, err := db.SubjectMajor.Where(db.SubjectMajor.SubjectID.Eq(subjectId)).Delete(); err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}
func (c ApplyServiceImpl) GetSubjectById(subjectId int64) (*model.Subject, *errno.Errno) {
	subject, err := db.Subject.Where(db.Subject.ID.Eq(subjectId)).Take()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return subject, nil
}
func (c ApplyServiceImpl) GetSubjectByIds(subjectIds []int64) ([]*model.Subject, *errno.Errno) {
	subject, err := db.Subject.Where(db.Subject.ID.In(subjectIds...)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return subject, nil
}
