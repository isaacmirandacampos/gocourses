package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/health-check", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status":    "OK",
				"timestamp": time.Now().UTC(),
			})
		})
	}

	v1 := api.Group("/v1")
	{
		v1.GET("/courses", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})

		v1.POST("/courses", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})

		v1.DELETE("/courses", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})

		v1.PUT("/courses", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})

		v1.GET("/courses/:course_id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})
	}
}
