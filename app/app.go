package app

import (
	"OSRailwayControl/handler"
	"OSRailwayControl/store"
	"fmt"
)

type App struct {
	Config    *Config
	Webserver handler.WebHandler
	Mqtt      handler.MQTTHandler

	TrainStore store.TrainStore
}

func (a *App) Run() {
	a.TrainStore = store.NewTrainStore()

	errCh := make(chan error)
	go func() {
		err := a.Webserver.Listen()
		if err != nil {
			errCh <- err
		}
		close(errCh)
	}()

	go func() {
		err := a.Mqtt.Listen()
		if err != nil {
			errCh <- err
		}
		close(errCh)
	}()

	err := <-errCh
	if err != nil {
		fmt.Print(err)
	}
}
