package server

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.New()

	// Add middlewares
	router.Use(
		gin.Logger(),
		gin.Recovery(),
		// TODO: Add tracer
	)

	// Setup routes
	v1 := router.Group("/v1")
	{
		internal := v1.Group("/internal")
		{
			internal.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "pong"})
			})
		}
	}

	return router
}
