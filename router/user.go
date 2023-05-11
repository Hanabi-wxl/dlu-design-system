package router

import (
	handler "github.com/Hanabi-wxl/dlu-design-system/handler/user"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(engin *gin.RouterGroup) {
	user := engin.Group("/user")
	loginGroup := user.Group("/login")
	{
		loginGroup.GET("/role/:number", handler.CheckLoginRole)
		loginGroup.POST("/student", handler.StudentLogin)
		loginGroup.POST("/teacher", handler.TeacherLogin)
	}

	manageGroup := user.Group("/manage")
	manageGroup.Use(jwt.Auth(consts.SchoolPermission))
	{
		manageGroup.POST("", handler.SaveUser)
		manageGroup.PUT("/:id/:isStu/password", handler.ResetPassword)
		manageGroup.PUT("/role", handler.UpdateTeacherRole)
		manageGroup.PUT("", handler.UpdateUser)
		manageGroup.DELETE("/:id/:isStu", handler.DeleteUser)
	}
}
