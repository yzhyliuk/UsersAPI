package main

import (
	"fmt"
	"log"
	"ms/usersAPI/application"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	user     string = "root"
	password string = "NewPassword"
	dbName   string = "usersapi"
)

//import "NixTwo/application"

func main() {
	defaultLogger := log.New(os.Stdout, "users-api: ", log.LstdFlags)
	defaultRouter := gin.Default()
	defaultAppServer := http.Server{
		Addr:     ":9090",
		ErrorLog: defaultLogger,
		Handler:  defaultRouter,
	}
	config := application.Config{
		Server: &defaultAppServer,
		Router: defaultRouter,
		// DataSource:           "mysql",
		// DataConnectionString: "root:NewPassword@/usersapi",
		DataSource:           "postgres",
		DataConnectionString: "postgres://postgres:080919@/UsersAPI?sslmode=disable",
	}
	app, err := application.NewApp(config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	app.Start()
}
