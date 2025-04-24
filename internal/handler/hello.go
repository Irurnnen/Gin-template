package handler

import (
	"net/http"

	"github.com/Irurnnen/gin-template/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HelloHandler struct {
	service services.HelloServiceInterface
	logger  *zap.Logger
}

type HelloHandlerInterface interface {
	GetHelloMessage(c *gin.Context)
}

func NewHelloHandler(service services.HelloServiceInterface, logger *zap.Logger) *HelloHandler {
	return &HelloHandler{
		service: service,
		logger:  logger,
	}
}

func (h *HelloHandler) GetHelloMessage(c *gin.Context) {
	h.logger.Debug("Get hello message in handler")
	message, err := h.service.GetHelloMessage()
	if err != nil {
		h.logger.Error("Failed to get hello message", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}
