package application

import (
	"github.com/Irurnnen/gin-template/internal/config"
	"github.com/Irurnnen/gin-template/internal/logger"
	"github.com/Irurnnen/gin-template/internal/server"
)

type Application struct {
	Config config.Config
	Debug  bool
}

func New() *Application {
	return &Application{
		Config: *config.NewConfig(),
		Debug:  false,
	}
}

func NewDebug() *Application {
	return &Application{
		Config: *config.NewConfig(),
		Debug:  true,
	}
}

func (a *Application) Run() error {
	// TODO: Init logger
	logger := logger.New(a.Config.LogLevel, a.Debug)

	// TODO: Init telemetry

	// TODO: Init database
	// TODO: Init cache
	// TOTO: Init server
	server := server.NewServer(&a.Config, logger, nil)

	return server.Start()
}
