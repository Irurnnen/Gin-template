package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/exceptionteapots/gin-template/internal/config"
	"github.com/exceptionteapots/gin-template/internal/handler"
	"github.com/exceptionteapots/gin-template/internal/logger"
	"github.com/exceptionteapots/gin-template/internal/repository"
	"github.com/exceptionteapots/gin-template/internal/server"
	"github.com/exceptionteapots/gin-template/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	@title			Gin-template
//	@version		0.0.1
//	@description	This is a sample server celler server.

// @host		localhost:8080
// @BasePath	/v1
func main() {
	// Read config
	cfg := config.NewConfig()

	// Setup logger
	log := logger.New(cfg.LogLevel)
	log.Info("Logger setup successfully")

	// Initialize database
	repo, err := repository.NewRepository(cfg.DatabaseConfig.GetDSN(), log)
	if err != nil {
		log.Fatal("Failed to initialize database", zap.Error(err))
	}
	log.Info("Database setup successfully")
	defer repo.Close()

	// Ping database
	if err := repo.Ping(); err != nil {
		log.Fatal("Failed to ping database", zap.Error(err), zap.String("host", cfg.DatabaseConfig.Host))
	}
	log.Info("Database connection ping successfully")

	// Initialize Hello handler
	HelloService := services.NewHelloService(repo.HelloRepository, log)
	HelloHandler := handler.NewHelloHandler(HelloService, log)

	// Setup gin level
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup server
	srv := server.NewServer(cfg.ServerConfig, log, HelloHandler)
	log.Debug("Server created successfully")

	// Launch application
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatal("Application failed to run", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server shutdown", zap.Error(err))
	}

	<-ctx.Done()

	log.Warn("Timeout of 5 seconds")
	log.Info("Server exiting")
}
