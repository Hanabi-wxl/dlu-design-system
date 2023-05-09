package service

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
)

type ManageService interface {
	SaveUser(s UserRequest) *errno.Errno
}
