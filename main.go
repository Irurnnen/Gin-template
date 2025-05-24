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
	"github.com/jackc/pgx/v5/pgxpool"
)

//	@title			Gin-template
//	@version		0.0.1
//	@description	This is a sample server caller server.

// @host		localhost:8080
// @BasePath	/v1
func main() {
	// Read config
	cfg := config.NewConfig()

	// Setup logger
	log := logger.New(cfg.LogLevel)
	log.Info().Msg("Logger setup successfully")

	// Initialize database
	dbPool, err := pgxpool.New(context.Background(), cfg.DatabaseConfig.GetDSN())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database connection pool")
	}
	log.Info().Msg("Database connection pool setup successfully")

	// Initialize repository
	helloRepoLogger := log.With().Str("repository", "hello").Logger()
	helloRepo := repository.NewHelloRepository(dbPool, &helloRepoLogger)
	log.Debug().Msg("Hello repository created successfully")

	// Ping database
	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatal().Str("host", cfg.DatabaseConfig.Host).Err(err).Msg("Failed to ping database")
	}
	log.Info().Msg("Database connection ping successfully")

	// Initialize Hello handler
	HelloService := services.NewHelloService(helloRepo, log)
	HelloHandler := handler.NewHelloHandler(HelloService, log)

	// Setup server
	srv := server.NewServer(cfg.ServerConfig, log, HelloHandler)
	log.Debug().Msg("Server created successfully")

	// Launch application
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatal().Err(err).Msg("Application failed to run")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server shutdown")
	}

	<-ctx.Done()

	log.Warn().Msg("Timeout of 5 seconds")
	log.Info().Msg("Server exiting")
}
