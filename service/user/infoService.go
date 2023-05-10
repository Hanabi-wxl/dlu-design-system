package service

import "github.com/gin-gonic/gin"

type InfoService interface {
	UserInfo()
	GetUserById(c *gin.Context)
	GetUserByNumber(c *gin.Context)
	GetUsersByMajor(c *gin.Context)
	GetUsersByCollege(c *gin.Context)
}
