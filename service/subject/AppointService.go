package Service

import "github.com/gin-gonic/gin"

type AppointService interface {
	GetAppointList(c *gin.Context)
	SetAppoint(c *gin.Context)
}
