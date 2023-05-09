package router

import (
	handler "github.com/Hanabi-wxl/dlu-design-system/handler/info"
	"github.com/gin-gonic/gin"
)

func RegisterInfoRouter(engin *gin.RouterGroup) {
	info := engin.Group("/info")
	{
		info.GET("/schools", handler.GetSchoolList)
		info.POST("/school", handler.AddSchool)
		info.DELETE("/school/:id", handler.DeleteSchool)
		info.PUT("/school/:id", handler.UpdateSchool)
		info.GET("/school/:id", handler.GetSchool)

		info.GET("/colleges/:size/:num", handler.GetCollegeList)
		info.POST("/college", handler.AddCollege)
		info.DELETE("/college/:id", handler.DeleteCollege)
		info.PUT("/college/:id", handler.UpdateCollege)
		info.GET("/college/:id", handler.GetCollege)

		info.GET("/majors/:size/:num", handler.GetMajorList)
		info.POST("/major", handler.AddMajor)
		info.DELETE("/major/:id", handler.DeleteMajor)
		info.PUT("/major/:id", handler.UpdateMajor)
		info.GET("/major/:id", handler.GetMajor)

		info.GET("/classes/:size/:num", handler.GetClassList)
		info.POST("/class", handler.AddClass)
		info.DELETE("/class/:id", handler.DeleteClass)
		info.PUT("/class/:id", handler.UpdateClass)
		info.GET("/class/:id", handler.GetClass)

		info.GET("/titles", handler.GetTitles)
		info.GET("/title/:id", handler.GetTitle)

		info.GET("/degrees", handler.GetDegrees)
		info.GET("/degree/:id", handler.GetDegree)

		info.GET("/sections", handler.GetSections)
		info.GET("/section/:id", handler.GetSection)
	}
}
