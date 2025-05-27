package controllers

import (
	"net/http"

	"github.com/exceptionteapots/gin-template/domains"
	"github.com/exceptionteapots/gin-template/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// HelloController responsible for processing HTTP requests related to the hello entity.
type HelloController struct {
	domain domains.HelloDomainInterface
	logger *zerolog.Logger
}

// HelloControllerInterface defines interface of hello entity controller
type HelloControllerInterface interface {
	GetHelloMessage(c *gin.Context)
	GetHelloMessageWithCache(c *gin.Context)
}

// NewHelloController creates new entity HelloController with domain and logger.
func NewHelloController(domain domains.HelloDomainInterface, logger *zerolog.Logger) *HelloController {
	return &HelloController{
		domain: domain,
		logger: logger,
	}
}

// GetHelloMessage process GET-request for get hello message
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

	// Get hello message in domain
	entity, err := h.domain.GetHelloMessage()

	// Process errors
	switch err {
	case nil:
		break
	default:
		h.logger.Error().Err(err).Msg("Failed to get hello message from domain")
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.HTTPError{Error: "unknown error", Message: "Unknown internal error"})
		return
	}

	// Send answer
	c.JSON(http.StatusOK, models.Message{Message: entity.Message})
}

// GetHelloMessageWithCache process GET-request for get hello message with cache
//
//	@Summary		Get Hello World message using database
//	@Description	get hello world
//	@Tags			Hello
//	@Produce		json
//	@Success		200	{object}	models.Message
//	@Failure		500	{object}	models.HTTPError
//	@Router			/hello-cached [GET]
func (h *HelloController) GetHelloMessageWithCache(c *gin.Context) {
	// Get hello message in domain
	entity, err := h.domain.GetHelloMessageWithCache(c.Request.Context())

	// Process errors
	switch err {
	case nil:
		break
	default:
		h.logger.Error().Err(err).Msg("Failed to get hello message from domain")
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.HTTPError{Error: "unknown error", Message: "Unknown internal error"})
		return
	}

	// Send answer
	c.JSON(http.StatusOK, models.Message{Message: entity.Message})
}
