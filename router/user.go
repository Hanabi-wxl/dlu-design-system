package router

import (
	handler "github.com/Hanabi-wxl/dlu-design-system/handler/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(engin *gin.RouterGroup) {
	user := engin.Group("/user")
	loginGroup := user.Group("/login")
	{
		loginGroup.GET("/role", handler.Role)
	}
	manageGroup := user.Group("/manage")
	{
		manageGroup.POST("", handler.SaveUser)
	}
}
