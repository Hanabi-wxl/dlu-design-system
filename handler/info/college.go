package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageRequest struct {
	Size int
	Num  int
}

type IdRequest struct {
	id int
}

func GetCollegeList(c *gin.Context) {
	var page PageRequest
	if err := c.ShouldBindUri(&page); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	infoService := service.GetInfoService()
	list, err := infoService().GetCollegeList(page.Size, page.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}

func AddCollege(c *gin.Context) {

}

func DeleteCollege(c *gin.Context) {

}

func UpdateCollege(c *gin.Context) {

}

func GetCollege(c *gin.Context) {

}
