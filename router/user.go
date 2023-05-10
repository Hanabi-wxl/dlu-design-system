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
		loginGroup.POST("/student", jwt.MustAuth(consts.StudentPermission), handler.StudentLogin)
		loginGroup.POST("/teacher", jwt.MustAuth(consts.TeacherPermission), handler.TeacherLogin)
	}
	manageGroup := user.Group("/manage")
	manageGroup.Use(jwt.Auth(consts.SchoolPermission))
	{
		manageGroup.POST("", handler.SaveUser)
	}
	infoGroup := user.Group("info")
	{
		infoGroup.GET("/:id/:isStu", handler.GetUserById)
		infoGroup.GET("/number/:number/:isStu", handler.GetUserByNumber)
		infoGroup.GET("/major/:majorId/:isStu/:size/:num", handler.GetUsersByMajor)
		infoGroup.GET("/college/:collegeId/:isStu/:size/:num", handler.GetUsersByCollege)
	}
}
