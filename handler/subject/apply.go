package handler

import (
	"net/http"
	"strconv"

	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	Service "github.com/Hanabi-wxl/dlu-design-system/service/subject"
	"github.com/gin-gonic/gin"
)

// AddSubject
//
//	@Description: 新增报题
//	@param c
func AddSubject(c *gin.Context) {
	var subjectRequest types.SubjectRequest
	if err := c.ShouldBindJSON(&subjectRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	count, err := Service.GetSubjectService().CountSubject(subjectRequest.FirstTeacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// 老师报题数量超过12条报题失败
	if count > 13 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	subject := types.SubjectRequestTosubject(subjectRequest)
	err2 := Service.GetSubjectService().AddSubject(&subject)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

func UpdateSubject(c *gin.Context) {
	var subjectRequest types.SubjectRequest
	userInfo := jwt.GetUserInfo(c.Request.Context())
	err := c.ShouldBindJSON(&subjectRequest)
	if userInfo.ID == subjectRequest.StudentID {
		if err != nil {
			c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
			return
		}
		subject := types.SubjectRequestTosubject(subjectRequest)
		err2 := Service.GetSubjectService().UpdateSubject(&subject)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, err2)
		} else {
			c.JSON(http.StatusOK, result.Ok())
		}
	} else {
		c.JSON(http.StatusBadRequest, result.Failed())
	}
}

func DeleteSubject(c *gin.Context) {

}

func GetOrigins(c *gin.Context) {

}
func GetTypes(c *gin.Context) {

}
func GetSelfSubjectsByYear(c *gin.Context) {
	var year Year
	userInfo := jwt.GetUserInfo(c.Request.Context())
	if err := c.ShouldBindUri(&year); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	yearDate,_ := strconv.ParseInt(year.year, 10, 64)
	list, err := Service.GetSubjectService().GetSelfSubjectsByYear(yearDate, userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}
