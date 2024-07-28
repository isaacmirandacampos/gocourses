package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"timestamp": time.Now().UTC(),
	})
}
