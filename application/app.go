package application

import (
	"context"
	"ms/usersAPI/services"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type defaultApp struct {
	Server  *http.Server
	Router  *gin.Engine
	Service services.Service
}

//NewApp : returns an new App interface with default settings
func NewApp(conf Config) (App, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}
	//creating new webApp with config
	app, err := createApp(conf)
	if err != nil {
		return nil, err
	}
	return app, nil
}

//Start : start function for server
func (app *defaultApp) Start() {
	go func() {
		app.Server.ErrorLog.Printf("Starting server on port %s \n", app.Server.Addr)
		err := app.Server.ListenAndServe()
		if err != nil {
			app.Server.ErrorLog.Printf("Error starting server: %s", err.Error())
			os.Exit(1)
		}
	}()
	//Create channel to communicate with os catching terminate commands
	signChan := make(chan os.Signal)
	signal.Notify(signChan, os.Kill)
	signal.Notify(signChan, os.Interrupt)

	//Block next part of code via wating response from channel
	_ = <-signChan
	//if recieving one - log command and Gracefully shutDown the Server with Timeout of 30 sec
	app.Server.ErrorLog.Printf("Recived terminate command, graceful shutdown in 30 seconds")

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	app.Server.Shutdown(tc)

}
