package server

import (
	"fmt"

	"github.com/Irurnnen/gin-template/internal/config"
	"github.com/Irurnnen/gin-template/internal/handler"
	"github.com/Irurnnen/gin-template/internal/repository"
	"github.com/Irurnnen/gin-template/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	config *config.Config
	logger *zap.Logger
	router *gin.Engine
	// repo   *repository.Repository
}

func NewServer(config *config.Config, logger *zap.Logger, repo *repository.Repository) *Server {
	// Create a new Gin router instance
	router := gin.New()

	// Setup hello
	HelloRepository := repository.NewHelloRepository(repo.DB)
	HelloService := services.NewHelloService(HelloRepository)
	HelloHandler := handler.NewHelloHandler(HelloService, logger)

	// Add middlewares
	router.Use(
		gin.Logger(),
		gin.Recovery(),
		// TODO: Add tracer
	)

	// Setup routes
	v1 := router.Group("/v1")
	{
		internal := v1.Group("/internal")
		{
			internal.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "pong"})
			})
		}

		v1.GET("/hello", HelloHandler.GetHelloMessage)
	}

	logger.Debug("Router initialized")

	return &Server{
		config: config,
		logger: logger,
		router: router,
		// repo:   repo,
	}
}

func (s *Server) Start() error {
	// Start the server on the configured port
	s.logger.Info("Starting server", zap.String("port", fmt.Sprint(s.config.Server.Port)))

	if err := s.router.Run(":" + fmt.Sprint(s.config.Server.Port)); err != nil {
		return err
	}
	return nil
}
