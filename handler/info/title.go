package handler

import (
	"net/http"
	"strconv"

	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
)

// GetTitles
//
//	@Description: 查询所有职称信息 参数:
//	@param c
func GetTitles(c *gin.Context) {
	infoService := service.GetInfoService()
	titles, err2 := infoService().GetTitles()
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(titles))
	}
}

// GetTitle
//
//	@Description: 根据id查询单个职称信息 参数: IdRequest
//	@param c
func GetTitle(c *gin.Context) {
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
	title, err2 := infoService().GetTitle(id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(title))
	}
}
