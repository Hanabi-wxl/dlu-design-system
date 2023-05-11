package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	service "github.com/Hanabi-wxl/dlu-design-system/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserById(c *gin.Context) {
	var IdRoleRequest IdRoleRequest
	if err := c.ShouldBindUri(&IdRoleRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	userService := service.GetUserService()
	user, err := userService().GetUserById(int(IdRoleRequest.Id), int(IdRoleRequest.IsStu))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(user))
	}
}

func GetUserByNumber(c *gin.Context) {
	var NumberRequest NumberRequest
	if err := c.ShouldBindUri(&NumberRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	userService := service.GetUserService()
	user, err := userService().GetUserByNumber(NumberRequest.Number, NumberRequest.IsStu)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(user))
	}
}

func GetUsersByMajor(c *gin.Context) {
	var MajorIdRequest MajorIdRequest
	if err := c.ShouldBindUri(&MajorIdRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	userService := service.GetUserService()
	users, err := userService().GetUsersByMajor(MajorIdRequest.MajorId, int(MajorIdRequest.IsStu),
		MajorIdRequest.Size, MajorIdRequest.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(users))
	}
}

func GetUsersByCollege(c *gin.Context) {
	var CollegeIdRequest CollegeIdRequest
	if err := c.ShouldBindUri(&CollegeIdRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	userService := service.GetUserService()
	users, err := userService().GetUsersByCollege(CollegeIdRequest.CollegeId, int(CollegeIdRequest.IsStu),
		CollegeIdRequest.Size, CollegeIdRequest.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(users))
	}
}
