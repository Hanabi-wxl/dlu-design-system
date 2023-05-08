package handler

import (
	service "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/gin-gonic/gin"
)

func Role(c *gin.Context) {
	userService := service.GetUserService()
	userService.CheckRole()
}
