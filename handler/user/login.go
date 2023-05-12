package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/utils"
	service "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckLoginRole(c *gin.Context) {
	number := c.Param("number")
	userService := service.GetUserService()
	role, errno := userService().CheckLoginRole(number)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(role))
	}
}

func StudentLogin(c *gin.Context) {
	var userLoginReq UserLoginRequest
	if err := c.ShouldBindJSON(&userLoginReq); err != nil {
		msg := utils.TranslateOverride(err)
		c.JSON(http.StatusBadRequest, result.NewFailedResult(msg))
	}
	userService := service.GetUserService()
	res, errno := userService().StudentLogin(userLoginReq.Number, userLoginReq.Password)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(res))
	}
}

func TeacherLogin(c *gin.Context) {
	var userLoginReq UserLoginRequest
	if err := c.ShouldBindJSON(&userLoginReq); err != nil {
		msg := utils.TranslateOverride(err)
		c.JSON(http.StatusBadRequest, result.NewFailedResult(msg))
	}
	userService := service.GetUserService()
	res, errno := userService().TeacherLogin(userLoginReq.Number, userLoginReq.Password)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(res))
	}
}
