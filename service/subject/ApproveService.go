package Service

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
)

type ApproveService interface {
	MajorApproveSubject(subjectIds []int64, approverId int64) *errno.Errno
	CollegeApproveSubject(subjectIds []int64, approverId int64) *errno.Errno
	GetApproveList(approveReq types.ApproveReq) ([]*model.Subject, *errno.Errno)
	AppointSubject(appointids []int64, userInfo *jwt.UserInfo) *errno.Errno
}
