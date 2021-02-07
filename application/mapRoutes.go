package application

import (
	"ms/usersAPI/application/mdlwr"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *defaultApp) mapRoutes() {

	//LoadDocumentation - Not implemented
	app.Router.StaticFile("/docs/swagger.json", "./docs/swagger.json")
	url := ginSwagger.URL("http://localhost:9090/docs/swagger.json") // The url pointing to API definition
	app.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//Login endpoint
	app.Router.POST("/login", app.Service.Login)
	//Create
	app.Router.POST("/users", app.Service.Create)
	//Authentication required endpoints
	app.Router.Use(mdlwr.AuthenticationRequired)
	//Recive
	app.Router.GET("/users/:id", app.Service.Recive)
	app.Router.GET("/users", app.Service.FindAll)
	//Update
	app.Router.PUT("/users/:id", app.Service.Update)
	app.Router.PUT("/users", app.Service.UpdateWhere)
	//Delete
	app.Router.DELETE("/users/:id", app.Service.Delete)
	app.Router.DELETE("/users", app.Service.DeleteWhere)

}
