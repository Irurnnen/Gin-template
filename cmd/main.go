package main

import (
	"github.com/Irurnnen/gin-template/internal/application"
	"github.com/Irurnnen/gin-template/internal/config"
	"github.com/Irurnnen/gin-template/internal/logger"
	"github.com/Irurnnen/gin-template/internal/repository"
	"github.com/Irurnnen/gin-template/internal/server"
	"go.uber.org/zap"
)

func main() {
	// Read config
	cfg := config.NewConfig()

	// Setup logger
	log := logger.New(cfg.LogLevel)

	// Setup database connection
	repo, err := repository.NewRepository(cfg.Database.GetDSN())
	if err != nil {
		log.Fatal("Failed to initialize repository", zap.Error(err))
	}

	// Ping database
	if err := repo.Ping(); err != nil {
		log.Fatal("Failed to ping database", zap.Error(err), zap.String("host", cfg.Database.Host))
	}
	log.Info("Database connection setup successfully")

	// Setup server
	srv := server.NewServer(cfg, log, repo)

	// Create application
	app := application.New(cfg, log, repo, srv)

	// Launch application
	app.Run()
}
