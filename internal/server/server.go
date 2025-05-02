package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Irurnnen/gin-template/internal/config"
	"github.com/Irurnnen/gin-template/internal/handler"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

type ServerInterface interface {
	Start() error
	Shutdown() error
}

type Server struct {
	logger     *zap.Logger
	httpServer *http.Server
}

func NewServer(cfg *config.ServerConfig, logger *zap.Logger, helloHandler handler.HelloHandlerInterface) *Server {

	// Create a new Gin router instance
	router := gin.New()

	// Add middlewares
	router.Use(
		otelgin.Middleware("gin-template"), // Указываем имя приложения
		gin.Logger(),
		gin.Recovery(),
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

		v1.GET("/hello", helloHandler.GetHelloMessage)
	}

	logger.Debug("Router initialized")

	return &Server{
		logger: logger,
		httpServer: &http.Server{
			Addr:    cfg.Host + ":" + strconv.Itoa(cfg.Port),
			Handler: router.Handler(),
		},
	}
}

func (s *Server) Start() error {
	s.logger.Info("Starting server", zap.String("address", s.httpServer.Addr))
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error("Server failed to start", zap.Error(err))
		return err
	}
	s.logger.Info("Server started successfully")
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down server...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("Server failed to shutdown", zap.Error(err))
		return err
	}
	return nil
}
