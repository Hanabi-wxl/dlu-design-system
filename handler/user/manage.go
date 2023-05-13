package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/utils"
	service "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(c *gin.Context) {
	var saveUserReq types.AddUserRequest
	if err := c.ShouldBindJSON(&saveUserReq); err != nil {
		msg := utils.TranslateOverride(err)
		c.JSON(http.StatusBadRequest, result.NewFailedResult(msg))
		return
	}
	userService := service.GetUserService()
	errno := userService().SaveUser(&saveUserReq)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

func ResetPassword(c *gin.Context) {
	var idReq IdRoleRequest
	if err := c.ShouldBindUri(&idReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userService := service.GetUserService()
	errno := userService().ResetPassword(idReq.Id, idReq.IsStu)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

func UpdateTeacherRole(c *gin.Context) {
	var req UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userService := service.GetUserService()
	errno := userService().UpdateTeacherRole(req.Id, req.RoleId)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

func UpdateUser(c *gin.Context) {
	var updateUserReq types.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userService := service.GetUserService()
	errno := userService().UpdateUser(&updateUserReq)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

func DeleteUser(c *gin.Context) {
	var idRoleReq IdRoleRequest
	if err := c.ShouldBindUri(&idRoleReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	userService := service.GetUserService()
	errno := userService().DeleteUser(idRoleReq.Id, idRoleReq.IsStu)
	if errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}
