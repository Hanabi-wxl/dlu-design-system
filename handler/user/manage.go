package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveUser(c *gin.Context) {
	var saveUserReq service.UserRequest
	if err := c.ShouldBindJSON(&saveUserReq); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	userService := service.GetUserService()
	errno := userService().SaveUser(saveUserReq)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}
