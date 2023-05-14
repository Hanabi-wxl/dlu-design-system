package service

import (
	"context"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
)

type LogService interface {
	// AddLog
	// @Description: 创建日志
	// @param content 内容
	// @param c 上下文
	// @param ip ip地址
	// @param state 状态 1正常 2错误
	// @param method 1增 2删 3改 4查
	AddLog(content string, c context.Context, ip string, state int8, method int8)

	// LogList
	// @Description: 获取日志列表
	// @param size 分页大小
	// @param num 页码
	// @return *[]model.Log
	// @return *errno.Errno
	LogList(size int, num int) (*types.PageResp, *errno.Errno)

	// LogDateList
	// @Description: 根据日期获取日志
	// @param req
	// @return *[]model.Log
	// @return *errno.Errno
	LogDateList(req *types.LogDateReq) (*types.PageResp, *errno.Errno)

	// LogStateList
	// @Description: 根据状态获取日志
	// @param req
	// @return *[]model.Log
	// @return *errno.Errno
	LogStateList(req *types.LogStateReq) (*types.PageResp, *errno.Errno)

	// LogMethodList
	// @Description: 根据方式获取日志
	// @param req
	// @return *[]model.Log
	// @return *errno.Errno
	LogMethodList(req *types.LogMethodReq) (*types.PageResp, *errno.Errno)
}
