package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetMajorList
//
//	@Description: 分页查询专业 参数: PageRequest
//	@param c
func GetMajorList(c *gin.Context) {
	var page types.PageRequest
	if err := c.ShouldBindUri(&page); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	list, err := service.GetInfoService().GetMajorList(page.Size, page.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}

// AddMajor
//
//	@Description: 新增专业信息 参数: major (JSON)
//	@param c
func AddMajor(c *gin.Context) {
	var major types.MajorReq
	if err := c.ShouldBindJSON(&major); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err := service.GetInfoService().AddMajor(&major)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// DeleteMajor
//
//	@Description: 删除专业信息 参数: IdRequest
//	@param c
func DeleteMajor(c *gin.Context) {
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err := service.GetInfoService().DeleteMajor(idRequest.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// UpdateMajor
//
//	@Description: 更新专业信息 参数: major (JSON)
//	@param c
func UpdateMajor(c *gin.Context) {
	var major types.MajorReq
	if err := c.ShouldBindJSON(&major); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err2 := service.GetInfoService().UpdateMajor(&major)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// GetMajor
//
//	@Description: 根据id查询单个专业信息 参数: IdRequest
//	@param c
func GetMajor(c *gin.Context) {
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	major, err := service.GetInfoService().GetMajor(idRequest.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(major))
	}
}

// GetMajorListByCollegeId
//
//	@Description: 分页查询专业 参数: collegeId
//	@param c
func GetMajorListByCollegeId(c *gin.Context) {
	var collegeIdReq types.CollegeIdReq
	if err := c.ShouldBindUri(&collegeIdReq); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	list, err2 := service.GetInfoService().GetMajorListByCollegeId(collegeIdReq.CollegeId)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}
