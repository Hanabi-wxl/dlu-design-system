package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoleList(c *gin.Context) {
	userService := service.GetUserService()
	if roleList, errno := userService().GetRoleList(); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
		return
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(roleList))
	}
}

func UpdateRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	userService := service.GetUserService()
	if errno := userService().UpdateRole(&role); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

func DeleteRole(c *gin.Context) {
	var roleIdReq RoleIdRequest
	if err := c.ShouldBindUri(&roleIdReq); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	userService := service.GetUserService()
	if errno := userService().DeleteRole(roleIdReq.RoleId); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

func AddRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	userService := service.GetUserService()
	if errno := userService().AddRole(&role); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}
