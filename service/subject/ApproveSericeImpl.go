package Service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal"
	"github.com/Hanabi-wxl/dlu-design-system/dal/db"
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	"github.com/sirupsen/logrus"
	"time"
	"unicode"
)

type ApproveServiceImpl struct {
}

func (c ApproveServiceImpl) MajorApproveSubject(subjectIds []int64, approverId int64) *errno.Errno {
	_, err := db.Subject.Where(db.Subject.ID.In(subjectIds...)).Updates(map[string]any{
		"major_approval_opinion": "同意",
		"major_approval_time":    time.Now(),
		"major_approver_id":      approverId,
		"progress_id":            2,
	})
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c ApproveServiceImpl) CollegeApproveSubject(subjectIds []int64, approverId int64) *errno.Errno {
	_, err := db.Subject.Where(db.Subject.ID.In(subjectIds...)).Updates(map[string]any{
		"college_approval_opinion": "同意",
		"college_approval_time":    time.Now(),
		"college_approver_id":      approverId,
		"progress_id":              3,
	})
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}
func (c ApproveServiceImpl) GetApproveList(approveReq types.ApproveReq) ([]*model.Subject, *errno.Errno) {
	var teacherIds []int64
	tx := dal.DB.Where(&types.ApproveReq{
		Year:       approveReq.Year,
		TeacherId:  approveReq.TeacherId,
		CollegeId:  approveReq.CollegeId,
		MajorId:    approveReq.MajorId,
		ProgressId: approveReq.ProgressId,
	})
	if len(approveReq.Condition) == 0 {
	} else if unicode.IsDigit([]rune(approveReq.Condition)[0]) {
		// 条件是数字
		err := db.Teacher.Where(db.Teacher.Number.Like("%"+approveReq.Condition+"%")).Pluck(db.Teacher.ID, &teacherIds)
		if err != nil {
			logrus.Error(err)
			return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
		}
		tx.Where("first_teacher_id in ?", teacherIds)
	} else {
		// 条件为名字
		err := db.Teacher.Where(db.Teacher.Name.Like("%"+approveReq.Condition+"%")).Pluck(db.Teacher.ID, &teacherIds)
		if err != nil {
			logrus.Error(err)
			return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
		}
		tx.Where("first_teacher_id in ?", teacherIds)
	}
	var subject []*model.Subject
	if result := tx.Find(&subject); result.Error != nil {
		logrus.Error(result.Error)
		return nil, errno.NewErrno(errno.DbErrorCode, result.Error.Error())
	}
	return subject, nil
}
func (c ApproveServiceImpl) AppointSubject(appointids []int64, userInfo *jwt.UserInfo) *errno.Errno {
	userId := userInfo.ID
	_, err := db.Appoint.Where(db.Appoint.ID.In(appointids...)).Updates(map[string]interface{}{
		"approve_id":   userId,
		"approve_time": time.Now(),
	})
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}
