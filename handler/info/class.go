package handler

import (
	"net/http"

	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
)

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

func AddClass(c *gin.Context) {

}

func DeleteClass(c *gin.Context) {

}

func UpdateClass(c *gin.Context) {

}

func GetClass(c *gin.Context) {

}
