package router

import (
	"github.com/Hanabi-wxl/dlu-design-system/middleware/limiter"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/logger"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Init() {
	engine := gin.Default()
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logger.Logger)
	engine.Use(limiter.Limiter())
	api := engine.Group("/api")
	RegisterUserRouter(api)
	RegisterInfoRouter(api)
	engine.Run()
}
