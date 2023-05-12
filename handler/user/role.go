package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	userService "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoleList(c *gin.Context) {
	usersvs := userService.GetUserService()
	if roleList, errno := usersvs().GetRoleList(); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
		return
	} else {
		// new(logService.LogServiceImpl).AddLog("获取角色列表", c, 1, 1)
		c.JSON(http.StatusOK, result.NewOkResult(roleList))
	}
}

func GetRole(c *gin.Context) {
	var roleIdReq RoleIdRequest
	if err := c.ShouldBindUri(&roleIdReq); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	usersvs := userService.GetUserService()
	if role, errno := usersvs().GetRole(roleIdReq.RoleId); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(role))
	}
}

func UpdateRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	usersvs := userService.GetUserService()
	if errno := usersvs().UpdateRole(&role); errno != nil {
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
	usersvs := userService.GetUserService()
	if errno := usersvs().DeleteRole(roleIdReq.RoleId); errno != nil {
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
	usersvs := userService.GetUserService()
	if errno := usersvs().AddRole(&role); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}
