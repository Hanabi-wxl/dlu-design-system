package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/middleware/jwt"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetUserById(c *gin.Context) {
	var idRoleRequest IdRoleRequest
	if err := c.ShouldBindUri(&idRoleRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	//userService := service.GetUserService()
	user, err := service.GetUserService().GetUserById(idRoleRequest.Id, idRoleRequest.IsStu)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(user))
	}
}

func GetUserByNumber(c *gin.Context) {
	var numberRequest NumberRequest
	if err := c.ShouldBindUri(&numberRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	//userService := service.GetUserService()
	user, err := service.GetUserService().GetUserByNumber(numberRequest.Number, numberRequest.IsStu)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(user))
	}
}

func GetUsersByMajor(c *gin.Context) {
	var majorIdRequest MajorIdRequest
	if err := c.ShouldBindUri(&majorIdRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	//userService := service.GetUserService()
	users, err := service.GetUserService().GetUsersByMajor(majorIdRequest.MajorId, majorIdRequest.IsStu,
		majorIdRequest.Size, majorIdRequest.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(users))
	}
}

func GetUsersByCollege(c *gin.Context) {
	userInfo := jwt.GetUserInfo(c.Request.Context())
	logrus.Info(userInfo)
	var collegeIdRequest CollegeIdRequest
	if err := c.ShouldBindUri(&collegeIdRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	//userService := service.GetUserService()
	users, err := service.GetUserService().GetUsersByCollege(collegeIdRequest.CollegeId, collegeIdRequest.IsStu,
		collegeIdRequest.Size, collegeIdRequest.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(users))
	}
}
