package main

import (
	api "image-receiver-api/api/app"
	"image-receiver-api/config"
)

var settings = config.GetSettings()

func main() {
	app := api.MakeApp()

	portListen := ":" + settings["PORT"]
	err := app.Listen(portListen)
	if err != nil {
		panic(err)
	}
}
