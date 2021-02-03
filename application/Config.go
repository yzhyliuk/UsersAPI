package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Config : contain fields with app configuration parameters
type Config struct {
	Server         *http.Server
	Router         *gin.Engine
	DataDriver     string
	DataSourceName string
}

//Validate : validate configuration
func (c *Config) Validate() error {
	if c.Server == nil {
		return fmt.Errorf("Configuration Error: server is nil")
	}
	if c.Router == nil {
		return fmt.Errorf("Configuration Error: router is nil")
	}
	if c.DataDriver == "" {
		return fmt.Errorf("Configuration Error: data driver is an empty string")
	}
	if c.DataSourceName == "" {
		return fmt.Errorf("Configuration Error: DataSource name is an empty string")
	}
	return nil
}
