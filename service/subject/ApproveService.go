package Service

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type ApproveService interface {
	MajorApproveSubject(subjectIds []int64) *errno.Errno
	CollegeApproveSubject(subjectIds []int64) *errno.Errno
}
