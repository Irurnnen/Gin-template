package handler

import (
	"net/http"

	"github.com/Irurnnen/gin-template/internal/models"
	"github.com/Irurnnen/gin-template/internal/services"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
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

// GetHelloMessage godoc
//
//	@Summary		Get Hello World message using database
//	@Description	get hello world
//	@Tags			Hello
//	@Produce		json
//	@Success		200	{object}	models.Message
//	@Failure		500	{object}	models.HTTPError
//	@Router			/hello [GET]
func (h *HelloHandler) GetHelloMessage(c *gin.Context) {
	ctx, span := otel.Tracer("handler").Start(c.Request.Context(), "GetHelloMessage")
	defer span.End()

	h.logger.Debug("Get hello message in handler")

	message, err := h.service.GetHelloMessage(ctx) // Передаем контекст с трассировкой
	switch err {
	case nil:
		break
	default:
		h.logger.Error("Failed to get hello message", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.HTTPError{Error: "unknown error", Message: "Unknown internal error"})
		return
	}

	c.JSON(http.StatusOK, models.Message{Message: message})
}
