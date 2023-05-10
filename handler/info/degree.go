package handler

import (
	"net/http"
	"strconv"

	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
)

// GetDegrees
//
//	@Description: 查询单个学位信息 参数:
//	@param c
func GetDegrees(c *gin.Context) {
	infoService := service.GetInfoService()
	degrees, err2 := infoService().GetDegrees()
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(degrees))
	}
}

// GetDegree
//
//	@Description: 根据id查询单个学位信息 参数: IdRequest
//	@param c
func GetDegree(c *gin.Context) {
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
	degree, err2 := infoService().GetDegree(id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(degree))
	}
}
