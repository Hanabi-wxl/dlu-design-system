// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameTeacherDegree = "teacher_degree"

// TeacherDegree mapped from table <teacher_degree>
type TeacherDegree struct {
	ID       int64                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`      // id
	Name     string                `gorm:"column:name;type:varchar(16);not null" json:"name"`                   // 名称
	IsDelete soft_delete.DeletedAt `gorm:"column:is_delete;type:int unsigned;softDelete:flag" json:"is_delete"` // 存在标志
}

// TableName TeacherDegree's table name
func (*TeacherDegree) TableName() string {
	return TableNameTeacherDegree
}
