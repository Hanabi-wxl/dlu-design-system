package router

import (
	handler "github.com/Hanabi-wxl/dlu-design-system/handler/subject"
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/consts"
	"github.com/gin-gonic/gin"
)

func RegisterSubjectRouter(engin *gin.RouterGroup) {
	subject := engin.Group("/subject")
	apply := subject.Group("/apply")
	{
		apply.POST("/", jwt.Auth(), handler.AddSubject)
		apply.PUT("/", jwt.Auth(), handler.UpdateSubject)
		apply.DELETE("/:subjectId", jwt.Auth(), handler.DeleteSubject)
		apply.GET("/origins", handler.GetOrigins)
		apply.GET("/types", handler.GetTypes)
		apply.GET("/:year", jwt.Auth(), handler.GetSelfSubjectsByYear)
	}
	approve := subject.Group("/approve")
	{
		approve.POST("/approvelist", jwt.NeedAuth(consts.MajorPermission), handler.GetApproveList)
		approve.POST("/major", jwt.NeedAuth(consts.MajorPermission), handler.MajorApproveSubject)
		approve.POST("/college", jwt.NeedAuth(consts.CollegePermission), handler.CollegeApproveSubject)
		approve.POST("/appoint", jwt.NeedAuth(consts.MajorPermission), handler.AppointSubject)
	}
	appoint := subject.Group("/appoint")
	{
		appoint.GET(":number", handler.GetAppointList)
		appoint.POST("", handler.SetAppoint)
	}
}
