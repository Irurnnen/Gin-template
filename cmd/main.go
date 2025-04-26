package main

import (
	"github.com/Irurnnen/gin-template/internal/application"
	"github.com/Irurnnen/gin-template/internal/config"
	"github.com/Irurnnen/gin-template/internal/database"
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

	// Setup database provider
	dbProvider, err := database.NewProvider(cfg.Database.GetDSN())
	if err != nil {
		log.Fatal("Failed to initialize database provider", zap.Error(err))
	}
	defer dbProvider.Close()

	// Ping database
	if err := dbProvider.Ping(); err != nil {
		log.Fatal("Failed to ping database", zap.Error(err), zap.String("host", cfg.Database.Host))
	}
	log.Info("Database connection setup successfully")

	// Initialize repository with database provider
	repo := repository.NewRepository(dbProvider)

	// Setup server
	srv := server.NewServer(cfg, log, repo)

	// Create application
	app := application.New(cfg, log, repo, srv)

	// Launch application
	if err := app.Run(); err != nil {
		log.Fatal("Application failed to run", zap.Error(err))
	}
}
