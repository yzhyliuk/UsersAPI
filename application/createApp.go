package application

import (
	"fmt"
	"ms/usersAPI/services"
)

// Creates new app with given configuration
func createApp(conf Config) (*webApp, error) {
	//Connecting to db
	userservice, err := services.NewUserService(conf.DataSource, conf.DataConnectionString)
	if err != nil {
		return nil, fmt.Errorf(err.Message)
	}
	app := &webApp{
		UserService: userservice,
	}
	app.applyConfiguration(conf)
	app.mapRoutes()
	return app, nil
}
