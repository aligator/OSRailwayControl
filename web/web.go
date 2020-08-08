// package web is for serving the html, css and js as website
package web

import (
	"OSRailwayControl/app"
	"OSRailwayControl/handler"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type web struct {
	app    *app.App
	port   int
	router *mux.Router
	socket handler.Socket
}

func NewWeb(app *app.App) handler.WebHandler {
	w := web{
		app:  app,
		port: app.Config.Web.Port,
	}
	return &w
}

func (w *web) Listen() error {
	w.socket = newSocket()
	w.setupSocketListeners()
	w.router = mux.NewRouter()
	w.setupRoutes()
	w.setupMiddlewares()

	srv := &http.Server{
		Handler:      w.router,
		Addr:         ":" + strconv.Itoa(w.port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listen on :" + strconv.Itoa(w.port))
	return srv.ListenAndServe()
}

func (w *web) Socket() handler.Socket {
	return w.socket
}
