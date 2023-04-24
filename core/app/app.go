package app

import (
	"shopy/config"
	"shopy/logger"

	"github.com/gin-gonic/gin"
)

type App struct {
	ApplicationName string
	Router          *gin.Engine
}

func New(applicationName string) *App {
	return &App{
		Router:          gin.Default(),
		ApplicationName: applicationName,
	}
}

func (a *App) Init() {
	logger.InitLogger()
	config.InitConfig()
}

func (a *App) Start() {
	logger.Info("App is started")
	a.Router.Run()
}
