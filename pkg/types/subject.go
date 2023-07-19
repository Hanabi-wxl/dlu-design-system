package types

import (
	"gorm.io/plugin/soft_delete"
)

type SubjectRequest struct {
	ID              int64                 `json:"id"`                // id
	Headline        string                `json:"headline"`          // 标题
	FirstTeacherID  int64                 `json:"first_teacher_id"`  // 第一指导老师
	SecondTeacherID int64                 `json:"second_teacher_id"` // 第二指导老师
	StudentID       int64                 `json:"student_id"`        // 学生id
	ProgressID      int64                 `json:"progress_id"`       // 题目状态
	CollegeID       int64                 `json:"college_id"`        // 学院
	MajorID         []int64               `json:"major_id"`          // 专业id
	RoleID          int64                 `json:"role_id"`           // 申报人角色
	Year            int64                 `json:"year"`              // 年份
	Abstract        string                `json:"abstract"`          // 摘要
	Necessity       string                `json:"necessity"`         // 选题的必要性
	Feasibility     string                `json:"feasibility"`       // 选题的可行性
	Identical       int8                  `json:"identical"`         // 三年内是否有雷同题目
	TypeID          int8                  `json:"type_id"`           // 毕业论文类型
	OriginID        int8                  `json:"origin_id"`         // 题目来源
	NeedTotal       int64                 `json:"need_total"`        // 拟需人数
	IsDelete        soft_delete.DeletedAt `json:"is_delete"`
}
type ApproveReq struct {
	Year       int64  `gorm:"column:year" json:"year" `
	TeacherId  int64  `gorm:"column:first_teacher_id" json:"teacherId"`
	CollegeId  int64  `gorm:"column:college_id" json:"collegeId"`
	MajorId    int64  `gorm:"column:major_id" json:"majorId"`
	ProgressId int64  `gorm:"column:progress_id"`
	Condition  string `json:"condition"`
}
