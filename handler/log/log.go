package handler

import (
	"github.com/Hanabi-wxl/dlu-design-system/pkg/result"
	"github.com/Hanabi-wxl/dlu-design-system/pkg/types"
	service "github.com/Hanabi-wxl/dlu-design-system/service/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLogList(c *gin.Context) {
	var page types.PageRequest
	if err := c.ShouldBindUri(&page); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if list, errno := service.GetLogService().LogList(page.Size, page.Num); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}

func GetLogListByDate(c *gin.Context) {
	var req types.LogDateReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if list, errno := service.GetLogService().LogDateList(&req); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}

func GetLogListByMethod(c *gin.Context) {
	var req types.LogMethodReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if list, errno := service.GetLogService().LogMethodList(&req); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}

func GetLogListByState(c *gin.Context) {
	var req types.LogStateReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if list, errno := service.GetLogService().LogStateList(&req); errno != nil {
		c.JSON(http.StatusBadRequest, errno)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(list))
	}
}
