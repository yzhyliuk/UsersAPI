package application

import (
	"fmt"
	"ms/usersAPI/services/commonservice"
)

// Creates new app with given configuration
func createApp(conf Config) (*defaultApp, error) {
	//Connecting to db
	userservice, err := commonservice.NewUserService(conf.DataDriver, conf.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf(err.Message)
	}
	app := &defaultApp{
		Service: userservice,
		Server:  conf.Server,
		Router:  conf.Router,
	}
	app.mapRoutes()
	return app, nil
}
