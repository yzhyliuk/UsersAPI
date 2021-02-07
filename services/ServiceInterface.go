package services

import (
	"github.com/gin-gonic/gin"
)

//Service : provides an interface for serving requests
type Service interface {
	Create(c *gin.Context)
	Recive(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	UpdateWhere(c *gin.Context)
	Delete(c *gin.Context)
	DeleteWhere(c *gin.Context)
	Login(c *gin.Context)
}
