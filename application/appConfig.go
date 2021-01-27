package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Config : contain fields with app configuration parameters
type Config struct {
	Server               *http.Server
	Router               *gin.Engine
	DataSource           string
	DataConnectionString string
}
