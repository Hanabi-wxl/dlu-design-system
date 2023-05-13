package service

import (
	"context"
)

type LogService interface {
	AddLog(content string, c context.Context, ip string, state int8, method int8)
}
