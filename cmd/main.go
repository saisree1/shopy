package main

import (
	"shopy/core/app"
	"shopy/handler"
)

func main() {
	appEngine := app.New("Shopy")
	appEngine.Init()
	handler.SetupAppRoutes(appEngine.Router)
	appEngine.Start()
}
