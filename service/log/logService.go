package service

import "github.com/gin-gonic/gin"

type LogService interface {
	AddLog(content string, c *gin.Context, state int8, method int8)
}
