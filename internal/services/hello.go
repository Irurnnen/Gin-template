package services

import (
	"context"

	"github.com/Irurnnen/gin-template/internal/repository"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type HelloService struct {
	repo   repository.HelloRepositoryInterface
	logger *zap.Logger
}

type HelloServiceInterface interface {
	GetHelloMessage(ctx context.Context) (string, error)
}

func NewHelloService(repo repository.HelloRepositoryInterface, logger *zap.Logger) *HelloService {
	return &HelloService{
		repo:   repo,
		logger: logger,
	}
}

func (s *HelloService) GetHelloMessage(ctx context.Context) (string, error) {
	ctx, span := otel.Tracer("service").Start(ctx, "GetHelloMessage")
	defer span.End()

	message, err := s.repo.GetHelloMessage(ctx) // Передаем контекст с трассировкой
	if err != nil {
		return "", err
	}
	return message, nil
}
