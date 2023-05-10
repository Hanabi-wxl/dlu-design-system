package handler

import (
	"net/http"
	"strconv"

	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
)

// GetSections
//
//	@Description: 查询所有科室信息 参数:
//	@param c
func GetSections(c *gin.Context) {
	infoService := service.GetInfoService()
	sections, err2 := infoService().GetSections()
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(sections))
	}
}

// GetSection
//
//	@Description: 根据id查询单个科室信息 参数: IdRequest
//	@param c
func GetSection(c *gin.Context) {
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
	section, err2 := infoService().GetSection(id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(section))
	}
}
