package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	log.Info("Logger setup successfully")

	// Setup database provider
	dbProvider, err := database.NewProvider(cfg.Database.GetDSN())
	if err != nil {
		log.Fatal("Failed to initialize database provider", zap.Error(err))
	}
	defer dbProvider.Close()
	log.Info("Database provider setup successfully")

	// Ping database
	if err := dbProvider.Ping(); err != nil {
		log.Fatal("Failed to ping database", zap.Error(err), zap.String("host", cfg.Database.Host))
	}
	log.Info("Database connection ping successfully")

	// Initialize repository with database provider
	repo := repository.NewRepository(dbProvider, log)
	log.Debug("Repository created successful")

	// Setup server
	srv := server.NewServer(cfg, log, repo)
	log.Debug("Server created successfully")

	// Create application
	app := application.New(cfg, log, repo, srv)
	log.Debug("Application created successfully")

	// Launch application
	go func() {
		if err := app.Run(); err != nil {
			log.Fatal("Application failed to run", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Error("Server shutdown", zap.Error(err))
	}

	<-ctx.Done()

	log.Warn("Timeout of 5 seconds")
	log.Info("Server exiting")
}
