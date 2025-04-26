package application

import (
	"context"

	"github.com/Irurnnen/gin-template/internal/config"
	"github.com/Irurnnen/gin-template/internal/repository"
	"github.com/Irurnnen/gin-template/internal/server"
	"go.uber.org/zap"
)

type Application struct {
	Config     *config.Config
	Logger     *zap.Logger
	Repository *repository.Repository
	Server     *server.Server
}

type ApplicationInterface interface {
	Run() error
	Shutdown() error
}

func New(config *config.Config, logger *zap.Logger, repository *repository.Repository, server *server.Server) *Application {
	return &Application{
		Config:     config,
		Logger:     logger,
		Repository: repository,
		Server:     server,
	}
}

func (a *Application) Run() error {
	// Run a server
	return a.Server.Start()
}

func (a *Application) Shutdown(ctx context.Context) error {
	a.Server.Shutdown(ctx)
	a.Repository.Close()
	return nil
}
