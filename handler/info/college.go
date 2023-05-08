package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetCollegeListReq struct {
	Size int
	Num  int
}

func GetCollegeList(c *gin.Context) {
	var getCollegeListReq GetCollegeListReq
	if err := c.ShouldBindUri(&getCollegeListReq); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
	}
	infoService := service.GetInfoService()
	list, err := infoService.GetCollegeList(getCollegeListReq.Size, getCollegeListReq.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}
