package router

import (
	"github.com/Hanabi-wxl/dlu-design-system/middleware/cors"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/limiter"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func Init() {
	engine := gin.Default()
	store := cookie.NewStore([]byte("dlu-design-system-cookie"))
	engine.Use(sessions.Sessions("dlu-design-system-session", store))
	engine.ForwardedByClientIP = true
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.Logger)
	// 配置跨域 限流
	engine.Use(cors.Cors(), limiter.Limiter())
	// 总路由
	api := engine.Group("/api")
	// 注册用户模块路由
	RegisterUserRouter(api)
	// 注册信息模块路由
	RegisterInfoRouter(api)
	err := engine.Run()
	if err != nil {
		logrus.Fatal("服务启动失败")
	}
}
