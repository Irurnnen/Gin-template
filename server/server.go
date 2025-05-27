package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/exceptionteapots/gin-template/config"
	"github.com/exceptionteapots/gin-template/controllers"
	"github.com/exceptionteapots/gin-template/server/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	ginSwaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ServerInterface defines interface of server of this project
type ServerInterface interface {
	Start() error
	Shutdown() error
}

// Server responsible for up HTTP server
type Server struct {
	logger     *zerolog.Logger
	httpServer *http.Server
}

// NewServer creates new entity of Server
func NewServer(cfg *config.ServerConfig, logger *zerolog.Logger, helloController controllers.HelloControllerInterface, debug bool) *Server {
	// Create a new Gin router instance
	router := gin.New()

	// Add middlewares
	router.Use(
		gin.Logger(),
		gin.Recovery(),
		// Example middleware
		middlewares.CorrelationIDMiddleware(),
	)

	// If debug is enabled add swagger endpoint
	if debug {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(ginSwaggerFiles.Handler))
	}

	// Setup routes
	v1 := router.Group("/v1")
	{
		internal := v1.Group("/internal")
		{
			internal.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "pong"})
			})
		}

		v1.GET("/hello", helloController.GetHelloMessage)
	}

	logger.Debug().Msg("Routes initialized")

	return &Server{
		logger: logger,
		httpServer: &http.Server{
			Addr:    cfg.Host + ":" + strconv.Itoa(cfg.Port),
			Handler: router.Handler(),
		},
	}
}

func (s *Server) Start() error {
	s.logger.Info().Str("address", s.httpServer.Addr).Msg("Starting server...")
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error().Err(err).Msg("Server failed to start")
		return err
	}
	s.logger.Info().Msg("Server started successfully")
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info().Msg("Shutting down server...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error().Err(err).Msg("Server failed to shutdown")
		return err
	}
	return nil
}
