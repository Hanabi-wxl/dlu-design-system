package router

import (
	"github.com/Hanabi-wxl/dlu-design-system/middleware/cors"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/limiter"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func Init() {
	engine := gin.Default()
	engine.ForwardedByClientIP = true
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.Logger)
	// 配置跨域 限流
	engine.Use(cors.Cors(), limiter.Limiter())
	api := engine.Group("/api")
	RegisterUserRouter(api)
	RegisterInfoRouter(api)
	err := engine.Run()
	if err != nil {
		logrus.Fatal("服务启动失败")
	}
}
