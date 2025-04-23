package server

import (
	"fmt"

	"github.com/Irurnnen/gin-template/internal/config"
	"github.com/Irurnnen/gin-template/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	config *config.Config
	logger *zap.Logger
	router *gin.Engine
	repo   *repository.Repository
}

func NewServer(config *config.Config, logger *zap.Logger, repo *repository.Repository) *Server {

	// Create a new Gin router instance
	router := NewRouter()

	return &Server{
		config: config,
		logger: logger,
		router: router,
		repo:   repo,
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
