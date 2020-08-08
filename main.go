//go:generate pkger -o bindata -include /client
package main

import (
	"OSRailwayControl/app"
	"OSRailwayControl/mqtt"
	"OSRailwayControl/web"
)

func main() {
	osRailway := app.App{
		Config: app.ParseFlags(),
	}
	osRailway.Webserver = web.NewWeb(&osRailway)
	osRailway.Mqtt = mqtt.NewMQTT(&osRailway)

	osRailway.Run()
}
