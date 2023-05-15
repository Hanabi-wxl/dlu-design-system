package router

import (
	handler "github.com/Hanabi-wxl/dlu-design-system/handler/subject"
	"github.com/gin-gonic/gin"
)

func RegisterSubjectRouter(engin *gin.RouterGroup) {
	subject := engin.Group("/subject")
	apply := subject.Group("/apply")
	{
		apply.POST("/", handler.AddSubject)
		apply.PUT("/", handler.UpdateSubject)
		apply.DELETE("/:subjectId", handler.DeleteSubject)
		apply.GET("/origins", handler.GetOrigins)
		apply.GET("/types", handler.GetTypes)
		apply.GET("/:year", handler.GetSelfSubjectsByYear)
	}
	approve := subject.Group("/approve")
	{
		approve.GET("/approvelist", handler.GetApproveList)
		approve.POST("/major", handler.MajorApproveSubject)
		approve.POST("/college", handler.CollegeApproveSubject)
		approve.POST("/appoint", handler.AppointSubject)
	}
}
