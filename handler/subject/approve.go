package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	Service "github.com/Hanabi-wxl/dlu-design-system/service/subject"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetApproveList
//
//	@Description: 获取报题审批列表
//	@param c
func GetApproveList(c *gin.Context) {
	var approveReq types.ApproveReq
	if err := c.ShouldBindJSON(&approveReq); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	userInfo := jwt.GetUserInfo(c.Request.Context())
	// 判断当前角色为专业管理员
	if userInfo.RoleId == 3 {
		approveReq.ProgressId = 1
	}
	// 判断当前角色为学院管理员
	if userInfo.RoleId == 4 {
		approveReq.ProgressId = 2
	}
	approveList, err := Service.GetSubjectService().GetApproveList(approveReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(approveList))
		return
	}
}

func MajorApproveSubject(c *gin.Context) {
	var subjectIdRequests types.SubjectIdsRequest
	if err := c.ShouldBindJSON(&subjectIdRequests); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	subjects, err1 := Service.GetSubjectService().GetSubjectByIds(subjectIdRequests.SubjectIds)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
		return
	}
	for _, v := range subjects {
		if v.ProgressID != 1 {
			// 不在专业审批的范围
			c.JSON(http.StatusBadRequest, result.Failed())
			return
		}
	}
	if err := Service.GetSubjectService().MajorApproveSubject(subjectIdRequests.SubjectIds, jwt.GetUserInfo(c.Request.Context()).ID); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, result.Ok())
}

func CollegeApproveSubject(c *gin.Context) {
	var subjectIdRequests types.SubjectIdsRequest
	if err := c.ShouldBindJSON(&subjectIdRequests); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	subjects, err1 := Service.GetSubjectService().GetSubjectByIds(subjectIdRequests.SubjectIds)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
		return
	}
	for _, v := range subjects {
		if v.ProgressID != 2 {
			// 不在学院审批的范围
			c.JSON(http.StatusBadRequest, result.Failed())
			return
		}
	}
	if err := Service.GetSubjectService().CollegeApproveSubject(subjectIdRequests.SubjectIds, jwt.GetUserInfo(c.Request.Context()).ID); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, result.Ok())
}

func AppointSubject(c *gin.Context) {
	var appointIdsRequest types.AppointIdsRequest
	if err := c.ShouldBindJSON(&appointIdsRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err := Service.GetSubjectService().AppointSubject(appointIdsRequest.Appointids, jwt.GetUserInfo(c.Request.Context()))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, result.Ok())
		return
	}
}
