// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameAppoint = "appoint"

// Appoint mapped from table <appoint>
type Appoint struct {
	ID        int64                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`      // id
	SendID    int64                 `gorm:"column:send_id;type:int(11);not null" json:"send_id"`                 // 发放权限的教师id
	ReceiveID int64                 `gorm:"column:receive_id;type:int(11);not null" json:"receive_id"`           // 接收权限的教师id
	SubjectID int64                 `gorm:"column:subject_id;type:int(11);not null" json:"subject_id"`           // 题目id
	CreatedAt time.Time             `gorm:"column:created_at;type:datetime;not null" json:"created_at"`          // 委任时间
	IsDelete  soft_delete.DeletedAt `gorm:"column:is_delete;type:int unsigned;softDelete:flag" json:"is_delete"` // 存在标志
}

// TableName Appoint's table name
func (*Appoint) TableName() string {
	return TableNameAppoint
}