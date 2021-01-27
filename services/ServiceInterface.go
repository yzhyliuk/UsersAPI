package services

import (
	"ms/usersAPI/utils/errors"

	"github.com/gin-gonic/gin"
)

//Service : provides an interface for serving requests
type Service interface {
	Create(c *gin.Context)
	Recive(c *gin.Context)
	List(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	configDataSource(dataSource, ConnectionString string) *errors.APIError
}
