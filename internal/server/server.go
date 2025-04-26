package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Irurnnen/gin-template/internal/config"
	"github.com/Irurnnen/gin-template/internal/handler"
	"github.com/Irurnnen/gin-template/internal/repository"
	"github.com/Irurnnen/gin-template/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ServerInterface interface {
	Start() error
	Shutdown() error
}

type Server struct {
	config *config.Config
	logger *zap.Logger
	server *http.Server
	router *gin.Engine
	// repo   *repository.Repository
}

func NewServer(config *config.Config, logger *zap.Logger, repo *repository.Repository) *Server {
	// Create a new Gin router instance
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Setup hello
	HelloService := services.NewHelloService(repo.HelloRepository, logger)
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
		// router: router,
		server: &http.Server{
			Addr:    ":" + strconv.Itoa(config.Server.Port),
			Handler: router.Handler(),
		},
	}
}

func (s *Server) Start() error {
	s.logger.Info("Starting server", zap.String("port", strconv.Itoa(s.config.Server.Port)))

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error("Server failed to start", zap.Error(err))
		return err
	}
	s.logger.Info("Server started successfully")
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
