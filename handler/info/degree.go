package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetDegrees
//
//	@Description: 查询全部学位信息 参数:
//	@param c
func GetDegrees(c *gin.Context) {
	degrees, err2 := service.GetInfoService().GetDegrees()
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
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	degree, err2 := service.GetInfoService().GetDegree(idRequest.Id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(degree))
	}
}
