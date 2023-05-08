package limiter

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"time"
)

// 容量为10 每100ms放一个token 最多10请求/s
var limiter = rate.NewLimiter(rate.Every(time.Second/10), 10)

func Limiter() gin.HandlerFunc {
	return func(context *gin.Context) {
		reservation := limiter.Reserve()
		waitTime := reservation.Delay().Milliseconds()
		logrus.Info("等待毫秒", waitTime, limiter.Limit(), limiter.Burst(), "   ", context.FullPath())
		if waitTime > 1 { // 如果有等待时间，则放弃处理当前请求
			reservation.Cancel()
			context.JSON(400, errno.LimitRequestErrno)
			context.Abort()
		}
		context.Next()
	}
}
