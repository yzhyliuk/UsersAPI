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

type webApp struct {
	webServer   *http.Server
	webRouter   *gin.Engine
	UserService services.Service
}

//NewApp : returns an new App interface with default settings
func NewApp(conf Config) (App, error) {
	//creating new webApp with config
	app, err := createApp(conf)
	if err != nil {
		return nil, err
	}
	return app, nil
}

//Start : start function for server
func (w *webApp) Start() {
	go func() {
		w.webServer.ErrorLog.Printf("Starting server on port %s \n", w.webServer.Addr)
		err := w.webServer.ListenAndServe()
		if err != nil {
			w.webServer.ErrorLog.Printf("Error starting server: %s", err.Error())
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
	w.webServer.ErrorLog.Printf("Recived terminate command, graceful shutdown in 30 seconds")

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	w.webServer.Shutdown(tc)

}

func (w *webApp) applyConfiguration(conf Config) {
	w.webServer = conf.Server
	w.webServer.Handler = conf.Router
	w.webRouter = conf.Router
}
