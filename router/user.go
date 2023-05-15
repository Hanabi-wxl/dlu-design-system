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
	manageGroup.Use(jwt.NeedAuth(consts.SchoolPermission))
	{
		manageGroup.POST("", handler.AddUser)
		manageGroup.PUT("/:id/:isStu/password", handler.ResetPassword)
		manageGroup.PUT("/role", handler.UpdateTeacherRole)
		manageGroup.PUT("", handler.UpdateUser)
		manageGroup.DELETE("/:id/:isStu", handler.DeleteUser)
	}
	user.GET("/manager/:roleId", jwt.NeedAuth(consts.SchoolPermission), handler.GetManagerList)
	user.PUT("/manager/:id", jwt.NeedAuth(consts.SchoolPermission), handler.DeleteManager)

	user.GET("/roles", jwt.NeedAuth(consts.SchoolPermission), handler.GetRoleList)
	user.POST("/role", jwt.NeedAuth(consts.SchoolPermission), handler.AddRole)
	user.GET("/role/:roleId", jwt.NeedAuth(consts.SchoolPermission), handler.GetRole)
	user.PUT("/role", jwt.NeedAuth(consts.SchoolPermission), handler.UpdateRole)
	user.DELETE("/role/:roleId", jwt.NeedAuth(consts.SchoolPermission), handler.DeleteRole)

	infoGroup := user.Group("/info")
	infoGroup.Use(jwt.Auth())
	{
		infoGroup.GET("/:id/:isStu", handler.GetUserById)
		infoGroup.GET("/number/:number/:isStu", handler.GetUserByNumber)
		infoGroup.GET("/numberMajor/:number/:majorId/:isStu", jwt.NeedAuth(consts.SchoolPermission), handler.GetUserByNumberMajor)
		infoGroup.GET("/major/:majorId/:isStu/:size/:num", handler.GetUsersByMajor)
		infoGroup.GET("/college/:collegeId/:isStu/:size/:num", handler.GetUsersByCollege)
	}
}
