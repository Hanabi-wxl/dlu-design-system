// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameClass = "class"

// Class mapped from table <class>
type Class struct {
	ID       int64                 `gorm:"column:id;type:int(11);primaryKey" json:"id"`                         // id
	Name     string                `gorm:"column:name;type:varchar(16);not null" json:"name"`                   // 名称
	MajorID  int64                 `gorm:"column:major_id;type:int(11);not null" json:"major_id"`               // 专业id
	Grade    int64                 `gorm:"column:grade;type:int(11);not null" json:"grade"`                     // 年级
	IsDelete soft_delete.DeletedAt `gorm:"column:is_delete;type:int unsigned;softDelete:flag" json:"is_delete"` // 是否有效
}

// TableName Class's table name
func (*Class) TableName() string {
	return TableNameClass
}
