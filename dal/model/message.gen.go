// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameMessage = "message"

// Message mapped from table <message>
type Message struct {
	ID               int64                 `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`            // id
	Content          string                `gorm:"column:content;type:varchar(128);not null" json:"content"`                  // 内容
	SendID           int64                 `gorm:"column:send_id;type:int(11);not null" json:"send_id"`                       // 发件人
	MajorID          int64                 `gorm:"column:major_id;type:int(11);not null" json:"major_id"`                     // 专业id
	ReceiveStudentID int64                 `gorm:"column:receive_student_id;type:int(11);not null" json:"receive_student_id"` // 接收学生
	ReceiveTeacherID int64                 `gorm:"column:receive_teacher_id;type:int(11);not null" json:"receive_teacher_id"` // 接收老师
	CreatedAt        time.Time             `gorm:"column:created_at;type:datetime;not null" json:"created_at"`                // 发布时间
	State            int8                  `gorm:"column:state;type:tinyint(4);not null" json:"state"`                        // 消息状态
	IsDelete         soft_delete.DeletedAt `gorm:"column:is_delete;type:int unsigned;softDelete:flag" json:"is_delete"`       // 存在标志
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}
