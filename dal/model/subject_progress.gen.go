// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSubjectProgress = "subject_progress"

// SubjectProgress mapped from table <subject_progress>
type SubjectProgress struct {
	ID   int64  `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"` // id
	Name string `gorm:"column:name;type:varchar(16);not null" json:"name"`              // 名称
}

// TableName SubjectProgress's table name
func (*SubjectProgress) TableName() string {
	return TableNameSubjectProgress
}
