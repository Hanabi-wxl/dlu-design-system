package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCollegeList
//
//	@Description: 分页查询学院 参数: PageRequest
//	@param c
func GetCollegeList(c *gin.Context) {
	var page types.PageRequest
	if err := c.ShouldBindUri(&page); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	list, err := service.GetInfoService().GetCollegeList(page.Size, page.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}

// AddCollege
//
//	@Description: 新增学院信息 参数: college (JSON)
//	@param c
func AddCollege(c *gin.Context) {
	var college types.CollegeReq
	if err := c.ShouldBindJSON(&college); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err := service.GetInfoService().AddCollege(&college)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// DeleteCollege
//
//	@Description: 删除学院信息 参数: IdRequest
//	@param c
func DeleteCollege(c *gin.Context) {
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err := service.GetInfoService().DeleteCollege(idRequest.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// UpdateCollege
//
//	@Description: 更新学院信息 参数: college (JSON)
//	@param c
func UpdateCollege(c *gin.Context) {
	var college types.CollegeReq
	if err := c.ShouldBindJSON(&college); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err2 := service.GetInfoService().UpdateCollege(&college)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// GetCollege
//
//	@Description: 根据id查询单个学院信息 参数: IdRequest
//	@param c
func GetCollege(c *gin.Context) {
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	college, err := service.GetInfoService().GetCollege(idRequest.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(college))
	}
}
