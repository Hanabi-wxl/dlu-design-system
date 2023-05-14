package router

import (
	handler "github.com/Hanabi-wxl/dlu-design-system/handler/log"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	"github.com/gin-gonic/gin"
)

func RegisterLogRouter(engin *gin.RouterGroup) {
	log := engin.Group("/log")
	log.Use(jwt.NeedAuth(consts.SchoolPermission))
	{
		log.GET("/:size/:num", handler.GetLogList)
		log.GET("/date/:start/:end/:size/:num", handler.GetLogListByDate)
		log.GET("/method/:methodId/:size/:num", handler.GetLogListByMethod)
		log.GET("/state/:stateId/:size/:num", handler.GetLogListByState)
	}
}
