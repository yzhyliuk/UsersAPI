// Package classification of Users API.
//
// This application is microservice for LightERP that provides basic CRUD actions with
// users of system.
//
// Datasource for this app is PostgresSQL database. All communication with database
// occurs through GORM ORM.
// Data Access Object implemented in dao pacakge as struct with methods that are
// executes db transactions. Dao avaliable via DAO interface
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /users
//     Version: 0.1.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key: in_development
// swagger:meta
package main

import (
	"fmt"
	"log"
	"ms/usersAPI/application"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	defaultLogger := log.New(os.Stdout, "users-api: ", log.LstdFlags)
	defaultRouter := gin.Default()
	defaultAppServer := http.Server{
		Addr:     ":9090",
		ErrorLog: defaultLogger,
		Handler:  defaultRouter,
	}
	config := application.Config{
		Server:         &defaultAppServer,
		Router:         defaultRouter,
		DataDriver:     "postgres",
		DataSourceName: "postgres://postgres:080919@/UsersAPI?sslmode=disable",
	}
	app, err := application.NewApp(config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	app.Start()
}

// @title Lighterp Users API
// @version 0.1.0
// @description This application is microservice for LightERP that provides basic CRUD actions with users of system.

// @host localhost:8080
// @BasePath /users
