package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTitles
//
//	@Description: 查询所有职称信息 参数:
//	@param c
func GetTitles(c *gin.Context) {
	titles, err2 := service.GetInfoService().GetTitles()
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
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	title, err := service.GetInfoService().GetTitle(idRequest.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(title))
	}
}
