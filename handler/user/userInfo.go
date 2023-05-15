package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	service "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserById(c *gin.Context) {
	var idRoleRequest types.IdRoleRequest
	if err := c.ShouldBindUri(&idRoleRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	user, err := service.GetUserService().GetUserById(idRoleRequest.Id, idRoleRequest.IsStu)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(user))
	}
}

func GetUserByNumber(c *gin.Context) {
	var numberRequest types.NumberRequest
	if err := c.ShouldBindUri(&numberRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	user, err := service.GetUserService().GetUserByNumber(numberRequest.Number, numberRequest.IsStu)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(user))
	}
}

func GetUserByNumberMajor(c *gin.Context) {
	var numberMajorRequest types.NumberMajorRequest
	if err := c.ShouldBindUri(&numberMajorRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	user, err := service.GetUserService().GetUserByNumberMajor(&numberMajorRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(user))
	}
}

func GetUsersByMajor(c *gin.Context) {
	var majorIdRequest types.UserQueryByMajorReq
	if err := c.ShouldBindUri(&majorIdRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	users, err := service.GetUserService().GetUsersByMajor(&majorIdRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(users))
	}
}

func GetUsersByCollege(c *gin.Context) {
	var collegeIdRequest types.UserQueryByCollegeReq
	if err := c.ShouldBindUri(&collegeIdRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	users, err := service.GetUserService().GetUsersByCollege(&collegeIdRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(users))
	}
}
