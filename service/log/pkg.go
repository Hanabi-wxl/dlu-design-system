package service

import "sync"

var logService *LogServiceImpl
var logServiceOnce sync.Once

func GetLogService() *LogServiceImpl {
	logServiceOnce.Do(func() {
		logService = new(LogServiceImpl)
	})
	return logService
}
