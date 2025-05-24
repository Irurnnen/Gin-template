package services

import (
	"github.com/exceptionteapots/gin-template/internal/repository"
	"github.com/rs/zerolog"
)

type HelloService struct {
	repo   repository.HelloRepositoryInterface
	logger *zerolog.Logger
}

type HelloServiceInterface interface {
	GetHelloMessage() (string, error)
}

func NewHelloService(repo repository.HelloRepositoryInterface, logger *zerolog.Logger) *HelloService {
	return &HelloService{
		repo:   repo,
		logger: logger,
	}
}

func (s *HelloService) GetHelloMessage() (string, error) {
	message, err := s.repo.GetHelloMessage()
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to get hello message from repository")
		return "", err
	}
	return message, nil
}
