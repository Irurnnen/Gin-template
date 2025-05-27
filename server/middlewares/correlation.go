package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	// Name of header for Correlation ID
	CorrelationIDHeader = "X-Correlation-ID"
	// Name of the key in context for storage Correlation ID
	CorrelationIDContextKey = "correlation_id"
)

// CorrelationIDMiddleware Create and processes Correlation ID for requests
func CorrelationIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to get Correlation ID from header
		correlationID := c.GetHeader(CorrelationIDHeader)

		// If header is not found - generate new UUID
		if correlationID == "" {
			correlationID = uuid.New().String()
		}

		// Set Correlation ID in context of gin
		c.Set(CorrelationIDContextKey, correlationID)

		// Add Correlation ID in header of response
		c.Header(CorrelationIDHeader, correlationID)

		// Next middleware
		c.Next()
	}
}

// GetCorrelationIDFromContext Get Correlation ID from context
func GetCorrelationIDFromContext(c *gin.Context) string {
	return c.GetString(CorrelationIDContextKey)
}
