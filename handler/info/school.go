package handler

import (
	service "github.com/Hanabi-wxl/dlu-design-system/service/info"
	"github.com/gin-gonic/gin"
)

func InfoSchools(c *gin.Context) {
	infoService := service.GetInfoService()
	schools, err := infoService.GetAllSchool()
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, schools)
	}
}
