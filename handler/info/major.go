package handler

import (
	"net/http"
	"strconv"

	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
)

// GetMajorList
//
//	@Description: 分页查询专业 参数: PageRequest
//	@param c
func GetMajorList(c *gin.Context) {
	var page PageRequest
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
	var major model.Major
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
	var idRequest IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	id, err := strconv.ParseInt(idRequest.Id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err2 := service.GetInfoService().DeleteMajor(id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// UpdateMajor
//
//	@Description: 更新专业信息 参数: major (JSON)
//	@param c
func UpdateMajor(c *gin.Context) {
	var major model.Major
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
	var idRequest IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	id, err := strconv.ParseInt(idRequest.Id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	major, err2 := service.GetInfoService().GetMajor(id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(major))
	}
}

// GetMajorListByCollegeId
//
//	@Description: 分页查询专业 参数: CollegeId
//	@param c
func GetMajorListByCollegeId(c *gin.Context) {
	var collegeId types.CollegeId
	if err := c.ShouldBindUri(&collegeId); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	id, err := strconv.ParseInt(collegeId.CollegeId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	list, err2 := service.GetInfoService().GetMajorListByCollegeId(id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}
