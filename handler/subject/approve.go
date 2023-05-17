package handler

import (
	"net/http"
	"strconv"

	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	Service "github.com/Hanabi-wxl/dlu-design-system/service/subject"
	"github.com/gin-gonic/gin"
)

func GetApproveList(c *gin.Context) {

}

func MajorApproveSubject(c *gin.Context) {
	var majorIds []MajorIds
	if err := c.ShouldBindJSON(&majorIds); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	var subjects []int64
	for k, v := range majorIds {
		id, _ := strconv.ParseInt(v.majorId, 10, 64)
		subjects[k] = id
	}
	err2 := Service.GetSubjectService().MajorApproveSubject(subjects)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

func CollegeApproveSubject(c *gin.Context) {
	var collegeIds []CollegeIds
	if err := c.ShouldBindJSON(&collegeIds); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	var subjects []int64
	for k, v := range collegeIds {
		id, _ := strconv.ParseInt(v.collegeId, 10, 64)
		subjects[k] = id
	}
	err2 := Service.GetSubjectService().CollegeApproveSubject(subjects)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

func AppointSubject(c *gin.Context) {

}
