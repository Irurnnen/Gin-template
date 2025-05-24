package domains

import (
	"github.com/exceptionteapots/gin-template/repositories"
	"github.com/rs/zerolog"
)

type HelloDomain struct {
	repo   repositories.HelloRepositoryInterface
	logger *zerolog.Logger
}

type HelloDomainInterface interface {
	GetHelloMessage() (string, error)
}

func NewHelloDomain(repo repositories.HelloRepositoryInterface, logger *zerolog.Logger) *HelloDomain {
	return &HelloDomain{
		repo:   repo,
		logger: logger,
	}
}

func (s *HelloDomain) GetHelloMessage() (string, error) {
	message, err := s.repo.GetHelloMessage()
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to get hello message from repository")
		return "", err
	}
	return message, nil
}
