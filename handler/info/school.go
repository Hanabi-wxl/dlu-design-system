package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetSchoolList
//
//	@Description: 分页查询学校 参数: PageRequest
//	@param c
func GetSchoolList(c *gin.Context) {
	var page types.PageRequest
	if err := c.ShouldBindUri(&page); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	list, err := service.GetInfoService().GetSchoolList(page.Size, page.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}

// AddSchool
//
//	@Description: 新增学校信息 参数: school (JSON)
//	@param c
func AddSchool(c *gin.Context) {
	var school model.School
	if err := c.ShouldBindJSON(&school); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err := service.GetInfoService().AddSchool(&school)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// DeleteSchool
//
//	@Description: 删除学校信息 参数: IdRequest
//	@param c
func DeleteSchool(c *gin.Context) {
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	resp, errno := service.GetInfoService().GetSchool(idRequest.Id)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
		return
	}
	if resp == nil {
		// 学校不存在
		c.JSON(http.StatusBadRequest, result.Failed())
	}
	err := service.GetInfoService().DeleteSchool(idRequest.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// UpdateSchool
//
//	@Description: 更新学校信息 参数: school (JSON)
//	@param c
func UpdateSchool(c *gin.Context) {
	var school model.School
	if err := c.ShouldBindJSON(&school); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err2 := service.GetInfoService().UpdateSchool(&school)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// GetSchool
//
//	@Description: 根据id查询单个学校信息 参数: IdRequest
//	@param c
func GetSchool(c *gin.Context) {
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	school, err := service.GetInfoService().GetSchool(idRequest.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(school))
	}
}
