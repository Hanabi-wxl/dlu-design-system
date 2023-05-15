package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetClassList
//
//	@Description: 分页查询班级 参数: PageRequest
//	@param c
func GetClassList(c *gin.Context) {
	var query types.ClassQueryReq
	if err := c.ShouldBindUri(&query); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	list, err := service.GetInfoService().GetClassList(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}

// AddClass
//
//	@Description: 新增班级信息 参数: class (JSON)
//	@param c
func AddClass(c *gin.Context) {
	var class types.ClassReq
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err := service.GetInfoService().AddClass(&class)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// DeleteClass
//
//	@Description: 删除班级信息 参数:IdRequest
//	@param c
func DeleteClass(c *gin.Context) {
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err := service.GetInfoService().DeleteClass(idRequest.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// UpdateClass
//
//	@Description: 更新班级信息 参数: class (JSON)
//	@param c
func UpdateClass(c *gin.Context) {
	var class types.ClassReq
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	err2 := service.GetInfoService().UpdateClass(&class)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}

}

// GetClass
//
//	@Description: 根据id查询单个班级信息 参数:IdRequest
//	@param c
func GetClass(c *gin.Context) {
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	class, err := service.GetInfoService().GetClass(idRequest.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(class))
	}
}
