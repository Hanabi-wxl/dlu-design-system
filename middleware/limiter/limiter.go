package limiter

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"log"
	"time"
)

// 容量为10 每100ms放一个token 最多10请求/s
var limiter = rate.NewLimiter(rate.Every(time.Second/10), 10)

func Limiter() gin.HandlerFunc {
	return func(context *gin.Context) {
		reservation := limiter.Reserve()
		waitTime := reservation.Delay().Milliseconds()
		log.Println("等待毫秒", waitTime, limiter.Limit(), limiter.Burst(), "   ", context.FullPath())
		if waitTime > 1 { // 如果有等待时间，则放弃处理当前请求
			reservation.Cancel()
			context.JSON(400, "请求频繁请稍后再试")
			context.Abort()
		}
		context.Next()
	}
}
