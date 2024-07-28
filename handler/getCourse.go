package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
