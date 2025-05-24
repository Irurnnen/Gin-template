package controllers

import (
	"net/http"

	"github.com/exceptionteapots/gin-template/internal/domains"
	"github.com/exceptionteapots/gin-template/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type HelloController struct {
	domain domains.HelloDomainInterface
	logger *zerolog.Logger
}

type HelloControllerInterface interface {
	GetHelloMessage(c *gin.Context)
}

func NewHelloController(domain domains.HelloDomainInterface, logger *zerolog.Logger) *HelloController {
	return &HelloController{
		domain: domain,
		logger: logger,
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
func (h *HelloController) GetHelloMessage(c *gin.Context) {
	h.logger.Debug().Msg("Get hello message in controller")

	message, err := h.domain.GetHelloMessage()
	switch err {
	case nil:
		break
	default:
		h.logger.Error().Err(err).Msg("Failed to get hello message from domain")
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.HTTPError{Error: "unknown error", Message: "Unknown internal error"})
		return
	}

	c.JSON(http.StatusOK, models.Message{Message: message})
}
