// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSubjectOrigin = "subject_origin"

// SubjectOrigin mapped from table <subject_origin>
type SubjectOrigin struct {
	ID       int64                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`      // id
	Name     string                `gorm:"column:name;type:varchar(16);not null" json:"name"`                   // 名称
	IsDelete soft_delete.DeletedAt `gorm:"column:is_delete;type:int unsigned;softDelete:flag" json:"is_delete"` // 存在标志
}

// TableName SubjectOrigin's table name
func (*SubjectOrigin) TableName() string {
	return TableNameSubjectOrigin
}
