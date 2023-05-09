package handler

import (
	"net/http"

	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
)

func GetSchoolList(c *gin.Context) {
	var page PageRequest
	if err := c.ShouldBindUri(&page); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	infoService := service.GetInfoService()
	list, err := infoService().GetSchoolList(page.Size, page.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}

func AddSchool(c *gin.Context) {

}

func DeleteSchool(c *gin.Context) {

}

func UpdateSchool(c *gin.Context) {

}

func GetSchool(c *gin.Context) {

}
