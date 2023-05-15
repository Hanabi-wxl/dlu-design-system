package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOffices
//
//	@Description: 查询所有科室信息 参数:
//	@param c
func GetOffices(c *gin.Context) {
	Offices, err2 := service.GetInfoService().GetOffices()
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(Offices))
	}
}

// GetOffice
//
//	@Description: 根据id查询单个科室信息 参数: IdRequest
//	@param c
func GetOffice(c *gin.Context) {
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	Office, err2 := service.GetInfoService().GetOffice(idRequest.Id)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(Office))
	}
}
