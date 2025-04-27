package handler

import (
	"reflect"
	"testing"

	"github.com/Irurnnen/gin-template/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type MockHelloService struct {
	mock.Mock
}

func (m *MockHelloService) GetHelloMessage() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestNewHelloHandler(t *testing.T) {
	type args struct {
		service services.HelloServiceInterface
		logger  *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want *HelloHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHelloHandler(tt.args.service, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHelloHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelloHandler_GetHelloMessage(t *testing.T) {
	// router := gin.New()

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *HelloHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetHelloMessage(tt.args.c)
		})
	}
}
