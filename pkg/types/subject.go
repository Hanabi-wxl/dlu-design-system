package types

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"gorm.io/plugin/soft_delete"
	"time"
)

type SubjectRequest struct {
	ID                     int64                 `json:"id"`                       // id
	Headline               string                `json:"headline"`                 // 标题
	FirstTeacherID         int64                 `json:"first_teacher_id"`         // 第一指导老师
	SecondTeacherID        int64                 `json:"second_teacher_id"`        // 第二指导老师
	StudentID              int64                 `json:"student_id"`               // 学生id
	ProgressID             int64                 `json:"progress_id"`              // 题目状态
	CollegeID              int64                 `json:"college_id"`               // 学院
	MajorID                int64                 `json:"major_id"`                 // 专业id
	RoleID                 int64                 `json:"role_id"`                  // 申报人角色
	Year                   int64                 `json:"year"`                     // 年份
	Abstract               string                `json:"abstract"`                 // 摘要
	Necessity              string                `json:"necessity"`                // 选题的必要性
	Feasibility            string                `json:"feasibility"`              // 选题的可行性
	Identical              int8                  `json:"identical"`                // 三年内是否有雷同题目
	TypeID                 int8                  `json:"type_id"`                  // 毕业论文类型
	OriginID               int8                  `json:"origin_id"`                // 题目来源
	NeedTotal              int64                 `json:"need_total"`               // 拟需人数
	MajorApprovalOpinion   string                `json:"major_approval_opinion"`   // 专业审批意见
	MajorApprovalTime      time.Time             `json:"major_approval_time"`      // 专业审批通过时间
	MajorApproverID        int64                 `json:"major_approver_id"`        // 专业审批人
	CollegeApprovalOpinion string                `json:"college_approval_opinion"` // 学院审批意见
	CollegeApprovalTime    time.Time             `json:"college_approval_time"`    // 学院审批通过时间
	CollegeApproverID      int64                 `json:"college_approver_id"`      // 学院审批人
	IsDelete               soft_delete.DeletedAt `json:"is_delete"`
}

// SubjectRequestTosubject
//
//	 @author: 567jiong
//		@Description: 将 SubjectRequest转为 subject
//		@param subjectRequest
//		@return model.Subject
func SubjectRequestTosubject(subjectRequest SubjectRequest) model.Subject {
	return model.Subject{
		ID:                     subjectRequest.ID,
		Headline:               subjectRequest.Headline,
		FirstTeacherID:         subjectRequest.FirstTeacherID,
		SecondTeacherID:        &(subjectRequest.SecondTeacherID),
		StudentID:              &(subjectRequest.StudentID),
		ProgressID:             subjectRequest.ProgressID,
		CollegeID:              subjectRequest.CollegeID,
		MajorID:                &(subjectRequest.MajorID),
		RoleID:                 &(subjectRequest.RoleID),
		Year:                   subjectRequest.Year,
		Abstract:               subjectRequest.Abstract,
		Necessity:              subjectRequest.Necessity,
		Feasibility:            subjectRequest.Feasibility,
		Identical:              subjectRequest.Identical,
		TypeID:                 subjectRequest.TypeID,
		OriginID:               subjectRequest.OriginID,
		NeedTotal:              subjectRequest.NeedTotal,
		MajorApprovalOpinion:   &(subjectRequest.MajorApprovalOpinion),
		MajorApprovalTime:      &(subjectRequest.MajorApprovalTime),
		MajorApproverID:        &(subjectRequest.MajorApproverID),
		CollegeApprovalOpinion: &(subjectRequest.CollegeApprovalOpinion),
		CollegeApprovalTime:    &(subjectRequest.CollegeApprovalTime),
		CollegeApproverID:      &(subjectRequest.CollegeApproverID),
		IsDelete:               subjectRequest.IsDelete,
	}
}
