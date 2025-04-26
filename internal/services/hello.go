package services

import "github.com/Irurnnen/gin-template/internal/repository"

type HelloService struct {
	repo repository.HelloRepositoryInterface
}

type HelloServiceInterface interface {
	GetHelloMessage() (string, error)
}

func NewHelloService(repo repository.HelloRepositoryInterface) *HelloService {
	return &HelloService{
		repo: repo,
	}
}

func (s *HelloService) GetHelloMessage() (string, error) {
	message, err := s.repo.GetHelloMessage()
	return message, err
}
