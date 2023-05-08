package router

import (
	handler "github.com/Hanabi-wxl/dlu-design-system/handler/info"
	"github.com/gin-gonic/gin"
)

func RegisterSchoolRouter(engin *gin.RouterGroup) {
	info := engin.Group("/info")
	{
		info.GET("/schools", handler.InfoSchools)
	}
}
