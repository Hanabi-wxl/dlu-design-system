package router

import (
	handler "github.com/Hanabi-wxl/dlu-design-system/handler/info"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	"github.com/gin-gonic/gin"
)

func RegisterInfoRouter(engin *gin.RouterGroup) {
	info := engin.Group("/info")
	{
		info.GET("/schools/:size/:num", handler.GetSchoolList)
		info.POST("/school", jwt.Auth(consts.SchoolPermission), handler.AddSchool)
		info.DELETE("/school/:id", jwt.Auth(consts.SchoolPermission), handler.DeleteSchool)
		info.PUT("/school", jwt.Auth(consts.SchoolPermission), handler.UpdateSchool)
		info.GET("/school/:id", handler.GetSchool)

		info.GET("/colleges/:size/:num", handler.GetCollegeList)
		info.POST("/college", jwt.Auth(consts.SchoolPermission), handler.AddCollege)
		info.DELETE("/college/:id", jwt.Auth(consts.SchoolPermission), handler.DeleteCollege)
		info.PUT("/college", jwt.Auth(consts.SchoolPermission), handler.UpdateCollege)
		info.GET("/college/:id", handler.GetCollege)

		info.GET("/majors/:size/:num", handler.GetMajorList)
		info.POST("/major", jwt.Auth(consts.SchoolPermission), handler.AddMajor)
		info.DELETE("/major/:id", jwt.Auth(consts.SchoolPermission), handler.DeleteMajor)
		info.PUT("/major", jwt.Auth(consts.SchoolPermission), handler.UpdateMajor)
		info.GET("/major/:id", handler.GetMajor)

		info.GET("/classes/:size/:num", handler.GetClassList)
		info.POST("/class", jwt.Auth(consts.SchoolPermission), handler.AddClass)
		info.DELETE("/class/:id", jwt.Auth(consts.SchoolPermission), handler.DeleteClass)
		info.PUT("/class", jwt.Auth(consts.SchoolPermission), handler.UpdateClass)
		info.GET("/class/:id", handler.GetClass)

		info.GET("/titles", handler.GetTitles)
		info.GET("/title/:id", handler.GetTitle)

		info.GET("/degrees", handler.GetDegrees)
		info.GET("/degree/:id", handler.GetDegree)

		info.GET("/offices", handler.GetOffices)
		info.GET("/office/:id", handler.GetOffice)
	}
}
