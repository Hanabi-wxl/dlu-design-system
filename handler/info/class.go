package handler

import (
	"net/http"
	"strconv"

	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
)

// GetClassList
//
//	@Description: 分页查询班级 参数: PageRequest
//	@param c
func GetClassList(c *gin.Context) {
	var page PageRequest
	if err := c.ShouldBindUri(&page); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	infoService := service.GetInfoService()
	list, err := infoService().GetClassList(page.Size, page.Num)
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
	var class model.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	infoService := service.GetInfoService()
	err := infoService().AddClass(&class)
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
	infoService := service.GetInfoService()
	err2 := infoService().DeleteClass(id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// UpdateClass
//
//	@Description: 更新班级信息 参数: class (JSON)
//	@param c
func UpdateClass(c *gin.Context) {
	var class model.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	infoService := service.GetInfoService()
	err2 := infoService().UpdateClass(&class)
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
	infoService := service.GetInfoService()
	class, err2 := infoService().GetClass(id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(class))
	}
}
