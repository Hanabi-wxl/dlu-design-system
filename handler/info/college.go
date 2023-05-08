package handler

import (
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
)

func GetCollegeList(c *gin.Context) {
	infoService := service.GetInfoService()
	list, err := infoService.GetCollegeList(1, 1)
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, list)
	}
}
