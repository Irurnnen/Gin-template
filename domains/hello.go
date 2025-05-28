package domains

import (
	"context"

	"github.com/exceptionteapots/gin-template/repositories"
	"github.com/rs/zerolog"
)

type HelloDomain struct {
	repo   repositories.HelloRepositoryInterface
	logger *zerolog.Logger
}

type HelloDomainInterface interface {
	GetHelloMessage(context.Context) (*HelloEntity, error)
	GetHelloMessageWithCache(context.Context) (*HelloEntity, error)
}

func NewHelloDomain(repo repositories.HelloRepositoryInterface, logger *zerolog.Logger) *HelloDomain {
	return &HelloDomain{
		repo:   repo,
		logger: logger,
	}
}

func (s *HelloDomain) GetHelloMessage(ctx context.Context) (*HelloEntity, error) {
	entity, err := s.repo.GetHelloMessage(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to get hello message from repository")
		return nil, err
	}
	return &HelloEntity{
		Message: entity.Message,
	}, nil
}

func (s *HelloDomain) GetHelloMessageWithCache(ctx context.Context) (*HelloEntity, error) {
	entity, err := s.repo.GetHelloMessageWithCache(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to get hello message from repository")
		return nil, err
	}
	return &HelloEntity{
		Message: entity.Message,
	}, nil
}
