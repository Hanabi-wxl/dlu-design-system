package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/dal/model"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	logService "github.com/Hanabi-wxl/dlu-design-system/service/log"
	userService "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoleList(c *gin.Context) {
	//usersvs := userService.GetUserService()
	if roleList, errno := userService.GetUserService().GetRoleList(); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
		return
	} else {
		logService.GetLogService().AddLog("获取角色列表", c.Request.Context(), c.ClientIP(), 1, 1)
		c.JSON(http.StatusOK, result.NewOkResult(roleList))
	}
}

func GetRole(c *gin.Context) {
	var roleIdReq RoleIdRequest
	if err := c.ShouldBindUri(&roleIdReq); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	//usersvs := userService.GetUserService()
	if role, errno := userService.GetUserService().GetRole(roleIdReq.RoleId); errno != nil {
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
	//usersvs := userService.GetUserService()
	if errno := userService.GetUserService().UpdateRole(&role); errno != nil {
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
	//usersvs := userService.GetUserService()
	if errno := userService.GetUserService().DeleteRole(roleIdReq.RoleId); errno != nil {
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
	//usersvs := userService.GetUserService()
	if errno := userService.GetUserService().AddRole(&role); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}
